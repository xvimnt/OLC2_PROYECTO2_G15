package translator

import (
	"fmt"

	"github.com/xvimnt/OLC2_PROYECTO1_G15/ast"
)

// Translator translates AST nodes into assembly code.
type Translator struct {
	// TODO: Add necessary fields for translation (e.g., output buffer, symbol table)
	output    []string // For simplicity, stores generated assembly lines
	DebugMode bool
}

// NewTranslator creates a new Translator instance.
func NewTranslator(debugMode bool) *Translator {
	return &Translator{
		output:    make([]string, 0),
		DebugMode: debugMode,
	}
}

// GetAssembly returns the generated assembly code.
func (t *Translator) GetAssembly() []string {
	return t.output
}

func (t *Translator) addAsm(instr string, args ...interface{}) {
	if len(args) > 0 {
		t.output = append(t.output, fmt.Sprintf(instr, args...))
	} else {
		t.output = append(t.output, instr)
	}
}

// --- Visitor Interface Implementations ---

func (t *Translator) VisitProgram(node *ast.Program) interface{} {
	if t.DebugMode {
		fmt.Println("Translator.Visiting Program")
	}
	// TODO: Implement Program translation (e.g., setup, global declarations)
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
		fmt.Println("Translator.Visiting FunctionDecl")
	}
	// TODO: Implement FunctionDecl translation (e.g., function prologue, body, epilogue)
	if node.Body != nil {
		node.Body.Accept(t)
	}
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
	// TODO: Implement VarDecl translation (e.g., memory allocation)
	if node.Initializer != nil {
		node.Initializer.Accept(t)
	}
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
	// TODO: Implement ExpressionStmt translation
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
	// TODO: Implement TypeOfExpr translation
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
		fmt.Println("Translator.Visiting IdentifierExpr")
	}
	// TODO: Implement IdentifierExpr translation (e.g., load variable)
	return nil
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
		fmt.Println("Translator.Visiting CallExpr")
	}
	// TODO: Implement CallExpr translation (e.g., function call setup, arg passing)
	node.Function.Accept(t)
	for _, arg := range node.Arguments {
		arg.Accept(t)
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
		fmt.Println("Translator.Visiting StringLiteral")
	}
	// TODO: Implement StringLiteral translation (e.g., data segment, load address)
	return nil
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
