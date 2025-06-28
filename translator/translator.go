package translator

import (
	"fmt"
	"os"
	"strings"

	"github.com/xvimnt/OLC2_PROYECTO2_G15/ast"
)

// VarType represents the type of a variable within the translator.
type VarType int

const (
	TypeUnknown VarType = iota
	TypeInt
	TypeFloat
	TypeString
	TypeBool
	TypeVoid
	TypeSlice
	TypeArray
	TypeStruct
)

func (v VarType) String() string {
	switch v {
	case TypeInt:
		return "int"
	case TypeFloat:
		return "float64"
	case TypeString:
		return "string"
	case TypeBool:
		return "bool"
	case TypeVoid:
		return "void"
	case TypeSlice:
		return "[]int" // Note: This is a simplified representation for now
	case TypeArray:
		return "array" // Note: Simplified
	case TypeStruct:
		return "struct" // Note: Simplified
	default:
		return "unknown"
	}
}

// Translator translates AST nodes into assembly code.
type Translator struct {
	asm                []string            // Stores generated .text section assembly lines
	dataSection        []string            // Stores generated .data section assembly lines
	symbolTables       []map[string]string // Stack of maps: original name -> mangled name
	varInfo            map[string]VarType  // Map: mangled name -> type
	scopeCounter       int
	scopeIDStack       []int
	stringCounter      int               // For generating unique string labels
	labelCounter       int               // For generating unique labels
	needsPrintf        bool              // Tracks if printf is used (for .extern printf)
	needsAtoi          bool              // Tracks if atoi is used
	needsAtof          bool              // Tracks if atof is used
	hasIntFormatStr    bool              // Tracks if the integer format string has been added
	hasFloatFormatStr  bool              // Tracks if the float format string has been added
	hasStringFormatStr bool              // Tracks if the string format string has been added
	trueStrLabel       string            // Label for the "true" string literal
	falseStrLabel      string            // Label for the "false" string literal
	emptyStrLabel      string            // Label for the "" string literal
	currentFuncDef     *ast.FunctionDecl // Keep track of the current function being defined
	funcSyms           map[string]*ast.FunctionDecl
	DebugMode          bool
	needsStringHelpers bool // Tracks if string concatenation helpers are needed
	needsStrcmp        bool // Tracks if strcmp is needed
	needsPrintIntSlice bool // Tracks if the print_int_slice helper is needed

	// Register allocation
	intRegs   []bool // Availability of general-purpose integer registers (X9-X15)
	floatRegs []bool // Availability of general-purpose float registers (D8-D15)

	// Loop context
	breakLabels    []string // Stack of labels for 'break' statements
	continueLabels []string // Stack of labels for 'continue' statements
}

const (
	numIntRegs    = 7 // X9-X15
	intRegStart   = 9
	numFloatRegs  = 8 // D8-D15
	floatRegStart = 8
)

// NewTranslator creates a new Translator instance.
func NewTranslator(debugMode bool) *Translator {
	t := &Translator{
		asm:                make([]string, 0),
		dataSection:        make([]string, 0),
		symbolTables:       []map[string]string{make(map[string]string)}, // Global scope
		varInfo:            make(map[string]VarType),
		scopeCounter:       0,
		scopeIDStack:       []int{0}, // Global scope ID
		stringCounter:      0,
		labelCounter:       0,
		needsPrintf:        false,
		needsAtoi:          false,
		needsAtof:          false,
		hasIntFormatStr:    false,
		hasFloatFormatStr:  false,
		hasStringFormatStr: false,
		trueStrLabel:       "",
		falseStrLabel:      "",
		emptyStrLabel:      "",

		currentFuncDef: nil,
		funcSyms:       make(map[string]*ast.FunctionDecl),
		DebugMode:      debugMode,
		needsStrcmp:    false,
		intRegs:        make([]bool, numIntRegs),
		floatRegs:      make([]bool, numFloatRegs),
	}

	// Initialize all registers as available
	for i := 0; i < numIntRegs; i++ {
		t.intRegs[i] = true
	}
	for i := 0; i < numFloatRegs; i++ {
		t.floatRegs[i] = true
	}

	return t
}

// --- Register Management ---

// acquireIntRegister finds and returns an available integer register.
func (t *Translator) acquireIntRegister() int {
	for i, available := range t.intRegs {
		if available {
			t.intRegs[i] = false
			return i + intRegStart
		}
	}
	panic("No more integer registers available!") // Or handle more gracefully
}

// releaseIntRegister marks an integer register as available.
func (t *Translator) releaseIntRegister(reg int) {
	if reg >= intRegStart && reg < intRegStart+numIntRegs {
		t.intRegs[reg-intRegStart] = true
	}
}

// acquireFloatRegister finds and returns an available float register.
func (t *Translator) acquireFloatRegister() int {
	for i, available := range t.floatRegs {
		if available {
			t.floatRegs[i] = false
			return i + floatRegStart
		}
	}
	panic("No more float registers available!")
}

// releaseFloatRegister marks a float register as available.
func (t *Translator) releaseFloatRegister(reg int) {
	if reg >= floatRegStart && reg < floatRegStart+numFloatRegs {
		t.floatRegs[reg-floatRegStart] = true
	}
}

// addFloatData adds a float to the .data section and returns its label.
func (t *Translator) addFloatData(floatStr string) string {
	label := fmt.Sprintf("F%d", t.stringCounter)
	t.stringCounter++
	// Align to 8-byte boundary for doubles
	t.dataSection = append(t.dataSection, ".align 3")
	t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .double %s", label, floatStr))
	return label
}

// addStringData adds a string to the .data section and returns its label.
// It uses %q to handle proper quoting and escaping for the assembler.
func (t *Translator) enterScope() {
	t.scopeCounter++
	t.scopeIDStack = append(t.scopeIDStack, t.scopeCounter)
	t.symbolTables = append(t.symbolTables, make(map[string]string))
}

func (t *Translator) exitScope() {
	t.scopeIDStack = t.scopeIDStack[:len(t.scopeIDStack)-1]
	t.symbolTables = t.symbolTables[:len(t.symbolTables)-1]
}

func (t *Translator) currentScopeID() int {
	return t.scopeIDStack[len(t.scopeIDStack)-1]
}

// defineSymbol creates a mangled name for a variable, stores it, and returns it.
func (t *Translator) defineSymbol(name string, vtype VarType) string {
	// Mangle name to be unique across scopes
	mangledName := fmt.Sprintf("%s_%d", name, t.currentScopeID())
	currentScope := t.symbolTables[len(t.symbolTables)-1]
	currentScope[name] = mangledName
	t.varInfo[mangledName] = vtype
	return mangledName
}

// lookupSymbol finds a variable's mangled name and type, searching from the current scope outwards.
func (t *Translator) lookupSymbol(name string) (mangledName string, vtype VarType, exists bool) {
	for i := len(t.symbolTables) - 1; i >= 0; i-- {
		if mangledName, ok := t.symbolTables[i][name]; ok {
			vtype := t.varInfo[mangledName]
			return mangledName, vtype, true
		}
	}
	return "", TypeUnknown, false
}

func (t *Translator) addStringData(strContent string) string {
	label := fmt.Sprintf("str%d", t.stringCounter)
	t.stringCounter++
	// Align to 4-byte boundary for strings, which is good practice.
	t.dataSection = append(t.dataSection, ".align 2")
	t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .asciz %q", label, strContent))
	return label
}

// addData adds a line to the .data section of the assembly code.
func (t *Translator) addData(line string) {
	t.dataSection = append(t.dataSection, line)
}

// newLabel generates a new unique label with a given prefix.
func (t *Translator) newLabel(prefix string) string {
	t.labelCounter++
	return fmt.Sprintf("%s%d", prefix, t.labelCounter)
}

// GetAssembly returns the generated assembly code.
func (t *Translator) GetAssembly() []string {
	var finalAsm []string

	// .data section
	if len(t.dataSection) > 0 {
		finalAsm = append(finalAsm, ".data")
		finalAsm = append(finalAsm, t.dataSection...)
		finalAsm = append(finalAsm, "") // Blank line
	}

	// .text section
	if len(t.asm) > 0 {
		if t.needsStringHelpers {
			finalAsm = append(finalAsm, ".extern malloc")
			finalAsm = append(finalAsm, ".extern strlen")
			finalAsm = append(finalAsm, ".extern strcpy")
			finalAsm = append(finalAsm, ".extern strcat")
		}
		if t.needsStrcmp {
			finalAsm = append(finalAsm, ".extern strcmp")
		}
		if t.needsAtof {
			finalAsm = append(finalAsm, ".extern atof")
		}
		if t.needsAtoi {
			finalAsm = append(finalAsm, ".extern atoi")
		}
		if t.needsPrintf {
			finalAsm = append(finalAsm, ".extern printf")
		}
		finalAsm = append(finalAsm, ".text")
		finalAsm = append(finalAsm, t.asm...)

		if t.needsPrintIntSlice {
			finalAsm = append(finalAsm, t.getPrintIntSliceAsm()...)
		}
	}

	return finalAsm
}

func (t *Translator) ensurePrintIntSliceHelper() {
	if t.needsPrintIntSlice {
		return
	}
	t.needsPrintIntSlice = true
	t.addData("open_bracket_str: .asciz \"[\"")
	t.addData("close_bracket_str: .asciz \"]\"")
	t.addData("comma_space_str: .asciz \", \"")
	t.addData("slice_int_format_str: .asciz \"%d\"")
}

func (t *Translator) getPrintIntSliceAsm() []string {
	return []string{
		"",
		"// --- Helper function to print a slice of integers ---",
		"print_int_slice:",
		"    STP X29, X30, [SP, #-48]!",
		"    STP X19, X20, [SP, #16]",
		"    STP X21, X22, [SP, #32]",
		"    MOV X29, SP",
		"    MOV X19, X0                       // Save slice descriptor address in X19",
		"",
		"    // Print opening bracket",
		"    LDR X0, =open_bracket_str",
		"    BL printf",
		"",
		"    // Load slice details",
		"    LDR X20, [X19, #0]                // X20 = pointer to slice data",
		"    LDR X21, [X19, #8]                // X21 = length of slice",
		"",
		"    // Loop setup",
		"    MOV X22, #0                       // X22 = loop counter (i)",
		"    B .Lprint_slice_loop_test",
		"",
		".Lprint_slice_loop_start:",
		"    // Print comma and space if not the first element",
		"    CMP X22, #0",
		"    BEQ .Lprint_slice_skip_comma",
		"    LDR X0, =comma_space_str",
		"    BL printf",
		"",
		".Lprint_slice_skip_comma:",
		"    // Load element: value = data[i]",
		"    LSL X1, X22, #3                   // Offset = i * 8 (since elements are .quad, 8 bytes)",
		"    LDR X1, [X20, X1]                 // Load element from data + offset into X1 (arg for printf)",
		"",
		"    // Print element",
		"    LDR X0, =slice_int_format_str",
		"    BL printf",
		"",
		"    // Increment and loop",
		"    ADD X22, X22, #1",
		"",
		".Lprint_slice_loop_test:",
		"    CMP X22, X21",
		"    B.LT .Lprint_slice_loop_start",
		"",
		"    // Print closing bracket",
		"    LDR X0, =close_bracket_str",
		"    BL printf",
		"",
		"    // Restore registers and return",
		"    LDP X21, X22, [SP, #32]",
		"    LDP X19, X20, [SP, #16]",
		"    LDP X29, X30, [SP], #48",
		"    RET",
	}
}

func (t *Translator) addAsm(instr string, args ...interface{}) {
	if len(args) > 0 {
		t.asm = append(t.asm, fmt.Sprintf(instr, args...))
	} else {
		t.asm = append(t.asm, instr)
	}
}

// --- Visitor Interface Implementations ---

func (t *Translator) VisitProgram(node *ast.Program) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting Program")
	}

	// First pass: Register all function declarations.
	// This allows forward references to functions.
	if t.DebugMode {
		fmt.Println("Translator: First pass - registering functions.")
	}
	for _, decl := range node.Declarations {
		if fnDecl, ok := decl.(*ast.FunctionDecl); ok {
			if fnDecl.Name != nil {
				if t.DebugMode {
					fmt.Printf("  - Registering function: %s\n", fnDecl.Name.Name)
				}
				t.funcSyms[fnDecl.Name.Name] = fnDecl
			}
		}
	}

	// Second pass: Translate all declarations.
	if t.DebugMode {
		fmt.Println("Translator: Second pass - translating declarations.")
	}
	for _, decl := range node.Declarations {
		decl.Accept(t)
	}

	return nil
}

// Declarations
func (t *Translator) VisitStructDecl(node *ast.StructDecl) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting StructDecl")
	}
	// TODO: Implement StructDecl translation
	return nil
}

func (t *Translator) VisitFieldDecl(node *ast.FieldDecl) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting FieldDecl")
	}
	// TODO: Implement FieldDecl translation
	return nil
}

func (t *Translator) VisitFunctionDecl(node *ast.FunctionDecl) interface{} {
	if t.DebugMode {
		if node.Name != nil {
			fmt.Printf("Translator.Visiting FunctionDecl: %s\n", node.Name.Name)
		} else {
			fmt.Println("Translator.Visiting FunctionDecl: <anonymous?>") // Should not happen for valid programs
		}
	}
	t.currentFuncDef = node

	if node.Name != nil && node.Name.Name == "main" {
		t.addAsm(".global main")
		t.addAsm("main:")
	} else if node.Name != nil {
		t.addAsm(".global %s", node.Name.Name)
		t.addAsm("%s:", node.Name.Name)
	} else {
		// Handle anonymous functions or error, though V doesn't have them at top level like this
		t.addAsm("anonymous_func_%d:", t.stringCounter) // Placeholder for unnamed funcs
		t.stringCounter++
	}

	// Prologue
	t.addAsm("    STP X29, X30, [SP, #-16]!") // Save FP, LR to stack, pre-decrement SP by 16 (write-back)
	t.addAsm("    MOV X29, SP")               // Set current stack pointer as the new Frame Pointer

	// TODO: Allocate space for local variables based on function needs

	t.enterScope() // Scope for parameters and locals

	// Process parameters
	if node.Parameters != nil {
		intArgReg := 0
		floatArgReg := 0
		for _, param := range node.Parameters {
			if param.Name == nil {
				continue // Should not happen in valid code
			}
			paramName := param.Name.Name

			// Use the helper to determine the parameter's type
			vtype := t.typeNodeToVarType(param.Type)
			if vtype == TypeVoid {
				panic(fmt.Sprintf("Function parameter '%s' cannot be void", paramName))
			}

			// Define the symbol in the current scope
			mangledName := t.defineSymbol(paramName, vtype)

			// Per ARM64 calling convention, first args are in registers.
			// We'll store them to memory (following the pattern for local vars in this compiler)
			switch vtype {
			case TypeInt, TypeBool, TypeString: // Strings are pointers (int-like)
				if intArgReg < 8 {
					// Allocate in .data section and store from register
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .quad 0", mangledName))
					t.addAsm("    // Store param '%s' from register X%d to memory", paramName, intArgReg)
					t.addAsm("    LDR X9, =%s", mangledName)
					t.addAsm("    STR X%d, [X9]", intArgReg)
					intArgReg++
				} else {
					// TODO: Handle parameters passed on the stack
				}
			case TypeFloat:
				if floatArgReg < 8 {
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .double 0.0", mangledName))
					t.addAsm("    // Store param '%s' from register D%d to memory", paramName, floatArgReg)
					t.addAsm("    LDR X9, =%s", mangledName)
					t.addAsm("    STR D%d, [X9]", floatArgReg)
					floatArgReg++
				} else {
					// TODO: Handle float parameters passed on the stack
				}
			}
		}
	}

	if node.Body != nil {
		node.Body.Accept(t)
	}

	t.exitScope()

	// Epilogue
	// Ensure a return path even if no explicit return statement for void functions (like main often is implicitly)
	// For non-void functions, an explicit return statement should handle loading the return value.
	// For main, or functions ending without explicit return, this provides a standard exit.
	// Epilogue for main/_start should handle process exit.
	// For other functions, it's a standard return.
	if node.Name != nil {
		epilogueLabel := fmt.Sprintf(".L%s_epilogue", node.Name.Name)
		t.addAsm(epilogueLabel + ":") // Label for potential jumps to epilogue
	} else {
		epilogueLabel := fmt.Sprintf(".L_anonymous_func_%d_epilogue", t.stringCounter-1)
		t.addAsm(epilogueLabel + ":") // Match potential anonymous label
	}
	// For void functions, ensure a default return value of 0.
	// Non-void functions must have explicit return statements that set X0.
	// The epilogue is jumped to from return statements, so we must not overwrite X0 here.
	functionReturnType := t.typeNodeToVarType(node.ReturnType)
	if functionReturnType == TypeVoid {
		t.addAsm("    MOV W0, #0") // Default return code 0 for void functions
	}
	t.addAsm("    LDP X29, X30, [SP], #16") // Restore FP, LR from stack, post-increment SP by 16
	t.addAsm("    RET")
	t.addAsm("") // Add a blank line for readability after function definition
	t.currentFuncDef = nil
	return nil
}

func (t *Translator) VisitReceiverDecl(node *ast.ReceiverDecl) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ReceiverDecl")
	}
	// TODO: Implement ReceiverDecl translation
	return nil
}

func (t *Translator) VisitParameterDecl(node *ast.ParameterDecl) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ParameterDecl")
	}
	// TODO: Implement ParameterDecl translation
	return nil
}

func (t *Translator) VisitVarDecl(node *ast.VarDecl) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting VarDecl for '%s'\n", node.Name.Name)
	}

	// Handle declaration without initializer (e.g., var x int)
	if node.Initializer == nil {
		varType := t.typeNodeToVarType(node.ExplicitType)
		if varType == TypeUnknown || varType == TypeVoid {
			panic(fmt.Sprintf("Cannot declare variable '%s' with invalid type", node.Name.Name))
		}

		mangledName := t.defineSymbol(node.Name.Name, varType)

		// Add variable to .data section, initializing to its zero-value.
		t.addAsm("    // Declaring %s without initializer", node.Name.Name)
		switch varType {
		case TypeInt, TypeBool:
			t.addData(fmt.Sprintf("%s: .quad 0", mangledName))
		case TypeFloat:
			// .double requires alignment
			t.addData(".align 3")
			t.addData(fmt.Sprintf("%s: .double 0.0", mangledName))
		case TypeString, TypeSlice: // Pointers
			if t.emptyStrLabel == "" {
				// Create the global empty string if it doesn't exist yet.
				t.emptyStrLabel = t.addStringData("")
			}
			t.addData(fmt.Sprintf("%s: .quad %s", mangledName, t.emptyStrLabel)) // Initialize to pointer to empty string
		default:
			panic(fmt.Sprintf("Unhandled zero-value initialization for type %s", varType))
		}
		return nil
	}

	// Evaluate the initializer expression first
	initResult := node.Initializer.Accept(t)
	res, ok := initResult.(ExpressionResult)
	if !ok {
		panic(fmt.Sprintf("Initializer for %s did not return an ExpressionResult", node.Name.Name))
	}

	// Define the symbol with the type from the expression result
	mangledName := t.defineSymbol(node.Name.Name, res.Type)

	// Add variable to .data section, initializing to zero/null.
	switch res.Type {
	case TypeInt, TypeBool:
		t.addData(fmt.Sprintf("%s: .quad 0", mangledName))
	case TypeFloat:
		t.addData(fmt.Sprintf("%s: .double 0.0", mangledName))
	case TypeString, TypeSlice:
		t.addData(fmt.Sprintf("%s: .quad 0", mangledName)) // Store pointer, init to null
	}

	// Spill the initializer result to the stack to prevent register conflicts.
	t.addAsm("    SUB SP, SP, #16")
	switch res.Type {
	case TypeFloat:
		t.addAsm("    STR D%d, [SP]", res.Reg)
		t.releaseFloatRegister(res.Reg)
	default: // Int, String, Bool, Slice, etc.
		t.addAsm("    STR X%d, [SP]", res.Reg)
		t.releaseIntRegister(res.Reg)
	}

	// Acquire registers for storing the value from stack to variable
	addrReg := t.acquireIntRegister()
	valReg := t.acquireIntRegister() // Use a separate register for the value

	t.addAsm("    LDR X%d, =%s", addrReg, mangledName)

	switch res.Type {
	case TypeInt:
		t.addAsm("    LDR W%d, [SP]", valReg)
		t.addAsm("    STR W%d, [X%d]", valReg, addrReg)
	case TypeFloat:
		valFloatReg := t.acquireFloatRegister()
		t.addAsm("    LDR D%d, [SP]", valFloatReg)
		t.addAsm("    STR D%d, [X%d]", valFloatReg, addrReg)
		t.releaseFloatRegister(valFloatReg)
	case TypeString, TypeSlice:
		t.addAsm("    LDR X%d, [SP]", valReg)
		t.addAsm("    STR X%d, [X%d]", valReg, addrReg)
	case TypeBool:
		t.addAsm("    LDRB W%d, [SP]", valReg)
		t.addAsm("    STRB W%d, [X%d]", valReg, addrReg)
	}

	// Release registers and clean up stack
	t.releaseIntRegister(addrReg)
	t.releaseIntRegister(valReg)
	t.addAsm("    ADD SP, SP, #16")
	return nil
}

// Statements
func (t *Translator) VisitBlockStmt(node *ast.BlockStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting BlockStmt")
	}
	t.enterScope()
	for _, stmt := range node.Statements {
		stmt.Accept(t)
	}
	t.exitScope()
	return nil
}

func (t *Translator) VisitAssignStmt(node *ast.AssignStmt) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting AssignStmt, Operator: %s\n", node.Operator)
	}

	// 1. Get the variable name from the left side (LHS)
	varName, ok := node.Left.(*ast.IdentifierExpr)
	if !ok {
		panic(fmt.Sprintf("Unsupported L-value in assignment: %T", node.Left))
	}

	// 2. Handle short variable declaration (:=) vs. simple assignment (=)
	if node.Operator == ":=" {
		// --- Short Variable Declaration (:=) ---
		// Evaluate the RHS to get the value and type
		initResult := node.Right.Accept(t)
		res, ok := initResult.(ExpressionResult)
		if !ok {
			panic(fmt.Sprintf("Initializer for %s did not return an ExpressionResult", varName.Name))
		}

		// Define the new symbol in the symbol table with the inferred type
		mangledName := t.defineSymbol(varName.Name, res.Type)

		// Add variable to .data section, initializing to zero/null.
		switch res.Type {
		case TypeInt, TypeBool:
			t.addData(fmt.Sprintf("%s: .quad 0", mangledName))
		case TypeFloat:
			t.addData(fmt.Sprintf("%s: .double 0.0", mangledName))
		case TypeString:
			t.addData(fmt.Sprintf("%s: .quad 0", mangledName)) // Store pointer, init to null
		default:
			panic(fmt.Sprintf("Unsupported type for variable declaration: %s", res.Type))
		}

		// Spill the initializer result to the stack to prevent register conflicts.
		t.addAsm("    SUB SP, SP, #16")
		switch res.Type {
		case TypeFloat:
			t.addAsm("    STR D%d, [SP]", res.Reg)
			t.releaseFloatRegister(res.Reg)
		default: // Int, String, Bool, Slice, etc.
			t.addAsm("    STR X%d, [SP]", res.Reg)
			t.releaseIntRegister(res.Reg)
		}

		// Acquire registers for storing the value from stack to variable
		addrReg := t.acquireIntRegister()
		valReg := t.acquireIntRegister() // Use a separate register for the value

		t.addAsm("    LDR X%d, =%s", addrReg, mangledName)

		switch res.Type {
		case TypeInt:
			t.addAsm("    LDR W%d, [SP]", valReg)
			t.addAsm("    STR W%d, [X%d]", valReg, addrReg)
		case TypeFloat:
			valFloatReg := t.acquireFloatRegister()
			t.addAsm("    LDR D%d, [SP]", valFloatReg)
			t.addAsm("    STR D%d, [X%d]", valFloatReg, addrReg)
			t.releaseFloatRegister(valFloatReg)
		case TypeString:
			t.addAsm("    LDR X%d, [SP]", valReg)
			t.addAsm("    STR X%d, [X%d]", valReg, addrReg)
		case TypeBool:
			t.addAsm("    LDRB W%d, [SP]", valReg)
			t.addAsm("    STRB W%d, [X%d]", valReg, addrReg)
		}

		// Release registers and clean up stack
		t.releaseIntRegister(addrReg)
		t.releaseIntRegister(valReg)
		t.addAsm("    ADD SP, SP, #16")
	} else if node.Operator == "=" {
		// --- Simple Assignment (=) ---
		// Lookup the existing variable
		mangledName, varType, exists := t.lookupSymbol(varName.Name)
		if !exists {
			panic(fmt.Sprintf("Assignment to undeclared variable: %s", varName.Name))
		}

		// Evaluate the RHS
		rightResultRaw := node.Right.Accept(t)
		rightResult, ok := rightResultRaw.(ExpressionResult)
		if !ok {
			panic("RHS of assignment did not return an ExpressionResult")
		}

		// Basic type check
		if varType != rightResult.Type {
			// Allow int to float promotion
			if !(varType == TypeFloat && rightResult.Type == TypeInt) {
				panic(fmt.Sprintf("Type mismatch in assignment to %s. Expected %s, got %s", varName.Name, varType, rightResult.Type))
			}
		}

		// Acquire a dedicated register for the address to avoid clobbering the value register.
		addrReg := t.acquireIntRegister()
		t.addAsm("    // Storing value for assignment to %s", varName.Name)
		t.addAsm("    LDR X%d, =%s", addrReg, mangledName) // Load address of variable into addrReg

		// Handle type promotion if necessary (int to float)
		if varType == TypeFloat && rightResult.Type == TypeInt {
			t.addAsm("    // Promoting RHS from INT to FLOAT for assignment")
			promotedFloatReg := t.acquireFloatRegister()
			t.addAsm("    SCVTF D%d, W%d", promotedFloatReg, rightResult.Reg)
			t.releaseIntRegister(rightResult.Reg)
			rightResult = ExpressionResult{Reg: promotedFloatReg, Type: TypeFloat}
		}

		// Spill the right-hand side result to the stack to prevent register conflicts.
		t.addAsm("    SUB SP, SP, #16")
		switch rightResult.Type {
		case TypeFloat:
			t.addAsm("    STR D%d, [SP]", rightResult.Reg)
			t.releaseFloatRegister(rightResult.Reg)
		default: // Int, String, Bool, Slice, etc.
			t.addAsm("    STR X%d, [SP]", rightResult.Reg)
			t.releaseIntRegister(rightResult.Reg)
		}

		// Acquire registers for storing the value from stack to variable
		valReg := t.acquireIntRegister() // Use a separate register for the value

		switch varType {
		case TypeInt:
			t.addAsm("    LDR W%d, [SP]", valReg)
			t.addAsm("    STR W%d, [X%d]", valReg, addrReg)
		case TypeFloat:
			valFloatReg := t.acquireFloatRegister()
			t.addAsm("    LDR D%d, [SP]", valFloatReg)
			t.addAsm("    STR D%d, [X%d]", valFloatReg, addrReg)
			t.releaseFloatRegister(valFloatReg)
		case TypeString:
			t.addAsm("    LDR X%d, [SP]", valReg)
			t.addAsm("    STR X%d, [X%d]", valReg, addrReg)
		case TypeBool:
			t.addAsm("    LDRB W%d, [SP]", valReg)
			t.addAsm("    STRB W%d, [X%d]", valReg, addrReg)
		}

		// Release registers and clean up stack
		t.releaseIntRegister(addrReg)
		t.releaseIntRegister(valReg)
		t.addAsm("    ADD SP, SP, #16")
	} else { // Compound assignment operators
		// 1. Get the variable name from the left side (LHS)
		varName, ok := node.Left.(*ast.IdentifierExpr)
		if !ok {
			panic(fmt.Sprintf("Unsupported L-value in assignment: %T", node.Left))
		}

		// 2. Lookup the existing variable
		mangledName, varType, exists := t.lookupSymbol(varName.Name)
		if !exists {
			panic(fmt.Sprintf("Assignment to undeclared variable: %s", varName.Name))
		}

		// 3. Evaluate the RHS
		rightResultRaw := node.Right.Accept(t)
		rightResult, ok := rightResultRaw.(ExpressionResult)
		if !ok {
			panic("RHS of compound assignment did not return an ExpressionResult")
		}

		// 4. Acquire a register to hold the address of the variable.
		addrReg := t.acquireIntRegister()
		t.addAsm("    // --- Start Compound Assignment: %s ---", node.Operator)
		t.addAsm("    LDR X%d, =%s", addrReg, mangledName)

		if varType == TypeFloat {
			// 4a. Load the current value of the float variable
			currentValReg := t.acquireFloatRegister()
			t.addAsm("    LDR D%d, [X%d]", currentValReg, addrReg)

			// 5a. Perform the float operation
			switch node.Operator {
			case "+=":
				if rightResult.Type == TypeFloat {
					t.addAsm("    FADD D%d, D%d, D%d", currentValReg, currentValReg, rightResult.Reg)
				} else {
					panic(fmt.Sprintf("Type mismatch for '+=' operator: %s and %s", varType, rightResult.Type))
				}
			default:
				panic(fmt.Sprintf("Unsupported compound assignment operator for floats: %s", node.Operator))
			}

			// 6a. Store the new float value back
			t.addAsm("    STR D%d, [X%d]", currentValReg, addrReg)

			// 7a. Release float registers
			t.releaseFloatRegister(currentValReg)
			t.releaseFloatRegister(rightResult.Reg)
		} else { // Integer operation
			// 4b. Load the current value of the integer variable
			currentValReg := t.acquireIntRegister()
			t.addAsm("    LDRSW X%d, [X%d]", currentValReg, addrReg)

			// 5b. Perform the integer operation
			switch node.Operator {
			case "+=":
				if rightResult.Type == TypeInt {
					t.addAsm("    ADD X%d, X%d, X%d", currentValReg, currentValReg, rightResult.Reg)
				} else {
					panic(fmt.Sprintf("Type mismatch for '+=' operator: %s and %s", varType, rightResult.Type))
				}
			case "-=":
				if rightResult.Type == TypeInt {
					t.addAsm("    SUB X%d, X%d, X%d", currentValReg, currentValReg, rightResult.Reg)
				} else {
					panic(fmt.Sprintf("Type mismatch for '-=' operator: %s and %s", varType, rightResult.Type))
				}
			case "*=":
				if rightResult.Type == TypeInt {
					t.addAsm("    MUL X%d, X%d, X%d", currentValReg, currentValReg, rightResult.Reg)
				} else {
					panic(fmt.Sprintf("Type mismatch for '*=' operator: %s and %s", varType, rightResult.Type))
				}
			case "/=":
				if rightResult.Type == TypeInt {
					t.addAsm("    SDIV X%d, X%d, X%d", currentValReg, currentValReg, rightResult.Reg)
				} else {
					panic(fmt.Sprintf("Type mismatch for '/=' operator: %s and %s", varType, rightResult.Type))
				}
			case "%=":
				if rightResult.Type == TypeInt {
					// result = a - (a/n) * n
					quotientReg := t.acquireIntRegister()
					t.addAsm("    SDIV X%d, X%d, X%d", quotientReg, currentValReg, rightResult.Reg)
					t.addAsm("    MUL X%d, X%d, X%d", quotientReg, quotientReg, rightResult.Reg)
					t.addAsm("    SUB X%d, X%d, X%d", currentValReg, currentValReg, quotientReg)
					t.releaseIntRegister(quotientReg)
				} else {
					panic(fmt.Sprintf("Type mismatch for '%%=' operator: %s and %s", varType, rightResult.Type))
				}
			default:
				panic(fmt.Sprintf("Unsupported compound assignment operator for ints: %s", node.Operator))
			}

			// 6b. Store the new integer value back
			t.addAsm("    STR W%d, [X%d]", currentValReg, addrReg)

			// 7b. Release integer registers
			t.releaseIntRegister(currentValReg)
			t.releaseIntRegister(rightResult.Reg)
		}

		// Release the address register
		t.releaseIntRegister(addrReg)
		t.addAsm("    // --- End Compound Assignment ---")
	}

	return nil
}

func (t *Translator) VisitExpressionStmt(node *ast.ExpressionStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ExpressionStmt")
	}
	// Evaluate the expression. This is typically a function call.
	result := node.Expression.Accept(t)

	// If the expression returned a value (e.g., a function call that returns something),
	// we need to release the register it was stored in, since the result is not being used.
	if res, ok := result.(ExpressionResult); ok {
		switch res.Type {
		case TypeInt, TypeBool, TypeString:
			t.releaseIntRegister(res.Reg)
		case TypeFloat:
			t.releaseFloatRegister(res.Reg)
		}
	}
	return nil
}

func (t *Translator) VisitBreakStmt(node *ast.BreakStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting BreakStmt")
	}
	if len(t.breakLabels) == 0 {
		panic("break statement outside of loop or switch")
	}
	// Jump to the current loop's exit label
	breakLabel := t.breakLabels[len(t.breakLabels)-1]
	t.addAsm("    B %s", breakLabel)
	return nil
}

func (t *Translator) VisitContinueStmt(node *ast.ContinueStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ContinueStmt")
	}
	if len(t.continueLabels) == 0 {
		panic("continue statement outside of loop")
	}
	// Jump to the current loop's post-statement label
	continueLabel := t.continueLabels[len(t.continueLabels)-1]
	t.addAsm("    B %s", continueLabel)
	return nil
}

func (t *Translator) VisitIncDecStmt(node *ast.IncDecStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting IncDecStmt")
	}

	// Get the variable name from the left side (LHS)
	var varName string
	if ident, ok := node.LValue.(*ast.IdentifierExpr); ok {
		varName = ident.Name
	} else {
		fmt.Fprintf(os.Stderr, "Unsupported L-value in inc/dec statement: %T\n", node.LValue)
		return nil
	}

	// We need to know the type of the variable to use the correct store instruction.
	mangledName, varType, typeExists := t.lookupSymbol(varName)
	if !typeExists {
		fmt.Fprintf(os.Stderr, "Inc/dec on undeclared variable: %s\n", varName)
		return nil
	}

	switch varType {
	case TypeInt:
		t.addAsm("    // --- Start of integer inc/dec on %s ---", varName)
		t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
		t.addAsm("    LDR W11, [X10]")            // Load current value of var

		switch node.Operator {
		case "++":
			t.addAsm("    ADD W11, W11, #1") // Increment
		case "--":
			t.addAsm("    SUB W11, W11, #1") // Decrement
		}

		t.addAsm("    STR W11, [X10]") // Store result back
		t.addAsm("    // --- End of integer inc/dec on %s ---", varName)
		t.addAsm("")
	case TypeFloat:
		t.addAsm("    // --- Start of float inc/dec on %s ---", varName)
		// Create a label for 1.0 in the data section.
		oneLabel := t.newLabel("float_one")
		t.addData(fmt.Sprintf("%s: .double 1.0", oneLabel))

		t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
		t.addAsm("    LDR D8, [X10]")             // Load current value of var into float register D8

		t.addAsm("    LDR X11, =%s", oneLabel) // Load address of 1.0
		t.addAsm("    LDR D9, [X11]")          // Load 1.0 into D9

		switch node.Operator {
		case "++":
			t.addAsm("    FADD D8, D8, D9") // Increment
		case "--":
			t.addAsm("    FSUB D8, D8, D9") // Decrement
		}

		t.addAsm("    STR D8, [X10]") // Store result back
		t.addAsm("    // --- End of float inc/dec on %s ---", varName)
		t.addAsm("")
	default:
		fmt.Fprintf(os.Stderr, "Inc/dec on unsupported type for variable: %s (%s)\n", varName, varType)
		return nil
	}

	return nil
}

func (t *Translator) VisitIfStmt(node *ast.IfStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting IfStmt")
	}

	elseLabel := t.newLabel(".Lelse")
	endLabel := t.newLabel(".Lendif")

	// 1. Evaluate the condition
	conditionResult, ok := node.Condition.Accept(t).(ExpressionResult)
	if !ok {
		// This might happen if the condition is a function call that doesn't return a value
		// or another expression type we haven't handled to return an ExpressionResult.
		// For now, we'll panic. A more robust compiler would have better error handling.
		panic("Condition in if statement did not return a valid ExpressionResult.")
	}

	// 2. Compare the result of the condition and branch if false
	// The result of a boolean expression should be in a register (0 for false, 1 for true).
	t.addAsm("    // If statement condition check")
	t.addAsm("    CMP W%d, #0", conditionResult.Reg) // Compare the result with 0
	t.releaseIntRegister(conditionResult.Reg)        // Free up the register

	if node.Alternative != nil {
		t.addAsm("    B.EQ %s", elseLabel) // If condition is false (result == 0), jump to the 'else' part
	} else {
		t.addAsm("    B.EQ %s", endLabel) // If condition is false and no 'else', jump to the end
	}

	// 3. Translate the 'then' block (Consequence)
	t.addAsm("    // 'Then' block")
	node.Consequence.Accept(t)

	// 4. If there's an 'else' block, add an unconditional jump to the end to skip it
	if node.Alternative != nil {
		t.addAsm("    B %s", endLabel)
	}

	// 5. Emit the 'else' label and translate the 'else' block (Alternative)
	if node.Alternative != nil {
		t.addAsm("%s:", elseLabel)
		t.addAsm("    // 'Else' block")
		node.Alternative.Accept(t)
	}

	// 6. Emit the 'end' label
	t.addAsm("%s:", endLabel)

	return nil
}

func (t *Translator) VisitSwitchStmt(node *ast.SwitchStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting SwitchStmt")
	}

	endSwitchLabel := t.newLabel("switch_end")

	// Push the end label onto the break stack. This allows 'break' statements
	// within the switch to jump to the end of the switch.
	t.breakLabels = append(t.breakLabels, endSwitchLabel)
	var defaultLabel string
	if node.Default != nil {
		defaultLabel = t.newLabel("switch_default")
	} else {
		defaultLabel = endSwitchLabel // If no default, jump to the end
	}

	// 1. Evaluate the switch expression
	if node.Expression == nil {
		// This case (e.g., `switch {}`) is not handled.
		// The parser should enforce that an expression is present.
		return nil
	}
	switchExprResult := node.Expression.Accept(t).(ExpressionResult)

	// 2. Generate labels for each case body
	caseLabels := make([]string, len(node.Cases))
	for i := range node.Cases {
		caseLabels[i] = t.newLabel(fmt.Sprintf("switch_case_%d", i))
	}

	// 3. Generate comparison logic
	t.addAsm("    // --- Switch Statement ---")
	for i, caseClause := range node.Cases {
		for _, expr := range caseClause.Expressions {
			caseExprResult := expr.Accept(t).(ExpressionResult)

			// A more robust implementation would check switchExprResult.Type.
			if switchExprResult.Type == TypeString {
				t.needsStrcmp = true
				t.addAsm("    // Comparing with case: %s", expr.String())

				// Save the switch expression register because BL will clobber it
				t.addAsm("    SUB SP, SP, #16")
				t.addAsm("    STR X%d, [SP]", switchExprResult.Reg)

				// Set up args for strcmp
				t.addAsm("    MOV X0, X%d", switchExprResult.Reg)
				t.addAsm("    MOV X1, X%d", caseExprResult.Reg)

				t.addAsm("    BL strcmp")

				// Restore the switch expression register
				t.addAsm("    LDR X%d, [SP]", switchExprResult.Reg)
				t.addAsm("    ADD SP, SP, #16")

				t.addAsm("    CMP W0, #0") // strcmp returns 0 on match
				t.addAsm("    BEQ %s", caseLabels[i])

			} else { // Assuming integer comparison for other types
				t.addAsm("    // Comparing with case: %s", expr.String())
				// switchExprResult.Reg holds the integer value.
				// caseExprResult.Reg holds the case integer value.
				t.addAsm("    CMP W%d, W%d", switchExprResult.Reg, caseExprResult.Reg)
				t.addAsm("    BEQ %s", caseLabels[i])
			}
			t.releaseIntRegister(caseExprResult.Reg)
		}
	}

	// 4. If no cases match, jump to default
	t.addAsm("    B %s", defaultLabel)

	// 5. Generate code for case bodies
	t.addAsm("    // --- Switch Case Bodies ---")
	for i, caseClause := range node.Cases {
		t.addAsm("%s:", caseLabels[i])
		for _, stmt := range caseClause.Body {
			stmt.Accept(t)
		}
		t.addAsm("    B %s", endSwitchLabel) // Jump to end after case body
	}

	// 6. Generate code for default body
	t.addAsm("%s:", defaultLabel)
	if node.Default != nil {
		for _, stmt := range node.Default.Body {
			stmt.Accept(t)
		}
	}

	// 7. End of switch
	t.addAsm("%s:", endSwitchLabel)
	t.releaseIntRegister(switchExprResult.Reg)

	// Pop the break label now that the switch is done.
	t.breakLabels = t.breakLabels[:len(t.breakLabels)-1]

	return nil
}

func (t *Translator) VisitCaseClause(node *ast.CaseClause) interface{} {
	// This is now handled entirely within VisitSwitchStmt.
	if t.DebugMode {
		fmt.Println("Translator.Visiting CaseClause (should be handled by SwitchStmt)")
	}
	return nil
}

func (t *Translator) VisitDefaultClause(node *ast.DefaultClause) interface{} {
	// This is now handled entirely within VisitSwitchStmt.
	if t.DebugMode {
		fmt.Println("Translator.Visiting DefaultClause (should be handled by SwitchStmt)")
	}
	return nil

}

func (t *Translator) VisitForStmt(node *ast.ForStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ForStmt")
	}

	loopStartLabel := t.newLabel("loop_start")
	loopBodyLabel := t.newLabel("loop_body")
	loopPostLabel := t.newLabel("loop_post") // For 'continue'
	loopEndLabel := t.newLabel("loop_end")   // For 'break'

	// Push labels onto the stacks for break/continue statements within this loop
	t.breakLabels = append(t.breakLabels, loopEndLabel)
	t.continueLabels = append(t.continueLabels, loopPostLabel)

	t.enterScope()

	// 1. Initialization
	if node.Init != nil {
		node.Init.Accept(t)
	}

	// 2. Loop Start / Condition Check
	t.addAsm("%s:", loopStartLabel)
	if node.Condition != nil {
		condResult := node.Condition.Accept(t).(ExpressionResult)
		t.addAsm("    // For loop condition check")
		t.addAsm("    CMP W%d, #0", condResult.Reg) // Check if the condition is false
		t.addAsm("    B.EQ %s", loopEndLabel)       // If false, exit loop
		t.releaseIntRegister(condResult.Reg)
	} else {
		// Infinite loop `for {}` - no condition, so we jump straight to the body
		// but we still need a way to get from the post-statement back here.
	}

	// Jump to body, this can be optimized out if body is next
	t.addAsm("    B %s", loopBodyLabel)

	// 4. Post-loop statement (for continue)
	t.addAsm("%s:", loopPostLabel)
	if node.Post != nil {
		node.Post.Accept(t)
	}
	t.addAsm("    B %s", loopStartLabel) // Jump back to the condition check

	// 3. Body of the loop
	t.addAsm("%s:", loopBodyLabel)
	node.Body.Accept(t)
	t.addAsm("    B %s", loopPostLabel) // After body, execute post-statement

	// 5. End of the loop
	t.addAsm("%s:", loopEndLabel)

	t.exitScope()

	// Pop labels from the stacks
	t.breakLabels = t.breakLabels[:len(t.breakLabels)-1]
	t.continueLabels = t.continueLabels[:len(t.continueLabels)-1]

	return nil
}

func (t *Translator) VisitReturnStmt(node *ast.ReturnStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ReturnStmt")
	}

	if node.Value != nil {
		result := node.Value.Accept(t)
		if result == nil {
			// This can happen for calls to functions that return void, like println
			// Check the function's return type. If it's void, this is okay.
			if t.currentFuncDef != nil && t.typeNodeToVarType(t.currentFuncDef.ReturnType) == TypeVoid {
				// This is a return from a void function, but with a value (e.g. return println()).
				// This should probably be a semantic error caught earlier, but for now, we just ignore the value.
			} else {
				panic("Return expression evaluated to nil for a non-void function")
			}
		} else {
			exprResult, ok := result.(ExpressionResult)
			if !ok {
				panic(fmt.Sprintf("Return expression did not evaluate to an ExpressionResult, but to %T", result))
			}

			// Move the result to the appropriate return register
			switch exprResult.Type {
			case TypeInt, TypeString, TypeBool:
				// ARM64 calling convention returns integer/pointer types in X0
				t.addAsm("    MOV X0, X%d", exprResult.Reg)
				t.releaseIntRegister(exprResult.Reg)
			case TypeFloat:
				// ARM64 calling convention returns float types in D0
				t.addAsm("    FMOV D0, D%d", exprResult.Reg)
				t.releaseFloatRegister(exprResult.Reg)
			default:
				// This includes TypeVoid, which shouldn't happen here because node.Value is not nil.
				panic(fmt.Sprintf("Unsupported return type: %s", exprResult.Type))
			}
		}
	}

	// After handling the return value (or if there is none), jump to the function's epilogue
	// to restore the stack and return properly.
	if t.currentFuncDef != nil {
		epilogueLabel := fmt.Sprintf(".L%s_epilogue", t.currentFuncDef.Name.Name)
		t.addAsm("    B " + epilogueLabel)
	} else {
		// This might happen in the global scope, which is an error.
		panic("Return statement outside of a function")
	}

	return nil
}

type ExpressionResult struct {
	Reg  int
	Type VarType
}

// Expressions

func (t *Translator) VisitTypeOfExpr(node *ast.TypeOfExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting TypeOfExpr")
	}
	// TODO: Implement TypeOfExpr translation
	return nil
}

func isRelationalOp(op string) bool {
	switch op {
	case ">", "<", ">=", "<=", "==", "!=", "&&", "||":
		return true
	default:
		return false
	}
}

// handleShortCircuit handles logical AND (&&) and OR (||) with short-circuiting.
func (t *Translator) handleShortCircuit(node *ast.BinaryExpr) interface{} {
	resultReg := t.acquireIntRegister()
	endLabel := t.newLabel(".L_logic_end")

	if node.Operator == "&&" {
		falseLabel := t.newLabel(".L_logic_false")

		// Evaluate left side
		leftResult, ok := node.Left.Accept(t).(ExpressionResult)
		if !ok || leftResult.Type != TypeBool {
			panic("Left side of && must be a boolean expression")
		}

		// If left is false (0), jump to set result to 0 and finish
		t.addAsm("    // Short-circuit AND: check left operand")
		t.addAsm("    CMP W%d, #0", leftResult.Reg)
		t.releaseIntRegister(leftResult.Reg)
		t.addAsm("    B.EQ %s", falseLabel)

		// Left was true, so evaluate right side
		rightResult, ok := node.Right.Accept(t).(ExpressionResult)
		if !ok || rightResult.Type != TypeBool {
			panic("Right side of && must be a boolean expression")
		}

		// The result of the expression is the result of the right side
		t.addAsm("    // Left was true, result is right operand")
		t.addAsm("    MOV W%d, W%d", resultReg, rightResult.Reg)
		t.releaseIntRegister(rightResult.Reg)
		t.addAsm("    B %s", endLabel)

		// False label: set result to 0
		t.addAsm("%s:", falseLabel)
		t.addAsm("    MOV W%d, #0", resultReg)

	} else { // "||"
		trueLabel := t.newLabel(".L_logic_true")

		// Evaluate left side
		leftResult, ok := node.Left.Accept(t).(ExpressionResult)
		if !ok || leftResult.Type != TypeBool {
			panic("Left side of || must be a boolean expression")
		}

		// If left is true (not 0), jump to set result to 1 and finish
		t.addAsm("    // Short-circuit OR: check left operand")
		t.addAsm("    CMP W%d, #0", leftResult.Reg)
		t.releaseIntRegister(leftResult.Reg)
		t.addAsm("    B.NE %s", trueLabel)

		// Left was false, so evaluate right side
		rightResult, ok := node.Right.Accept(t).(ExpressionResult)
		if !ok || rightResult.Type != TypeBool {
			panic("Right side of || must be a boolean expression")
		}

		// The result of the expression is the result of the right side
		t.addAsm("    // Left was false, result is right operand")
		t.addAsm("    MOV W%d, W%d", resultReg, rightResult.Reg)
		t.releaseIntRegister(rightResult.Reg)
		t.addAsm("    B %s", endLabel)

		// True label: set result to 1
		t.addAsm("%s:", trueLabel)
		t.addAsm("    MOV W%d, #1", resultReg)
	}

	t.addAsm("%s:", endLabel)
	return ExpressionResult{Reg: resultReg, Type: TypeBool}
}

// concatenateStrings generates assembly to concatenate two strings.
// It uses C standard library functions (strlen, malloc, strcpy, strcat).
func (t *Translator) concatenateStrings(left, right ExpressionResult) ExpressionResult {
	t.addAsm("    // --- String Concatenation ---")

	// Save the pointers to the two strings on the stack, as the registers will be reused.
	t.addAsm("    SUB SP, SP, #16")
	t.addAsm("    STP X%d, X%d, [SP]", left.Reg, right.Reg)
	t.releaseIntRegister(left.Reg)
	t.releaseIntRegister(right.Reg)

	// 1. Get length of left string (pointer is at [SP, #0])
	t.addAsm("    LDR X0, [SP, #0]")
	t.addAsm("    BL strlen")
	t.addAsm("    MOV X9, X0") // Save len1 in X9

	// 2. Get length of right string (pointer is at [SP, #8])
	t.addAsm("    LDR X0, [SP, #8]")
	t.addAsm("    BL strlen")
	t.addAsm("    MOV X10, X0") // Save len2 in X10

	// 3. Malloc new buffer: len1 + len2 + 1
	t.addAsm("    ADD X0, X9, X10")
	t.addAsm("    ADD X0, X0, #1")
	t.addAsm("    BL malloc")
	resultReg := t.acquireIntRegister()
	t.addAsm("    MOV X%d, X0", resultReg) // Save new buffer pointer in resultReg

	// 4. strcpy(new_buffer, left_string)
	t.addAsm("    LDR X1, [SP, #0]")       // 2nd arg: src (left_string pointer)
	t.addAsm("    MOV X0, X%d", resultReg) // 1st arg: dest (new_buffer pointer)
	t.addAsm("    BL strcpy")

	// 5. strcat(new_buffer, right_string)
	t.addAsm("    LDR X1, [SP, #8]")       // 2nd arg: src (right_string pointer)
	t.addAsm("    MOV X0, X%d", resultReg) // 1st arg: dest (new_buffer pointer)
	t.addAsm("    BL strcat")

	// 6. Clean up stack
	t.addAsm("    ADD SP, SP, #16")

	// 7. Return the new string
	return ExpressionResult{Reg: resultReg, Type: TypeString}
}

func (t *Translator) VisitBinaryExpr(node *ast.BinaryExpr) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting BinaryExpr: %s\n", node.Operator)
	}

	// Handle short-circuiting for logical operators
	if node.Operator == "&&" || node.Operator == "||" {
		return t.handleShortCircuit(node)
	}

	leftResult := node.Left.Accept(t).(ExpressionResult)
	rightResult := node.Right.Accept(t).(ExpressionResult)

	// Handle string concatenation
	if leftResult.Type == TypeString && rightResult.Type == TypeString {
		if node.Operator == "+" {
			t.needsStringHelpers = true
			return t.concatenateStrings(leftResult, rightResult)
		} else if node.Operator == "==" || node.Operator == "!=" {
			t.needsStrcmp = true
			resultReg := t.acquireIntRegister()

			// strcmp(s1, s2) -> arguments in X0, X1
			t.addAsm("    MOV X0, X%d", leftResult.Reg)
			t.addAsm("    MOV X1, X%d", rightResult.Reg)
			t.addAsm("    BL strcmp")

			// strcmp returns 0 if equal. Compare result in W0 with 0.
			t.addAsm("    CMP W0, #0")

			cond := ""
			if node.Operator == "==" {
				cond = "EQ" // Set if equal
			} else {
				cond = "NE" // Set if not equal
			}
			// Set result register to 1 if condition is met, 0 otherwise.
			t.addAsm("    CSET W%d, %s", resultReg, cond)

			t.releaseIntRegister(leftResult.Reg)
			t.releaseIntRegister(rightResult.Reg)

			return ExpressionResult{Reg: resultReg, Type: TypeBool}
		} else {
			panic(fmt.Sprintf("Unsupported operator '%s' for strings", node.Operator))
		}
	}

	// Handle type promotion: Int -> Float
	if leftResult.Type != rightResult.Type {
		if leftResult.Type == TypeInt && rightResult.Type == TypeFloat {
			t.addAsm("    // Promoting left operand from INT to FLOAT")
			promotedFloatReg := t.acquireFloatRegister()
			t.addAsm("    SCVTF D%d, W%d", promotedFloatReg, leftResult.Reg)
			t.releaseIntRegister(leftResult.Reg)
			leftResult = ExpressionResult{Reg: promotedFloatReg, Type: TypeFloat}
		} else if leftResult.Type == TypeFloat && rightResult.Type == TypeInt {
			t.addAsm("    // Promoting right operand from INT to FLOAT")
			promotedFloatReg := t.acquireFloatRegister()
			t.addAsm("    SCVTF D%d, W%d", promotedFloatReg, rightResult.Reg)
			t.releaseIntRegister(rightResult.Reg)
			rightResult = ExpressionResult{Reg: promotedFloatReg, Type: TypeFloat}
		} else {
			panic(fmt.Sprintf("Type mismatch in binary expression: %s and %s", leftResult.Type, rightResult.Type))
		}
	}

	// Perform the operation
	switch leftResult.Type {
	case TypeInt:
		resultReg := leftResult.Reg
		switch node.Operator {
		case "+":
			t.addAsm("    ADD X%d, X%d, X%d", resultReg, leftResult.Reg, rightResult.Reg)
		case "-":
			t.addAsm("    SUB X%d, X%d, X%d", resultReg, leftResult.Reg, rightResult.Reg)
		case "*":
			t.addAsm("    MUL X%d, X%d, X%d", resultReg, leftResult.Reg, rightResult.Reg)
		case "/":
			t.addAsm("    SDIV X%d, X%d, X%d", resultReg, leftResult.Reg, rightResult.Reg)
		case "%":
			// a % n = a - (a/n) * n
			divResultReg := t.acquireIntRegister()
			t.addAsm("    SDIV X%d, X%d, X%d", divResultReg, leftResult.Reg, rightResult.Reg) // divResultReg = a / n
			t.addAsm("    MUL X%d, X%d, X%d", divResultReg, divResultReg, rightResult.Reg)    // divResultReg = (a / n) * n
			t.addAsm("    SUB X%d, X%d, X%d", resultReg, leftResult.Reg, divResultReg)        // resultReg = a - divResultReg
			t.releaseIntRegister(divResultReg)
		case "==", "!=", ">", "<", ">=", "<=":
			t.addAsm("    CMP W%d, W%d", leftResult.Reg, rightResult.Reg)
			cond := ""
			switch node.Operator {
			case "==":
				cond = "EQ"
			case "!=":
				cond = "NE"
			case ">":
				cond = "GT"
			case "<":
				cond = "LT"
			case ">=":
				cond = "GE"
			case "<=":
				cond = "LE"
			}
			t.addAsm("    CSET W%d, %s", resultReg, cond)
			t.releaseIntRegister(rightResult.Reg)
			return ExpressionResult{Reg: resultReg, Type: TypeBool}
		default:
			panic(fmt.Sprintf("Unsupported integer operator: %s", node.Operator))
		}
		// The result is in resultReg. The right register can be freed.
		if rightResult.Type == TypeFloat {
			t.releaseFloatRegister(rightResult.Reg)
		} else {
			// Default to releasing an integer register for Int, Bool, etc.
			t.releaseIntRegister(rightResult.Reg)
		}
		// For arithmetic operators, the type is Int. For relational, it's Bool.
		resultType := TypeInt
		if isRelationalOp(node.Operator) {
			resultType = TypeBool
		}
		return ExpressionResult{Reg: resultReg, Type: resultType}
	case TypeBool:
		resultReg := leftResult.Reg
		switch node.Operator {
		case "==", "!=":
			t.addAsm("    CMP W%d, W%d", leftResult.Reg, rightResult.Reg)
			cond := ""
			if node.Operator == "==" {
				cond = "EQ"
			} else {
				cond = "NE"
			}
			t.addAsm("    CSET W%d, %s", resultReg, cond)
			t.releaseIntRegister(rightResult.Reg)
			return ExpressionResult{Reg: resultReg, Type: TypeBool}
		default:
			panic(fmt.Sprintf("Unsupported boolean operator: %s. (&& and || are short-circuited)", node.Operator))
		}
	case TypeFloat:
		reg := leftResult.Reg
		switch node.Operator {
		case "+":
			t.addAsm("    FADD D%d, D%d, D%d", reg, leftResult.Reg, rightResult.Reg)
		case "-":
			t.addAsm("    FSUB D%d, D%d, D%d", reg, leftResult.Reg, rightResult.Reg)
		case "*":
			t.addAsm("    FMUL D%d, D%d, D%d", reg, leftResult.Reg, rightResult.Reg)
		case "/":
			t.addAsm("    FDIV D%d, D%d, D%d", reg, leftResult.Reg, rightResult.Reg)
		case "%":
			// fmod(x, y) = x - trunc(x/y) * y
			tmp1 := t.acquireFloatRegister()
			tmp2 := t.acquireFloatRegister()
			t.addAsm("    // Calculating float modulo (fmod)")
			t.addAsm("    FDIV D%d, D%d, D%d", tmp1, leftResult.Reg, rightResult.Reg) // tmp1 = x/y
			t.addAsm("    FRINTZ D%d, D%d", tmp2, tmp1)                               // tmp2 = trunc(x/y)
			t.addAsm("    FMUL D%d, D%d, D%d", tmp1, tmp2, rightResult.Reg)           // tmp1 = trunc(x/y) * y
			t.addAsm("    FSUB D%d, D%d, D%d", reg, leftResult.Reg, tmp1)             // reg = x - tmp1
			t.releaseFloatRegister(tmp1)
			t.releaseFloatRegister(tmp2)
		case "==", "!=", ">", "<", ">=", "<=":
			resultReg := t.acquireIntRegister()
			t.addAsm("    FCMP D%d, D%d", leftResult.Reg, rightResult.Reg)
			cond := ""
			switch node.Operator {
			case "==":
				cond = "EQ"
			case "!=":
				cond = "NE"
			case ">":
				cond = "GT"
			case "<":
				cond = "LT"
			case ">=":
				cond = "GE"
			case "<=":
				cond = "LE"
			}
			t.addAsm("    CSET W%d, %s", resultReg, cond)
			t.releaseFloatRegister(leftResult.Reg)
			t.releaseFloatRegister(rightResult.Reg)
			return ExpressionResult{Reg: resultReg, Type: TypeBool}
		default:
			panic(fmt.Sprintf("Unsupported float operator: %s", node.Operator))
		}
		t.releaseFloatRegister(rightResult.Reg)
		return ExpressionResult{Reg: reg, Type: TypeFloat}
	default:
		panic(fmt.Sprintf("Unsupported type in binary expression: %s", leftResult.Type))
	}
}

func (t *Translator) VisitUnaryExpr(node *ast.UnaryExpr) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting UnaryExpr: %s\n", node.Operator)
	}

	operandResult := node.Right.Accept(t).(ExpressionResult)

	switch node.Operator {
	case "!":
		if operandResult.Type != TypeBool {
			panic(fmt.Sprintf("Unsupported type for unary ! operator: %s", operandResult.Type))
		}
		// EOR with 1 flips the bit (0->1, 1->0)
		t.addAsm("    EOR W%d, W%d, #1", operandResult.Reg, operandResult.Reg)
		return operandResult
	case "-":
		switch operandResult.Type {
		case TypeInt:
			// NEG instruction negates the value in a register.
			// It's a two-operand instruction, so we can use the same register for source and destination.
			t.addAsm("    NEG X%d, X%d", operandResult.Reg, operandResult.Reg)
			return operandResult // The result is in the same register, with the same type.
		case TypeFloat:
			// FNEG for floating-point negation.
			t.addAsm("    FNEG D%d, D%d", operandResult.Reg, operandResult.Reg)
			return operandResult // The result is in the same register, with the same type.
		default:
			panic(fmt.Sprintf("Unsupported type for unary minus operator: %s", operandResult.Type))
		}
	default:
		panic(fmt.Sprintf("Unsupported unary operator: %s", node.Operator))
	}
}

func (t *Translator) VisitParenExpr(node *ast.ParenExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ParenExpr")
	}
	// Just evaluate the inner expression and return its result
	return node.Expression.Accept(t)
}

func (t *Translator) VisitIdentifierExpr(node *ast.IdentifierExpr) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting IdentifierExpr: %s\n", node.Name)
	}

	mangledName, varType, exists := t.lookupSymbol(node.Name)
	if !exists {
		panic(fmt.Sprintf("Undefined variable: %s", node.Name))
	}

	// Use a temporary register to load the address of the variable from the data section.
	// This is more robust and avoids PC-relative range issues.
	addrReg := t.acquireIntRegister()
	t.addAsm("    LDR X%d, =%s", addrReg, mangledName)

	switch varType {
	case TypeInt:
		valReg := t.acquireIntRegister()
		// Load 32-bit signed word from the address in addrReg.
		t.addAsm("    LDRSW X%d, [X%d]", valReg, addrReg)
		t.releaseIntRegister(addrReg) // Free the address register.
		return ExpressionResult{Reg: valReg, Type: TypeInt}
	case TypeFloat:
		valReg := t.acquireFloatRegister()
		// Load 64-bit double from the address in addrReg.
		t.addAsm("    LDR D%d, [X%d]", valReg, addrReg)
		t.releaseIntRegister(addrReg) // Free the address register.
		return ExpressionResult{Reg: valReg, Type: TypeFloat}
	case TypeString:
		valReg := t.acquireIntRegister()
		// For strings, we load the pointer to the string data.
		t.addAsm("    LDR X%d, [X%d]", valReg, addrReg)
		t.releaseIntRegister(addrReg) // Free the address register.
		return ExpressionResult{Reg: valReg, Type: TypeString}
	case TypeBool:
		valReg := t.acquireIntRegister()
		// Load a byte from the address in addrReg.
		t.addAsm("    LDRB W%d, [X%d]", valReg, addrReg)
		t.releaseIntRegister(addrReg) // Free the address register.
		return ExpressionResult{Reg: valReg, Type: TypeBool}
	case TypeSlice:
		valReg := t.acquireIntRegister()
		// For slices, we load the pointer to the slice descriptor.
		t.addAsm("    LDR X%d, [X%d]", valReg, addrReg)
		t.releaseIntRegister(addrReg)
		return ExpressionResult{Reg: valReg, Type: TypeSlice}
	default:
		t.releaseIntRegister(addrReg) // Release register even on panic
		panic(fmt.Sprintf("Loading for type %s not implemented for identifier '%s'", varType, node.Name))
	}
}

func (t *Translator) VisitTypeConversionExpr(node *ast.TypeConversionExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting TypeConversionExpr")
	}
	// TODO: Implement TypeConversionExpr translation
	return nil
}

func (t *Translator) VisitCompositeLiteralExpr(node *ast.CompositeLiteralExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting CompositeLiteralExpr")
	}

	// For now, we only support slice literals, e.g., []int{1, 2, 3}
	sliceType, ok := node.Type.(*ast.SliceTypeNode)
	if !ok {
		panic("Currently only slice composite literals are supported")
	}

	// Determine the element type
	elementType := t.typeNodeToVarType(sliceType.ElementType)
	if elementType != TypeInt {
		panic("Only slices of integers are currently supported")
	}

	// Create a label for the static data
	dataLabel := t.newLabel("slice_data")

	// Emit the slice data to the .data section
	var elementValues []string
	for _, element := range node.Elements {
		lit, ok := element.Value.(*ast.IntegerLiteral)
		if !ok {
			panic("Slice elements must be integer literals for now")
		}
		// lit.Value is a string representation of the number, so use it directly.
		elementValues = append(elementValues, lit.Value)
	}
	dataValues := strings.Join(elementValues, ", ")
	t.addData(fmt.Sprintf("%s: .quad %s", dataLabel, dataValues))

	// Now, we need to create a slice descriptor in memory.
	// A slice is a struct { data_ptr, len, cap }.
	// We'll allocate this descriptor on the stack or in the data section.
	// For a simple static initializer, data section is fine.

	sliceStructLabel := t.newLabel("slice_descriptor")
	t.addData(fmt.Sprintf("%s:", sliceStructLabel))
	t.addData(fmt.Sprintf("    .quad %s  // Pointer to data", dataLabel))
	t.addData(fmt.Sprintf("    .quad %d    // Length", len(node.Elements)))
	t.addData(fmt.Sprintf("    .quad %d    // Capacity", len(node.Elements)))

	// Load the address of the slice descriptor into a register
	reg := t.acquireIntRegister()
	t.addAsm("    LDR X%d, =%s", reg, sliceStructLabel)

	return ExpressionResult{
		Reg:  reg,
		Type: TypeSlice, // Special type for the slice descriptor pointer
	}
}

func (t *Translator) VisitCompositeElement(node *ast.CompositeElement) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting CompositeElement")
	}
	// TODO: Implement CompositeElement translation
	return nil
}

func (t *Translator) VisitIndexAccessExpr(node *ast.IndexAccessExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting IndexAccessExpr")
	}
	// TODO: Implement IndexAccessExpr translation
	return nil
}

func (t *Translator) VisitFieldAccessExpr(node *ast.FieldAccessExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting FieldAccessExpr")
	}
	// TODO: Implement FieldAccessExpr translation
	return nil
}

func (t *Translator) VisitCallExpr(node *ast.CallExpr) interface{} {
	if t.DebugMode {
		if ident, ok := node.Function.(*ast.IdentifierExpr); ok {
			fmt.Printf("Translator.VisitCallExpr: Visiting call to identifier: '%s'\n", ident.Name)
		} else {
			fmt.Println("Translator.VisitCallExpr: Visiting complex call expression")
		}
	}

	ident, ok := node.Function.(*ast.IdentifierExpr)
	if !ok {
		// For now, we only handle direct function calls like `myFunc()`
		panic("Unhandled function call type")
	}
	funcName := ident.Name

	// Handle special built-in functions
	switch funcName {
	case "print", "println":
		t.needsPrintf = true
		t.addAsm("    // --- Start of %s call ---", funcName)
		t.handlePrintln(node.Arguments)
		t.addAsm("    // --- End of %s call ---", funcName)
		return nil // println does not return a value
	case "Atoi":
		return t.handleAtoi(node.Arguments)
	case "parseFloat":
		return t.handleParseFloat(node.Arguments)
	case "typeof":
		return t.handleTypeOf(node.Arguments)
	}

	// ... (rest of the code remains the same)
	// --- General Function Call for user-defined functions ---
	funcDef, exists := t.funcSyms[funcName]
	if !exists {
		panic(fmt.Sprintf("Call to undefined function: %s", funcName))
	}

	// Check if the number of arguments matches the function definition
	if len(node.Arguments) != len(funcDef.Parameters) {
		panic(fmt.Sprintf("Incorrect number of arguments for function %s: expected %d, got %d", funcName, len(funcDef.Parameters), len(node.Arguments)))
	}

	// --- Argument Passing ---
	// Per ARM64 calling convention, first 8 integer/pointer args are in X0-X7,
	// and first 8 float args are in D0-D7.
	t.addAsm("    // --- Start of call to %s ---", funcName)

	intArgReg := 0
	floatArgReg := 0

	for i, arg := range node.Arguments {
		paramType := t.typeNodeToVarType(funcDef.Parameters[i].Type)
		argResult := arg.Accept(t).(ExpressionResult)

		switch paramType {
		case TypeInt, TypeBool, TypeString: // Strings are pointers
			if intArgReg < 8 {
				t.addAsm("    MOV X%d, X%d  // Move arg %d for %s", intArgReg, argResult.Reg, i, funcName)
				intArgReg++
			} else {
				// TODO: Handle stack-passed arguments
			}
		case TypeFloat:
			if floatArgReg < 8 {
				t.addAsm("    FMOV D%d, D%d // Move arg %d for %s", floatArgReg, argResult.Reg, i, funcName)
				floatArgReg++
			} else {
				// TODO: Handle stack-passed float arguments
			}
		}
		// Release the register used by the argument expression now that it's moved
		if argResult.Type == TypeFloat {
			t.releaseFloatRegister(argResult.Reg)
		} else {
			t.releaseIntRegister(argResult.Reg)
		}
	}

	t.addAsm("    BL %s", funcName)
	t.addAsm("    // --- End of call to %s ---", funcName)

	// --- Return Value ---
	// The return value will be in X0 (for int/ptr) or D0 (for float).
	// We need to move it to a temporary register from our pool.
	returnType := t.typeNodeToVarType(funcDef.ReturnType)
	if returnType != TypeVoid {
		var resultReg int
		if returnType == TypeFloat {
			resultReg = t.acquireFloatRegister()
			t.addAsm("    FMOV D%d, D0", resultReg)
		} else {
			resultReg = t.acquireIntRegister()
			t.addAsm("    MOV X%d, X0", resultReg)
		}
		return ExpressionResult{Reg: resultReg, Type: returnType}
	}

	return nil // No value for void functions
}

// typeNodeToVarType converts an ast.TypeNode to a VarType.
func (t *Translator) typeNodeToVarType(typeNode ast.TypeNode) VarType {
	if typeNode == nil {
		// This typically means a void return type for a function.
		return TypeVoid
	}

	switch n := typeNode.(type) {
	case *ast.PrimitiveTypeNode:
		switch n.Kind {
		case ast.IntKind:
			return TypeInt
		case ast.Float64Kind:
			return TypeFloat
		case ast.StringKind:
			return TypeString
		case ast.BoolKind:
			return TypeBool
		default:
			panic(fmt.Sprintf("Unsupported primitive type kind: %v", n.Kind))
		}
	// TODO: Add cases for other types like Array, Struct, etc. as they are implemented.
	case *ast.TypeName:
		switch n.Name {
		case "int":
			return TypeInt
		case "f64":
			return TypeFloat
		case "string":
			return TypeString
		case "bool":
			return TypeBool
		case "void":
			return TypeVoid
		default:
			panic(fmt.Sprintf("Unsupported type name: %s", n.Name))
		}
	default:
		panic(fmt.Sprintf("Unsupported type node: %T", n))
	}
}

func (t *Translator) handleTypeOf(args []ast.Expression) interface{} {
	if len(args) != 1 {
		panic("typeof() expects exactly one argument")
	}

	// Evaluate the expression to find its type.
	result := args[0].Accept(t).(ExpressionResult)

	// The result register isn't needed, just its type. Release the register.
	switch result.Type {
	case TypeInt, TypeString, TypeBool, TypeSlice:
		t.releaseIntRegister(result.Reg)
	case TypeFloat:
		t.releaseFloatRegister(result.Reg)
		// Note: Slices and other complex types might need special handling for register release.
		// For now, we assume simple types or types whose registers can be released.
	}

	// Get the string representation of the type, e.g., "int", "float64".
	typeString := result.Type.String()

	// Add the type string to the .data section.
	label := t.addStringData(typeString)

	// Load the address of the string into a new register.
	reg := t.acquireIntRegister()
	t.addAsm("    LDR X%d, =%s", reg, label)

	// Return the result, which is a pointer to the string.
	return ExpressionResult{Reg: reg, Type: TypeString}
}

func (t *Translator) handleParseFloat(args []ast.Expression) interface{} {
	if len(args) != 1 {
		panic("parseFloat expects exactly one argument")
	}

	argResult := args[0].Accept(t)
	if argResult == nil {
		panic("Argument to parseFloat is nil")
	}

	res, ok := argResult.(ExpressionResult)
	if !ok {
		panic(fmt.Sprintf("Unexpected result type from argument expression: %T", argResult))
	}

	if res.Type != TypeString {
		panic("parseFloat expects a string argument")
	}

	t.needsAtof = true

	// The register from res.Reg holds the address of the string.
	// Per ARM64 calling convention, the first argument to a function is in X0.
	t.addAsm("    MOV X0, X%d", res.Reg)
	t.releaseIntRegister(res.Reg) // Release the register that held the string address

	// Call atof
	t.addAsm("    BL atof")

	// The result is in D0. We move it to a new temporary register.
	resultReg := t.acquireFloatRegister()
	t.addAsm("    FMOV D%d, D0", resultReg)

	return ExpressionResult{Reg: resultReg, Type: TypeFloat}
}

func (t *Translator) handleAtoi(args []ast.Expression) interface{} {
	if len(args) != 1 {
		panic("Atoi expects exactly one argument")
	}

	argResult := args[0].Accept(t)
	if argResult == nil {
		panic("Argument to Atoi is nil")
	}

	res, ok := argResult.(ExpressionResult)
	if !ok {
		panic(fmt.Sprintf("Unexpected result type from argument expression: %T", argResult))
	}

	if res.Type != TypeString {
		panic("Atoi expects a string argument")
	}

	t.needsAtoi = true

	// The register from res.Reg holds the address of the string.
	// Per ARM64 calling convention, the first argument to a function is in X0.
	t.addAsm("    MOV X0, X%d", res.Reg)
	t.releaseIntRegister(res.Reg) // Release the register that held the string address

	// Call atoi
	t.addAsm("    BL atoi")

	// The result is in X0. We move it to a new temporary register.
	resultReg := t.acquireIntRegister()
	t.addAsm("    MOV X%d, X0", resultReg)

	return ExpressionResult{Reg: resultReg, Type: TypeInt}
}

func (t *Translator) handlePrintln(args []ast.Expression) interface{} {
	t.needsPrintf = true
	// Label for the space character, created only if needed.
	var spaceLabel string

	for i, arg := range args {
		// Print a space before all but the first argument.
		if i > 0 {
			if spaceLabel == "" {
				spaceLabel = t.addStringData(" ")
			}
			t.addAsm("    LDR X0, =%s", spaceLabel)
			t.addAsm("    BL printf")
		}

		result := arg.Accept(t).(ExpressionResult)

		if result.Type == TypeSlice {
			// This assumes the slice is of integers for now.
			t.ensurePrintIntSliceHelper()
			t.addAsm("    MOV X0, X%d", result.Reg) // Pass slice descriptor address to our helper
			t.addAsm("    BL print_int_slice")
			t.releaseIntRegister(result.Reg)
		} else {
			// It's a regular type, print it with printf
			var format string
			switch result.Type {
			case TypeInt:
				format = "%d"
			case TypeString:
				format = "%s"
			case TypeFloat:
				format = "%f"
			case TypeBool:
				format = "%s" // Bool is handled by passing "true" or "false" string address
			default:
				panic(fmt.Sprintf("Unsupported type for println: %s", result.Type))
			}

			formatLabel := t.addStringData(format)
			t.addAsm("    LDR X0, =%s", formatLabel)

			switch result.Type {
			case TypeInt, TypeString:
				t.addAsm("    MOV X1, X%d", result.Reg)
				t.releaseIntRegister(result.Reg)
			case TypeFloat:
				t.addAsm("    FMOV D0, D%d", result.Reg)
				t.releaseFloatRegister(result.Reg)
			case TypeBool:
				trueLabelReg := t.acquireIntRegister()
				falseLabelReg := t.acquireIntRegister()

				if t.trueStrLabel == "" {
					t.trueStrLabel = t.addStringData("true")
				}
				if t.falseStrLabel == "" {
					t.falseStrLabel = t.addStringData("false")
				}

				t.addAsm("    LDR X%d, =%s", trueLabelReg, t.trueStrLabel)
				t.addAsm("    LDR X%d, =%s", falseLabelReg, t.falseStrLabel)

				t.addAsm("    CMP X%d, #0", result.Reg) // Compare the bool value (0 or 1)
				// Select correct string address into X1. If NE (not equal to 0, i.e., true), use true label.
				t.addAsm("    CSEL X1, X%d, X%d, NE", trueLabelReg, falseLabelReg)

				t.releaseIntRegister(trueLabelReg)
				t.releaseIntRegister(falseLabelReg)
				t.releaseIntRegister(result.Reg)
			}
			t.addAsm("    BL printf")
		}
	}

	// Finally, print a newline character.
	newlineLabel := t.addStringData("\n")
	t.addAsm("    LDR X0, =%s", newlineLabel)
	t.addAsm("    BL printf")

	return nil
}

func (t *Translator) VisitIntegerLiteral(node *ast.IntegerLiteral) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting IntegerLiteral: %s\n", node.Value)
	}
	reg := t.acquireIntRegister()
	t.addAsm("    MOV X%d, #%s", reg, node.Value)
	return ExpressionResult{Reg: reg, Type: TypeInt}
}

func (t *Translator) VisitFloatLiteral(node *ast.FloatLiteral) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting FloatLiteral: %s\n", node.Value)
	}
	label := t.addFloatData(node.Value)
	reg := t.acquireFloatRegister()
	t.addAsm("    LDR D%d, %s", reg, label)
	return ExpressionResult{Reg: reg, Type: TypeFloat}
}

func (t *Translator) VisitStringLiteral(node *ast.StringLiteral) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting StringLiteral: %s\n", node.Value)
	}
	// Add the string literal to the .data section
	label := t.addStringData(node.Value)

	// Load the address of the string into a register
	reg := t.acquireIntRegister()
	t.addAsm("    LDR X%d, =%s", reg, label)

	// Return an ExpressionResult
	return ExpressionResult{Reg: reg, Type: TypeString}
}

func (t *Translator) VisitCharLiteral(node *ast.CharLiteral) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting CharLiteral: %s\n", node.Value)
	}

	// The parser provides the character itself (e.g., 'a' becomes string "a").
	// We just need to get the integer value of the first rune.
	var runeValue rune
	if len(node.Value) > 0 {
		runeValue = []rune(node.Value)[0]
	} else {
		panic("empty character literal")
	}

	reg := t.acquireIntRegister()
	t.addAsm("    MOV X%d, #%d", reg, runeValue)
	return ExpressionResult{Reg: reg, Type: TypeInt}
}

func (t *Translator) VisitBoolLiteral(node *ast.BoolLiteral) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting BoolLiteral: %v\n", node.Value)
	}
	reg := t.acquireIntRegister()
	value := 0
	if node.Value {
		value = 1
	}
	t.addAsm("    MOV X%d, #%d", reg, value)
	return ExpressionResult{Reg: reg, Type: TypeBool}
}

func (t *Translator) VisitNilLiteral(node *ast.NilLiteral) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting NilLiteral")
	}
	// TODO: Implement NilLiteral translation
	return nil
}

func (t *Translator) VisitArrayLiteral(node *ast.ArrayLiteral) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ArrayLiteral")
	}
	// TODO: Implement ArrayLiteral translation
	for _, el := range node.Elements {
		el.Accept(t)
	}
	return nil
}

func (t *Translator) VisitMapLiteral(node *ast.MapLiteral) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting MapLiteral")
	}
	// TODO: Implement MapLiteral translation
	for _, entry := range node.Entries {
		entry.Accept(t)
	}
	return nil
}

func (t *Translator) VisitMapEntry(node *ast.MapEntry) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting MapEntry")
	}
	// TODO: Implement MapEntry translation
	node.Key.Accept(t)
	node.Value.Accept(t)
	return nil
}

// Types
func (t *Translator) VisitTypeName(node *ast.TypeName) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting TypeName")
	}
	// TODO: Implement TypeName translation (usually for type checking or metadata)
	return nil
}

func (t *Translator) VisitArrayTypeNode(node *ast.ArrayTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ArrayTypeNode")
	}
	// TODO: Implement ArrayTypeNode translation
	node.ElementType.Accept(t)
	if node.Size != nil {
		// node.Size is an Expression, typically an IntegerLiteral
		node.Size.Accept(t)
	}
	return nil
}

func (t *Translator) VisitSliceTypeNode(node *ast.SliceTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting SliceTypeNode")
	}
	// TODO: Implement SliceTypeNode translation
	node.ElementType.Accept(t)
	return nil
}

func (t *Translator) VisitMapTypeNode(node *ast.MapTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting MapTypeNode")
	}
	// TODO: Implement MapTypeNode translation
	node.KeyType.Accept(t)
	node.ValueType.Accept(t)
	return nil
}

func (t *Translator) VisitPointerTypeNode(node *ast.PointerTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting PointerTypeNode")
	}
	// TODO: Implement PointerTypeNode translation
	node.ElementType.Accept(t)
	return nil
}

func (t *Translator) VisitFunctionTypeNode(node *ast.FunctionTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting FunctionTypeNode")
	}
	// TODO: Implement FunctionTypeNode translation
	for _, paramType := range node.ParameterTypes {
		if paramType != nil {
			paramType.Accept(t)
		}
	}
	if node.ReturnType != nil {
		node.ReturnType.Accept(t)
	}
	return nil
}

func (t *Translator) VisitOptionalTypeNode(node *ast.OptionalTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting OptionalTypeNode")
	}
	// TODO: Implement OptionalTypeNode translation
	node.ElementType.Accept(t)
	return nil
}

func (t *Translator) VisitAnonymousStructTypeNode(node *ast.AnonymousStructTypeNode) interface{} {
	if t.DebugMode {
		fmt.Printf("Visiting AnonymousStructTypeNode: %v\n", node)
	}
	// For now, we don't generate specific code for anonymous struct type definitions themselves,
	// as they are mainly for type checking and structure definition.
	// The actual instantiation is handled by CompositeLiteralExpr.
	return nil
}

func (t *Translator) VisitAnonymousInterfaceTypeNode(node *ast.AnonymousInterfaceTypeNode) interface{} {
	if t.DebugMode {
		fmt.Printf("Visiting AnonymousInterfaceTypeNode: %v\n", node)
	}
	// Placeholder implementation
	return nil
}

func (t *Translator) VisitMethodSignatureNode(node *ast.MethodSignatureNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting MethodSignatureNode")
	}
	// TODO: Implement MethodSignatureNode translation
	for _, param := range node.Parameters {
		if param != nil {
			param.Accept(t) // param is ast.ParameterDecl
		}
	}
	if node.ReturnType != nil {
		node.ReturnType.Accept(t) // ret is ast.TypeNode
	}
	return nil
}

func (t *Translator) VisitPrimitiveTypeNode(node *ast.PrimitiveTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting PrimitiveTypeNode")
	}
	// TODO: Implement PrimitiveTypeNode translation
	return nil
}
