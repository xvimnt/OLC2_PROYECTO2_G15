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
	TypeArray
	TypeStruct
)

func (v VarType) String() string {
	switch v {
	case TypeInt:
		return "Int"
	case TypeFloat:
		return "Float"
	case TypeString:
		return "String"
	case TypeBool:
		return "Bool"
	case TypeVoid:
		return "Void"
	case TypeArray:
		return "Array"
	case TypeStruct:
		return "Struct"
	default:
		return "Unknown"
	}
}

// Translator translates AST nodes into assembly code.
type Translator struct {
	asm                []string          // Stores generated .text section assembly lines
	dataSection        []string          // Stores generated .data section assembly lines
	symbolTables       []map[string]string // Stack of maps: original name -> mangled name
	varInfo            map[string]VarType  // Map: mangled name -> type
	scopeCounter       int
	scopeIDStack       []int
	stringCounter      int               // For generating unique string labels
	labelCounter       int               // For generating unique labels
	needsPrintf        bool              // Tracks if printf is used (for .extern printf)
	hasIntFormatStr    bool              // Tracks if the integer format string has been added
	hasFloatFormatStr  bool              // Tracks if the float format string has been added
	hasStringFormatStr bool              // Tracks if the string format string has been added
	currentFuncDef     *ast.FunctionDecl // Keep track of the current function being defined
	DebugMode          bool
	needsStringHelpers bool              // Tracks if string concatenation helpers are needed

	// Register allocation
	intRegs   []bool // Availability of general-purpose integer registers (X9-X15)
	floatRegs []bool // Availability of general-purpose float registers (D8-D15)
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
		hasIntFormatStr:    false,
		hasFloatFormatStr:  false,
		hasStringFormatStr: false,
		currentFuncDef:     nil,
		DebugMode:          debugMode,
		intRegs:            make([]bool, numIntRegs),
		floatRegs:          make([]bool, numFloatRegs),
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
	t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .asciz %q", label, strContent))
	t.stringCounter++
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
		if t.needsPrintf {
			finalAsm = append(finalAsm, ".extern printf")
		}
		finalAsm = append(finalAsm, ".text")
		finalAsm = append(finalAsm, t.asm...)
	}

	return finalAsm
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
	// Global directives and .data section will be prepended by GetAssembly().
	// Here, we just process declarations.
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

	// TODO: Process parameters and add them to the symbol table

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
		t.addAsm(".L%s_epilogue:", node.Name.Name) // Label for potential jumps to epilogue
	} else {
		t.addAsm(".L_anonymous_func_%d_epilogue:", t.stringCounter-1) // Match potential anonymous label
	}
	t.addAsm("    MOV W0, #0")              // Default return code 0 for other functions
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
		fmt.Println("Translator.Visiting VarDecl")
	}
	if node.Name == nil {
		return nil // Should not happen in a valid program
	}
	varName := node.Name.Name

	// Handle initializer
	if node.Initializer != nil {
		switch init := node.Initializer.(type) {
		case *ast.IntegerLiteral:
			// For global integers, we define them in the data section and store their type.
			mangledName := t.defineSymbol(varName, TypeInt)
			t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word %s", mangledName, init.Value))
		case *ast.UnaryExpr:
			// Handle unary expressions, e.g., negative numbers
			if init.Operator == "-" {
				if intLit, ok := init.Right.(*ast.IntegerLiteral); ok {
					mangledName := t.defineSymbol(varName, TypeInt)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word -%s", mangledName, intLit.Value))
				} else {
					// Unhandled unary expression operand, default to 0
					mangledName := t.defineSymbol(varName, TypeUnknown)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
				}
			} else {
				// Unhandled unary operator, default to 0
				mangledName := t.defineSymbol(varName, TypeUnknown)
				t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
			}
		case *ast.StringLiteral:
			// For global strings, we store the string, and the variable holds its address.
			strLabel := t.addStringData(init.Value)
			mangledName := t.defineSymbol(varName, TypeString)
			t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .quad %s", mangledName, strLabel))
		case *ast.FloatLiteral:
			// For global floats, we define them in the data section and store their type.
			mangledName := t.defineSymbol(varName, TypeFloat)
			t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .double %s", mangledName, init.Value))
		case *ast.BoolLiteral:
			// For global booleans, we define them as a byte.
			val := "0"
			if init.Value {
				val = "1"
			}
			mangledName := t.defineSymbol(varName, TypeBool)
			t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .byte %s", mangledName, val))
		default:
			// Unhandled initializer type, default to 0
			mangledName := t.defineSymbol(varName, TypeUnknown)
			t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
		}
	} else {
		// Uninitialized variable, assign default value based on type
		if node.ExplicitType != nil {
			if typeName, ok := node.ExplicitType.(*ast.TypeName); ok {
				switch typeName.Name {
				case "int":
					mangledName := t.defineSymbol(varName, TypeInt)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
				case "float64":
					mangledName := t.defineSymbol(varName, TypeFloat)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .double 0.0", mangledName))
				case "string":
					strLabel := t.addStringData("")
					mangledName := t.defineSymbol(varName, TypeString)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .quad %s", mangledName, strLabel))
				case "bool":
					mangledName := t.defineSymbol(varName, TypeBool)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .byte 0", mangledName)) // 0 for false
				default:
					// Default for unhandled explicit types
					mangledName := t.defineSymbol(varName, TypeUnknown)
					t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
				}
			} else {
				// Type is not a simple TypeName, default for now
				mangledName := t.defineSymbol(varName, TypeUnknown)
				t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
			}
		} else {
			// Type not specified (e.g. from := which requires an initializer), so this case is for declarations without initializer.
			// Default to integer 0.
			mangledName := t.defineSymbol(varName, TypeUnknown)
			t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word 0", mangledName))
		}
	}

	return nil
}

// ...
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
		fmt.Println("Translator.Visiting AssignStmt")
	}

	// Get the variable name from the left side (LHS)
	var varName string
	if ident, ok := node.Left.(*ast.IdentifierExpr); ok {
		varName = ident.Name
	} else {
		// This is a simplification. Real-world scenarios would handle struct fields, array elements, etc.
		fmt.Fprintf(os.Stderr, "Unsupported L-value in assignment: %T\n", node.Left)
		return nil
	}

	// We need to know the type of the variable to use the correct store instruction.
	mangledName, varType, typeExists := t.lookupSymbol(varName)
	if !typeExists {
		fmt.Fprintf(os.Stderr, "Assignment to undeclared variable: %s\n", varName)
		return nil
	}

	// Handle different assignment operators
	switch node.Operator {
	case "=":
		// Evaluate the right side (RHS) and generate code to store the value.
		// This is a simplified evaluation that handles literals directly.
		// A more robust implementation would have expression visitors return results in registers.
		switch rhs := node.Right.(type) {
		case *ast.BoolLiteral:
			if varType != TypeBool {
				fmt.Fprintf(os.Stderr, "Type mismatch in assignment to %s. Expected Bool.\n", varName)
				return nil
			}
			val := "0"
			if rhs.Value {
				val = "1"
			}
			t.addAsm("    // --- Start of assignment to %s ---", varName)
			t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
			t.addAsm("    MOV W11, #%s", val)        // Load immediate value (0 or 1)
			t.addAsm("    STRB W11, [X10]")         // Store byte value
			t.addAsm("    // --- End of assignment to %s ---", varName)
			t.addAsm("") // Add a blank line for readability after assignment

		case *ast.IntegerLiteral:
			if varType != TypeInt {
				fmt.Fprintf(os.Stderr, "Type mismatch in assignment to %s. Expected Int.\n", varName)
				return nil
			}
			t.addAsm("    // --- Start of assignment to %s ---", varName)
			t.addAsm("    LDR X10, =%s", mangledName)      // Load address of the variable
			t.addAsm("    MOV W11, #%s", rhs.Value)    // Load immediate integer value
			t.addAsm("    STR W11, [X10]")                // Store word value
			t.addAsm("    // --- End of assignment to %s ---", varName)
			t.addAsm("") // Add a blank line for readability after assignment

		case *ast.IdentifierExpr:
			rhsVarName := rhs.Name
			rhsMangledName, rhsVarType, rhsExists := t.lookupSymbol(rhsVarName)
			if !rhsExists {
				fmt.Fprintf(os.Stderr, "Assignment from undeclared variable: %s\n", rhsVarName)
				return nil
			}

			// Basic type check
			if varType != rhsVarType {
				fmt.Fprintf(os.Stderr, "Type mismatch in assignment to %s. Cannot assign value from %s.\n", varName, rhsVarName)
				return nil
			}

			t.addAsm("    // --- Start of assignment to %s from %s ---", varName, rhsVarName)
			switch varType {
			case TypeInt:
				t.addAsm("    LDR X9, =%s", rhsMangledName) // Load address of RHS
				t.addAsm("    LDR W11, [X9]")               // Load value from RHS
				t.addAsm("    LDR X10, =%s", mangledName)    // Load address of LHS
				t.addAsm("    STR W11, [X10]")               // Store value to LHS
			case TypeBool:
				t.addAsm("    LDR X9, =%s", rhsMangledName) // Load address of RHS
				t.addAsm("    LDRB W11, [X9]")              // Load byte from RHS
				t.addAsm("    LDR X10, =%s", mangledName)   // Load address of LHS
				t.addAsm("    STRB W11, [X10]")             // Store byte to LHS
			default:
				fmt.Fprintf(os.Stderr, "Unsupported type for variable-to-variable assignment: %s\n", varType)
				return nil
			}
			t.addAsm("    // --- End of assignment to %s from %s ---", varName, rhsVarName)
			t.addAsm("")

		// TODO: Add cases for other literal types like FloatLiteral, StringLiteral.
		// TODO: Add cases for BinaryExpr (assignment from an arithmetic operation).

		default:
			fmt.Fprintf(os.Stderr, "Unsupported R-value in assignment: %T\n", node.Right)
			if node.Right != nil {
				node.Right.Accept(t)
			}
		}
	case "+=":
		switch varType {
		case TypeInt:
			if rhs, ok := node.Right.(*ast.IntegerLiteral); ok {
				t.addAsm("    // --- Start of compound assignment (+=) to %s ---", varName)
				t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
				t.addAsm("    LDR W11, [X10]")           // Load current value of var
				t.addAsm("    MOV W12, #%s", rhs.Value) // Load immediate integer value from RHS
				t.addAsm("    ADD W11, W11, W12")        // Perform addition
				t.addAsm("    STR W11, [X10]")           // Store result back
				t.addAsm("    // --- End of compound assignment (+=) to %s ---", varName)
				t.addAsm("")
			} else {
				fmt.Fprintf(os.Stderr, "Unsupported R-value in compound assignment for Int: %T\n", node.Right)
			}
		case TypeFloat:
			t.addAsm("    // --- Start of compound assignment (+=) to %s ---", varName)
			t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
			t.addAsm("    LDR D8, [X10]")           // Load current value of var into float register D8
			switch rhs := node.Right.(type) {
			case *ast.IntegerLiteral:
				t.addAsm("    MOV W11, #%s", rhs.Value) // Load immediate integer value
				t.addAsm("    SCVTF D9, W11")           // Convert integer in W11 to float in D9
				t.addAsm("    FADD D8, D8, D9")         // Perform float addition
			case *ast.FloatLiteral:
				floatLabel := t.newLabel("float")
				t.addData(fmt.Sprintf("%s: .double %s", floatLabel, rhs.Value))
				t.addAsm("    LDR X11, =%s", floatLabel) // Load address of float literal
				t.addAsm("    LDR D9, [X11]")            // Load float literal into D9
				t.addAsm("    FADD D8, D8, D9")          // Perform float addition
			default:
				fmt.Fprintf(os.Stderr, "Unsupported R-value in compound assignment for Float: %T\n", node.Right)
				t.addAsm("    // --- Aborted compound assignment due to unsupported RHS ---")
				return nil
			}
			t.addAsm("    STR D8, [X10]") // Store result back
			t.addAsm("    // --- End of compound assignment (+=) to %s ---", varName)
			t.addAsm("")
		default:
			fmt.Fprintf(os.Stderr, "Unsupported type for compound assignment (+=): %s\n", varType)
			return nil
		}
	case "-=":
		switch varType {
		case TypeInt:
			if rhs, ok := node.Right.(*ast.IntegerLiteral); ok {
				t.addAsm("    // --- Start of compound assignment (-=) to %s ---", varName)
				t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
				t.addAsm("    LDR W11, [X10]")           // Load current value of var
				t.addAsm("    MOV W12, #%s", rhs.Value) // Load immediate integer value from RHS
				t.addAsm("    SUB W11, W11, W12")        // Perform subtraction
				t.addAsm("    STR W11, [X10]")           // Store result back
				t.addAsm("    // --- End of compound assignment (-=) to %s ---", varName)
				t.addAsm("")
			} else {
				fmt.Fprintf(os.Stderr, "Unsupported R-value in compound assignment for Int: %T\n", node.Right)
			}
		case TypeFloat:
			t.addAsm("    // --- Start of compound assignment (-=) to %s ---", varName)
			t.addAsm("    LDR X10, =%s", mangledName) // Load address of the variable
			t.addAsm("    LDR D8, [X10]")           // Load current value of var into float register D8
			switch rhs := node.Right.(type) {
			case *ast.IntegerLiteral:
				t.addAsm("    MOV W11, #%s", rhs.Value) // Load immediate integer value
				t.addAsm("    SCVTF D9, W11")           // Convert integer in W11 to float in D9
				t.addAsm("    FSUB D8, D8, D9")         // Perform float subtraction
			case *ast.FloatLiteral:
				floatLabel := t.newLabel("float")
				t.addData(fmt.Sprintf("%s: .double %s", floatLabel, rhs.Value))
				t.addAsm("    LDR X11, =%s", floatLabel) // Load address of float literal
				t.addAsm("    LDR D9, [X11]")            // Load float literal into D9
				t.addAsm("    FSUB D8, D8, D9")          // Perform float subtraction
			default:
				fmt.Fprintf(os.Stderr, "Unsupported R-value in compound assignment for Float: %T\n", node.Right)
				t.addAsm("    // --- Aborted compound assignment due to unsupported RHS ---")
				return nil
			}
			t.addAsm("    STR D8, [X10]") // Store result back
			t.addAsm("    // --- End of compound assignment (-=) to %s ---", varName)
			t.addAsm("")
		default:
			fmt.Fprintf(os.Stderr, "Unsupported type for compound assignment (-=): %s\n", varType)
			return nil
		}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported assignment operator: %s\n", node.Operator)
	}

	return nil
}

func (t *Translator) VisitExpressionStmt(node *ast.ExpressionStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ExpressionStmt")
	}
	// The result of the expression (if any) is usually discarded in an expression statement.
	// For example, a function call made for its side effect.
	node.Expression.Accept(t)
	return nil
}

func (t *Translator) VisitIfStmt(node *ast.IfStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting IfStmt")
	}
	// TODO: Implement IfStmt translation (e.g., conditional jumps, labels)
	node.Condition.Accept(t)
	node.Consequence.Accept(t)
	if node.Alternative != nil { // This handles both "else if" (if Alternative is IfStmt) and "else" (if Alternative is BlockStmt)
		node.Alternative.Accept(t)
	}
	return nil
}

func (t *Translator) VisitSwitchStmt(node *ast.SwitchStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting SwitchStmt")
	}
	// TODO: Implement SwitchStmt translation
	if node.Expression != nil {
		node.Expression.Accept(t)
	}
	for _, c := range node.Cases {
		c.Accept(t)
	}
	if node.Default != nil {
		node.Default.Accept(t)
	}
	return nil
}

func (t *Translator) VisitCaseClause(node *ast.CaseClause) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting CaseClause")
	}
	// TODO: Implement CaseClause translation
	for _, expr := range node.Expressions {
		expr.Accept(t)
	}
	for _, stmt := range node.Body {
		stmt.Accept(t)
	}
	return nil
}

func (t *Translator) VisitDefaultClause(node *ast.DefaultClause) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting DefaultClause")
	}
	// TODO: Implement DefaultClause translation
	for _, stmt := range node.Body {
		stmt.Accept(t)
	}
	return nil
}

func (t *Translator) VisitForStmt(node *ast.ForStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ForStmt")
	}
	// TODO: Implement ForStmt translation (e.g., loop setup, jumps)
	if node.Init != nil {
		node.Init.Accept(t)
	}
	if node.Condition != nil {
		node.Condition.Accept(t)
	}
	if node.Post != nil {
		node.Post.Accept(t)
	}
	node.Body.Accept(t)
	return nil
}

func (t *Translator) VisitReturnStmt(node *ast.ReturnStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ReturnStmt")
	}
	// TODO: Implement ReturnStmt translation
	if node.Value != nil {
		node.Value.Accept(t)
	}
	return nil
}

func (t *Translator) VisitBreakStmt(node *ast.BreakStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting BreakStmt")
	}
	// TODO: Implement BreakStmt translation
	return nil
}

func (t *Translator) VisitContinueStmt(node *ast.ContinueStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ContinueStmt")
	}
	// TODO: Implement ContinueStmt translation
	return nil
}

// Statements (continued)

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
		t.addAsm("    LDR W11, [X10]")           // Load current value of var

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
		t.addAsm("    LDR D8, [X10]")           // Load current value of var

		t.addAsm("    LDR X11, =%s", oneLabel)  // Load address of 1.0
		t.addAsm("    LDR D9, [X11]")           // Load 1.0 into D9

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

// ExpressionResult holds the result of an expression evaluation.
// It contains the register where the result is stored and its type.
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

func (t *Translator) VisitBinaryExpr(node *ast.BinaryExpr) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting BinaryExpr: %s\n", node.Operator)
	}

	leftResult := node.Left.Accept(t).(ExpressionResult)
	rightResult := node.Right.Accept(t).(ExpressionResult)

	// Handle string concatenation
	if leftResult.Type == TypeString && rightResult.Type == TypeString {
		if node.Operator == "+" {
			t.needsStringHelpers = true
			return t.concatenateStrings(leftResult, rightResult)
		}
		panic(fmt.Sprintf("Unsupported operator '%s' for strings", node.Operator))
	}

	// Handle type promotion: Int -> Float
	if leftResult.Type != rightResult.Type {
		if leftResult.Type == TypeInt && rightResult.Type == TypeFloat {
			t.addAsm("    // Promoting left operand from INT to FLOAT")
			promotedFloatReg := t.acquireFloatRegister()
			t.addAsm("    SCVTF D%d, X%d", promotedFloatReg, leftResult.Reg)
			t.releaseIntRegister(leftResult.Reg)
			leftResult = ExpressionResult{Reg: promotedFloatReg, Type: TypeFloat}
		} else if leftResult.Type == TypeFloat && rightResult.Type == TypeInt {
			t.addAsm("    // Promoting right operand from INT to FLOAT")
			promotedFloatReg := t.acquireFloatRegister()
			t.addAsm("    SCVTF D%d, X%d", promotedFloatReg, rightResult.Reg)
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
			t.addAsm("    MUL X%d, X%d, X%d", divResultReg, divResultReg, rightResult.Reg)   // divResultReg = (a / n) * n
			t.addAsm("    SUB X%d, X%d, X%d", resultReg, leftResult.Reg, divResultReg)      // resultReg = a - divResultReg
			t.releaseIntRegister(divResultReg)
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
		return ExpressionResult{Reg: resultReg, Type: TypeInt}
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
		default:
			panic(fmt.Sprintf("Unsupported float operator: %s", node.Operator))
		}
		t.releaseFloatRegister(rightResult.Reg)
		return ExpressionResult{Reg: reg, Type: TypeFloat}
	default:
		panic(fmt.Sprintf("Unsupported type in binary expression: %s", leftResult.Type))
	}
}

func (t *Translator) concatenateStrings(left, right ExpressionResult) ExpressionResult {
	t.addAsm("    // --- Start of string concatenation ---")

	// Save callee-saved registers we will use as temporaries (X19, X20, X21)
	// and the link register X30.
	t.addAsm("    STP X19, X20, [SP, #-16]!")
	t.addAsm("    STP X21, X30, [SP, #-16]!")

	// Move original string pointers into safe callee-saved registers
	t.addAsm("    MOV X19, X%d  // Pointer to left string", left.Reg)
	t.addAsm("    MOV X20, X%d  // Pointer to right string", right.Reg)

	// 1. Get length of left string
	t.addAsm("    MOV X0, X19")
	t.addAsm("    BL strlen")
	t.addAsm("    MOV X21, X0  // Store length of left string")

	// 2. Get length of right string
	t.addAsm("    MOV X0, X20")
	t.addAsm("    BL strlen")

	// 3. Allocate memory for new string (len(left) + len(right) + 1)
	t.addAsm("    ADD X0, X0, X21  // Total length")
	t.addAsm("    ADD X0, X0, #1     // Add 1 for null terminator")
	t.addAsm("    BL malloc")
	// X0 now holds the pointer to the new buffer. Save it in X21.
	t.addAsm("    MOV X21, X0      // X21 now holds the new string pointer")

	// 4. Copy left string into new buffer
	t.addAsm("    MOV X0, X21      // 1st arg for strcpy: destination")
	t.addAsm("    MOV X1, X19      // 2nd arg for strcpy: source (left string)")
	t.addAsm("    BL strcpy")

	// 5. Append right string to new buffer
	t.addAsm("    MOV X0, X21      // 1st arg for strcat: destination")
	t.addAsm("    MOV X1, X20      // 2nd arg for strcat: source (right string)")
	t.addAsm("    BL strcat")

	// The final concatenated string is in X21. Move it to a fresh register from our pool.
	newStringReg := t.acquireIntRegister()
	t.addAsm("    MOV X%d, X21", newStringReg)

	// Restore callee-saved registers and link register
	t.addAsm("    LDP X21, X30, [SP], #16")
	t.addAsm("    LDP X19, X20, [SP], #16")

	// Release the original registers
	t.releaseIntRegister(left.Reg)
	t.releaseIntRegister(right.Reg)

	t.addAsm("    // --- End of string concatenation ---")
	return ExpressionResult{Reg: newStringReg, Type: TypeString}
}

func (t *Translator) VisitUnaryExpr(node *ast.UnaryExpr) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting UnaryExpr: %s\n", node.Operator)
	}

	operandResult := node.Right.Accept(t).(ExpressionResult)

	switch node.Operator {
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
	default:
		t.releaseIntRegister(addrReg) // Release register even on panic
		panic(fmt.Sprintf("Loading for type %s not implemented", varType))
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
	// TODO: Implement CompositeLiteralExpr translation
	return nil
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
		fmt.Println("Translator.Visiting CallExpr")
	}

	if ident, ok := node.Function.(*ast.IdentifierExpr); ok && ident.Name == "println" {
		if t.DebugMode {
			fmt.Printf("Translator.VisitCallExpr: Visiting call to identifier: '%s'\n", ident.Name)
		}
		t.needsPrintf = true
		t.addAsm("    // --- Start of println call ---")
		t.handlePrintln(node.Arguments)
		t.addAsm("    // --- End of println call ---")
		return nil
	}

	if t.DebugMode {
		if ident, ok := node.Function.(*ast.IdentifierExpr); ok {
			fmt.Printf("Translator.VisitCallExpr: Unhandled call to identifier: '%s'\n", ident.Name)
		}
	}
	return nil
}

func (t *Translator) handlePrintln(args []ast.Expression) {
	var formatString strings.Builder
	var evaluatedArgs []ExpressionResult

	// 1. Build format string and evaluate expressions
	for _, arg := range args {
		if strLit, ok := arg.(*ast.StringLiteral); ok {
			// Sanitize the string to escape any '%' characters for printf.
			sanitizedStr := strings.ReplaceAll(strings.Trim(strLit.Value, "\""), "%", "%%")
			formatString.WriteString(sanitizedStr)
		} else {
			result := arg.Accept(t).(ExpressionResult)
			evaluatedArgs = append(evaluatedArgs, result)
			switch result.Type {
			case TypeInt:
				formatString.WriteString("%d")
			case TypeFloat:
				formatString.WriteString("%f")
			case TypeString:
				formatString.WriteString("%s")
			default:
				//panic(fmt.Sprintf("Unsupported type for println: %s", result.Type))
			}
		}
	}
	formatString.WriteString("\n")

	// 2. Add the format string to the data section
	formatLabel := t.addStringData(formatString.String())

	// 3. Load arguments into registers for printf
	// Load format string address into X0
	t.addAsm("    LDR X0, =%s", formatLabel)

	intArgCount := 0
	floatArgCount := 0

	for _, arg := range evaluatedArgs {
		switch arg.Type {
		case TypeInt:
			if intArgCount < 7 { // X1-X7 for integer/pointer arguments
				t.addAsm("    MOV X%d, X%d", intArgCount+1, arg.Reg)
				t.releaseIntRegister(arg.Reg)
				intArgCount++
			}
		case TypeFloat:
			if floatArgCount < 8 { // D0-D7 for float arguments
				t.addAsm("    FMOV D%d, D%d", floatArgCount, arg.Reg) // Note: printf varargs start D0
				t.releaseFloatRegister(arg.Reg)
				floatArgCount++
			}
		case TypeString: // Assuming string address is in an X register
			if intArgCount < 7 {
				t.addAsm("    MOV X%d, X%d", intArgCount+1, arg.Reg)
				t.releaseIntRegister(arg.Reg)
				intArgCount++
			}
		}
	}

	// 4. Call printf
	t.addAsm("    BL printf")
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
		fmt.Println("Translator.Visiting CharLiteral")
	}
	// TODO: Implement CharLiteral translation
	return nil
}

func (t *Translator) VisitBoolLiteral(node *ast.BoolLiteral) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting BoolLiteral")
	}
	// TODO: Implement BoolLiteral translation
	return nil
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
		fmt.Println("Translator.Visiting AnonymousStructTypeNode")
	}
	// TODO: Implement AnonymousStructTypeNode translation
	for _, field := range node.Fields {
		field.Accept(t)
	}
	return nil
}

func (t *Translator) VisitAnonymousInterfaceTypeNode(node *ast.AnonymousInterfaceTypeNode) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting AnonymousInterfaceTypeNode")
	}
	// TODO: Implement AnonymousInterfaceTypeNode translation
	for _, method := range node.Methods {
		method.Accept(t)
	}
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
