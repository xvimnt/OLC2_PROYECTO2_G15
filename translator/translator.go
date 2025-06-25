package translator

import (
	"fmt"
	"os" // Added for Fprintf to os.Stderr for errors

	"github.com/xvimnt/OLC2_PROYECTO2_G15/ast"
)

// Translator translates AST nodes into assembly code.
type Translator struct {
	asm            []string          // Stores generated .text section assembly lines
	dataSection    []string          // Stores generated .data section assembly lines
	stringCounter  int               // For generating unique string labels
	needsPrintf    bool              // Tracks if printf is used (for .extern printf)
	currentFuncDef *ast.FunctionDecl // Keep track of the current function being defined
	DebugMode      bool
}

// NewTranslator creates a new Translator instance.
func NewTranslator(debugMode bool) *Translator {
	return &Translator{
		asm:            make([]string, 0),
		dataSection:    make([]string, 0),
		stringCounter:  0,
		needsPrintf:    false,
		currentFuncDef: nil,
		DebugMode:      debugMode,
	}
}

// addStringData adds a string to the .data section and returns its label.
// It uses %q to handle proper quoting and escaping for the assembler.
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
		t.addAsm(".global _start")
		t.addAsm("_start:")
	} else if node.Name != nil {
		t.addAsm(".global %s", node.Name.Name)
		t.addAsm("%s:", node.Name.Name)
	} else {
		// Handle anonymous functions or error, though V doesn't have them at top level like this
		t.addAsm("anonymous_func_%d:", t.stringCounter) // Placeholder for unnamed funcs
		t.stringCounter++
	}

	// Prologue
	t.addAsm("    STP X29, X30, [SP, #-16]") // Save Frame Pointer (x29) and Link Register (x30) to stack, pre-decrement SP by 16
	t.addAsm("    MOV X29, SP")              // Set current stack pointer as the new Frame Pointer

	// TODO: Allocate space for local variables based on function needs

	if node.Body != nil {
		node.Body.Accept(t)
	}

	// Epilogue
	// Ensure a return path even if no explicit return statement for void functions (like main often is implicitly)
	// For non-void functions, an explicit return statement should handle loading the return value.
	// For main, or functions ending without explicit return, this provides a standard exit.
	// Epilogue for main/_start should handle process exit.
	// For other functions, it's a standard return.
	if node.Name != nil && node.Name.Name == "main" {
		// The body of main will now contain the syscalls for println.
		// We still need to ensure the process exits correctly.
		t.addAsm("    // Syscall: exit(code=0)")
		t.addAsm("    MOV X8, #93") // exit syscall number
		t.addAsm("    MOV X0, #0")  // exit code 0
		t.addAsm("    SVC #0")      // trigger syscall
	} else {
		if node.Name != nil {
			t.addAsm(".L%s_epilogue:", node.Name.Name) // Label for potential jumps to epilogue
		} else {
			t.addAsm(".L_anonymous_func_%d_epilogue:", t.stringCounter-1) // Match potential anonymous label
		}
		t.addAsm("    MOV W0, #0")              // Default return code 0 for other functions
		t.addAsm("    LDP X29, X30, [SP], #16") // Restore FP, LR from stack, post-increment SP by 16
		t.addAsm("    RET")
	}
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
	initialValue := "0" // Default to 0 if no initializer

	if node.Initializer != nil {
		// For now, only handle integer literal initializers
		if intLit, ok := node.Initializer.(*ast.IntegerLiteral); ok {
			initialValue = intLit.Value
		}
		// TODO: Handle other initializer types
	}
	// Add to .data section for global/static variables.
	// Note: This assumes all VarDecls are global. Local variables would need stack allocation.
	t.dataSection = append(t.dataSection, fmt.Sprintf("%s: .word %s", varName, initialValue))

	return nil
}

// Statements
func (t *Translator) VisitBlockStmt(node *ast.BlockStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting BlockStmt")
	}
	// TODO: Implement BlockStmt translation
	for _, stmt := range node.Statements {
		stmt.Accept(t)
	}
	return nil
}

func (t *Translator) VisitAssignStmt(node *ast.AssignStmt) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting AssignStmt")
	}
	// TODO: Implement AssignStmt translation
	// Based on ast.AssignStmt, Left and Right are single ast.Expression nodes.
	if node.Left != nil {
		node.Left.Accept(t)
	}
	if node.Right != nil {
		node.Right.Accept(t)
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
		// Enhanced debugging
		if ident, ok := node.Function.(*ast.IdentifierExpr); ok {
			fmt.Printf("Translator.VisitCallExpr: Visiting call to identifier: '%s'\n", ident.Name)
		} else {
			fmt.Printf("Translator.VisitCallExpr: Visiting call to expression of type %T\n", node.Function)
		}
	}

	if ident, ok := node.Function.(*ast.IdentifierExpr); ok && ident.Name == "println!" {
		for _, arg := range node.Arguments {
			if strLit, ok := arg.(*ast.StringLiteral); ok {
				// Add the string to the .data section
				strLabel := t.addStringData(strLit.Value + "\n") // Add newline for println
				strLen := len(strLit.Value) + 1

				// Generate syscall for write
				t.addAsm("    // Syscall: write(fd=1, buf=%s, count=%d)", strLabel, strLen)
				t.addAsm("    MOV X8, #64")      // syscall number for write
				t.addAsm("    MOV X0, #1")       // file descriptor 1 (stdout)
				t.addAsm("    LDR X1, =%s", strLabel) // address of the string
				t.addAsm("    MOV X2, #%d", strLen) // length of the string
				t.addAsm("    SVC #0")           // make the syscall
			} else {
				// Fallback for non-string arguments
				fmt.Fprintf(os.Stderr, "Warning: println argument is not a string literal, skipping.\n")
			}
		}
	} else {
		// Handle other function calls (including user-defined)
		// This is a simplified placeholder for calling other functions
		if node.Function != nil {
			if name, ok := node.Function.Accept(t).(string); ok {
				t.addAsm("    BL %s", name)
			}
		}
	}
	return nil
}

// Literals
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
