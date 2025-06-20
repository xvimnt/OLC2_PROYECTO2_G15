
package ast

// --- Visitor Interface ---

type Visitor interface {
	VisitProgram(node *Program) interface{}

	// Declarations
	VisitStructDecl(node *StructDecl) interface{}
	VisitFieldDecl(node *FieldDecl) interface{}
	VisitFunctionDecl(node *FunctionDecl) interface{}
	VisitReceiverDecl(node *ReceiverDecl) interface{}
	VisitParameterDecl(node *ParameterDecl) interface{}
	VisitVarDecl(node *VarDecl) interface{}

	// Statements
	VisitBlockStmt(node *BlockStmt) interface{}
	VisitAssignStmt(node *AssignStmt) interface{}
	VisitExpressionStmt(node *ExpressionStmt) interface{}
	VisitIfStmt(stmt *IfStmt) interface{}
	VisitSwitchStmt(node *SwitchStmt) interface{}
	VisitCaseClause(node *CaseClause) interface{}
	VisitDefaultClause(*DefaultClause) interface{}
	VisitForStmt(*ForStmt) interface{}
	VisitReturnStmt(*ReturnStmt) interface{}
	VisitBreakStmt(node *BreakStmt) interface{}
	VisitContinueStmt(node *ContinueStmt) interface{}
	VisitIncDecStmt(node *IncDecStmt) interface{}

	// Expressions
	VisitBinaryExpr(node *BinaryExpr) interface{}
	VisitUnaryExpr(node *UnaryExpr) interface{}
	VisitParenExpr(node *ParenExpr) interface{}
	VisitIdentifierExpr(node *IdentifierExpr) interface{}
	VisitTypeConversionExpr(node *TypeConversionExpr) interface{}
	VisitCompositeLiteralExpr(node *CompositeLiteralExpr) interface{}
	VisitCompositeElement(node *CompositeElement) interface{}
	VisitIndexAccessExpr(node *IndexAccessExpr) interface{}
	VisitFieldAccessExpr(node *FieldAccessExpr) interface{}
	VisitCallExpr(node *CallExpr) interface{}
	VisitTypeOfExpr(node *TypeOfExpr) interface{}

	// Literals
	VisitIntegerLiteral(node *IntegerLiteral) interface{}
	VisitFloatLiteral(node *FloatLiteral) interface{}
	VisitStringLiteral(node *StringLiteral) interface{}
	VisitCharLiteral(node *CharLiteral) interface{}
	VisitBoolLiteral(node *BoolLiteral) interface{}
	VisitNilLiteral(node *NilLiteral) interface{}
	VisitArrayLiteral(node *ArrayLiteral) interface{}
	VisitMapLiteral(node *MapLiteral) interface{}
	VisitMapEntry(node *MapEntry) interface{}

	// Types
	VisitTypeName(node *TypeName) interface{}
	VisitArrayTypeNode(node *ArrayTypeNode) interface{}
	VisitSliceTypeNode(node *SliceTypeNode) interface{}
	VisitMapTypeNode(node *MapTypeNode) interface{}
	VisitPointerTypeNode(node *PointerTypeNode) interface{}
	VisitFunctionTypeNode(node *FunctionTypeNode) interface{}
	VisitOptionalTypeNode(node *OptionalTypeNode) interface{}
	VisitAnonymousStructTypeNode(node *AnonymousStructTypeNode) interface{}
	VisitAnonymousInterfaceTypeNode(node *AnonymousInterfaceTypeNode) interface{}
	VisitMethodSignatureNode(node *MethodSignatureNode) interface{}
	VisitPrimitiveTypeNode(node *PrimitiveTypeNode) interface{}
}
