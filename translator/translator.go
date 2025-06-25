package translator

import (
	"fmt"
	"os"
	"regexp"
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
	needsPrintf        bool              // Tracks if printf is used (for .extern printf)
	hasIntFormatStr    bool              // Tracks if the integer format string has been added
	hasFloatFormatStr  bool              // Tracks if the float format string has been added
	hasStringFormatStr bool              // Tracks if the string format string has been added
	currentFuncDef     *ast.FunctionDecl // Keep track of the current function being defined
	DebugMode          bool
}

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
		needsPrintf:        false,
		hasIntFormatStr:    false,
		hasFloatFormatStr:  false,
		hasStringFormatStr: false,
		currentFuncDef:     nil,
		DebugMode:          debugMode,
	}
	return t
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

	// We don't visit node.Left because we've already processed it to get the varName.
	// Visiting it would be redundant or incorrect if it's not designed to be visited in this context.

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
	// TODO: Implement IncDecStmt translation
	// Example: node.LValue.Accept(t) // if LValue needs translation or analysis
	return nil
}

// Expressions

func (t *Translator) VisitTypeOfExpr(node *ast.TypeOfExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting TypeOfExpr")
	}
	// TODO: Implement TypeOfExpr translation (usually for type checking or metadata)
	// Example: node.Expression.Accept(t)
	return nil
}

func (t *Translator) VisitBinaryExpr(node *ast.BinaryExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting BinaryExpr")
	}
	// TODO: Implement BinaryExpr translation (e.g., arithmetic, logical ops)
	node.Left.Accept(t)
	node.Right.Accept(t)
	return nil
}

func (t *Translator) VisitUnaryExpr(node *ast.UnaryExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting UnaryExpr")
	}
	// TODO: Implement UnaryExpr translation
	if node.Right != nil {
		node.Right.Accept(t)
	}
	return nil
}

func (t *Translator) VisitParenExpr(node *ast.ParenExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting ParenExpr")
	}
	// TODO: Implement ParenExpr translation
	node.Expression.Accept(t)
	return nil
}

func (t *Translator) VisitIdentifierExpr(node *ast.IdentifierExpr) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting IdentifierExpr: %s\n", node.Name)
	}
	// For an identifier, we might need to load its value from memory or a register.
	// If it's a function name (like in CallExpr), the CallExpr handler uses the name.
	// If it's a variable, we'd generate code to load it.
	// For now, just return its name. This might be used by parent nodes.
	// TODO: Implement variable loading, etc.
	return node.Name
}

func (t *Translator) VisitTypeConversionExpr(node *ast.TypeConversionExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting TypeConversionExpr")
	}
	// TODO: Implement TypeConversionExpr translation
	if node.Expression != nil {
		node.Expression.Accept(t)
	}
	if node.TargetType != nil {
		node.TargetType.Accept(t)
	}
	return nil
}

func (t *Translator) VisitCompositeLiteralExpr(node *ast.CompositeLiteralExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting CompositeLiteralExpr")
	}
	// TODO: Implement CompositeLiteralExpr translation
	if node.Type != nil {
		node.Type.Accept(t)
	}
	for _, el := range node.Elements {
		el.Accept(t)
	}
	return nil
}

func (t *Translator) VisitCompositeElement(node *ast.CompositeElement) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting CompositeElement")
	}
	// TODO: Implement CompositeElement translation
	if node.Key != nil {
		node.Key.Accept(t)
	}
	node.Value.Accept(t)
	return nil
}

func (t *Translator) VisitIndexAccessExpr(node *ast.IndexAccessExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting IndexAccessExpr")
	}
	// TODO: Implement IndexAccessExpr translation
	node.Receiver.Accept(t)
	node.Index.Accept(t)
	return nil
}

func (t *Translator) VisitFieldAccessExpr(node *ast.FieldAccessExpr) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting FieldAccessExpr")
	}
	// TODO: Implement FieldAccessExpr translation
	node.Receiver.Accept(t)
	// node.FieldName is an IdentifierExpr, usually handled by its name string
	return nil
}

func (t *Translator) VisitCallExpr(node *ast.CallExpr) interface{} {
	if t.DebugMode {
		if ident, ok := node.Function.(*ast.IdentifierExpr); ok {
			fmt.Printf("Translator.VisitCallExpr: Visiting call to identifier: '%s'\n", ident.Name)
		} else {
			fmt.Printf("Translator.VisitCallExpr: Visiting call to expression of type %T\n", node.Function)
		}
	}

	if ident, ok := node.Function.(*ast.IdentifierExpr); ok {
		// Special handling for println
		if ident.Name == "println" {
			t.handlePrintln(node.Arguments)
		} else {
			// Generic function call handling
			t.addAsm("    BL %s", ident.Name)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Translator Error: Non-identifier function calls not yet supported\n")
	}

	return nil
}

// handlePrintlnStringLiteral processes a string literal within a 'println' call,
// handling string interpolation for variables.
func (t *Translator) handlePrintln(args []ast.Expression) {
	t.needsPrintf = true
	var formatParts []string

	type printArg struct {
		mangledName string
		varType     VarType
	}
	var printArgs []printArg

	for i, arg := range args {
		if i > 0 {
			formatParts = append(formatParts, " ")
		}

		switch v := arg.(type) {
		case *ast.StringLiteral:
			rawVal := v.Value
			if len(rawVal) >= 2 && rawVal[0] == '"' && rawVal[len(rawVal)-1] == '"' {
				rawVal = rawVal[1 : len(rawVal)-1]
			}

			re := regexp.MustCompile(`\$([a-zA-Z_]\w*)`)
			matches := re.FindAllStringSubmatchIndex(rawVal, -1)
			lastIndex := 0
			for _, match := range matches {
				formatParts = append(formatParts, rawVal[lastIndex:match[0]])
				varName := rawVal[match[2]:match[3]]
				mangledName, varType, exists := t.lookupSymbol(varName)
				if !exists {
					fmt.Fprintf(os.Stderr, "Error: Use of undeclared variable '%s' in println.\n", varName)
					formatParts = append(formatParts, "[UNDECLARED]")
					continue
				}

				switch varType {
				case TypeFloat:
					formatParts = append(formatParts, "%f")
				case TypeString, TypeBool:
					formatParts = append(formatParts, "%s")
				default:
					formatParts = append(formatParts, "%d")
				}
				printArgs = append(printArgs, printArg{mangledName: mangledName, varType: varType})
				lastIndex = match[1]
			}
			formatParts = append(formatParts, rawVal[lastIndex:])

		case *ast.IdentifierExpr:
			varName := v.Name
			mangledName, varType, exists := t.lookupSymbol(varName)
			if !exists {
				fmt.Fprintf(os.Stderr, "Error: Use of undeclared variable '%s' in println.\n", varName)
				formatParts = append(formatParts, "[UNDECLARED]")
				continue
			}

			switch varType {
			case TypeFloat:
				formatParts = append(formatParts, "%f")
			case TypeString, TypeBool:
				formatParts = append(formatParts, "%s")
			default:
				formatParts = append(formatParts, "%d")
			}
			printArgs = append(printArgs, printArg{mangledName: mangledName, varType: varType})

		default:
			formatParts = append(formatParts, "[?]")
		}
	}

	formatParts = append(formatParts, "\n")
	formatString := strings.Join(formatParts, "")
	formatLabel := t.addStringData(formatString)

	t.addAsm("    // --- Start of println call ---")
	t.addAsm("    LDR X0, =%s", formatLabel)

	intArgCount := 1
	floatArgCount := 0

	for _, arg := range printArgs {
		switch arg.varType {
		case TypeFloat:
			if floatArgCount < 8 {
				t.addAsm("    LDR X9, =%s", arg.mangledName)
				t.addAsm("    LDR D%d, [X9]", floatArgCount)
				floatArgCount++
			}
		case TypeString:
			if intArgCount < 8 {
				t.addAsm("    LDR X%d, =%s", intArgCount, arg.mangledName)
				t.addAsm("    LDR X%d, [X%d]", intArgCount, intArgCount)
				intArgCount++
			}
		case TypeBool:
			if intArgCount < 8 {
				trueLabel := t.addStringData("true")
				falseLabel := t.addStringData("false")
				argReg := fmt.Sprintf("X%d", intArgCount)
				valReg := "W9"
				trueAddrReg := "X10"
				falseAddrReg := "X11"

				t.addAsm("    LDR %s, =%s", trueAddrReg, trueLabel)
				t.addAsm("    LDR %s, =%s", falseAddrReg, falseLabel)
				t.addAsm("    LDR X12, =%s", arg.mangledName)
				t.addAsm("    LDRB %s, [X12]", valReg)
				t.addAsm("    CMP %s, #0", valReg)
				t.addAsm("    CSEL %s, %s, %s, EQ", argReg, falseAddrReg, trueAddrReg)
				intArgCount++
			}
		default: // TypeInt, TypeUnknown
			if intArgCount < 8 {
				t.addAsm("    LDR X9, =%s", arg.mangledName)
				t.addAsm("    LDR W%d, [X9]", intArgCount)
				intArgCount++
			}
		}
	}

	t.addAsm("    BL printf")
	t.addAsm("    // --- End of println call ---")
	t.addAsm("")
}

func (t *Translator) VisitIntegerLiteral(node *ast.IntegerLiteral) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting IntegerLiteral")
	}
	// TODO: Implement IntegerLiteral translation (e.g., load immediate value)
	return nil
}

func (t *Translator) VisitFloatLiteral(node *ast.FloatLiteral) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting FloatLiteral")
	}
	// TODO: Implement FloatLiteral translation
	return nil
}

func (t *Translator) VisitStringLiteral(node *ast.StringLiteral) interface{} {
	if t.DebugMode {
		fmt.Printf("Translator.Visiting StringLiteral: %s\n", node.Value)
	}
	// node.Value is like "Hello, world!" (includes the quotes).
	// .asciz directive in GAS expects the string content, typically also quoted in the assembly source.
	// However, our addStringData function expects the raw content to put inside "".
	// So, we need to strip the outer quotes from node.Value before passing.
	rawValue := node.Value
	if len(rawValue) >= 2 && rawValue[0] == '"' && rawValue[len(rawValue)-1] == '"' {
		rawValue = rawValue[1 : len(rawValue)-1]
	}
	// TODO: Handle escape sequences within the string if V lang supports them (e.g., \n, \t)
	// For now, assuming rawValue is what we want between the .asciz "".
	label := t.addStringData(rawValue) // addStringData now handles quoting for .asciz
	return label                       // Return the label for this string in the .data section
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
