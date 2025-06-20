package ast

import "fmt"

// Node is the interface that all AST nodes must implement.
type Node interface {
	Accept(visitor Visitor) interface{}
	String() string // For debugging and testing
}

// Expression is the interface for all expression nodes.
type Expression interface {
	Node
	expressionNode()
}

// Statement is the interface for all statement nodes.
type Statement interface {
	Node
	statementNode()
}

// Declaration is the interface for all top-level declarations.
type Declaration interface {
	Node
	declarationNode()
}

// TypeNode represents a type in the language.
type TypeNode interface {
	Node
	typeNode()
}

// --- Program Structure ---

// Program represents the root of the AST, a sequence of top-level declarations.
type Program struct {
	Declarations []Declaration
}

func (p *Program) Accept(visitor Visitor) interface{} { return visitor.VisitProgram(p) }
func (p *Program) String() string                     { return "Program" } // Placeholder
func (p *Program) declarationNode()                   {}                   // Program can be considered a top-level container

// --- Declarations ---

// StructDecl represents a struct declaration.
type StructDecl struct {
	Name   *IdentifierExpr
	Fields []*FieldDecl
}

func (sd *StructDecl) Accept(visitor Visitor) interface{} { return visitor.VisitStructDecl(sd) }
func (sd *StructDecl) declarationNode()                   {}
func (sd *StructDecl) String() string                     { return "StructDecl: " + sd.Name.Name }

// FieldDecl represents a field in a struct declaration.
type FieldDecl struct {
	Name *IdentifierExpr
	Type TypeNode
}

func (fd *FieldDecl) Accept(visitor Visitor) interface{} { return visitor.VisitFieldDecl(fd) }
func (fd *FieldDecl) String() string                     { return "FieldDecl: " + fd.Name.Name }

// FunctionDecl represents a function declaration.
type FunctionDecl struct {
	Receiver   *ReceiverDecl // Optional
	Name       *IdentifierExpr
	Parameters []*ParameterDecl
	ReturnType TypeNode // Optional
	Body       *BlockStmt
}

func (fd *FunctionDecl) Accept(visitor Visitor) interface{} { return visitor.VisitFunctionDecl(fd) }
func (fd *FunctionDecl) declarationNode()                   {}
func (fd *FunctionDecl) String() string                     { return "FunctionDecl: " + fd.Name.Name }

// ReceiverDecl represents a method receiver.
type ReceiverDecl struct {
	Name *IdentifierExpr
	Type TypeNode // Type can be PointerTypeNode if '*' is used
}

func (rd *ReceiverDecl) Accept(visitor Visitor) interface{} { return visitor.VisitReceiverDecl(rd) }
func (rd *ReceiverDecl) String() string {
	return "ReceiverDecl: (" + rd.Name.Name + " " + rd.Type.String() + ")"
}

// ParameterDecl represents a function parameter.
type ParameterDecl struct {
	Name   *IdentifierExpr
	Type   TypeNode
	Line   int // <-- AGREGADO
	Column int // <-- AGREGADO
}

func (pd *ParameterDecl) Accept(visitor Visitor) interface{} { return visitor.VisitParameterDecl(pd) }
func (pd *ParameterDecl) String() string {
	return "ParameterDecl: " + pd.Name.Name + " " + pd.Type.String()
}

// VarDecl represents a variable declaration.
type VarDecl struct {
	Name         *IdentifierExpr
	IsMutable    bool
	ExplicitType TypeNode   // Optional, nil if not provided (e.g., in := or inferred)
	Initializer  Expression // Optional for 'mut x int', required for ':='
	IsShortHand  bool       // True if ':=' was used
	Line         int        // <-- AGREGADO
	Column       int        // <-- AGREGADO
}

func (vd *VarDecl) Accept(visitor Visitor) interface{} { return visitor.VisitVarDecl(vd) }
func (vd *VarDecl) declarationNode()                   {}
func (vd *VarDecl) statementNode()                     {} // VarDecl can also be a statement
func (vd *VarDecl) String() string                     { return "VarDecl: " + vd.Name.Name }

// --- Statements ---

// BlockStmt represents a block of statements.
type BlockStmt struct {
	Statements []Statement
	Line       int // <-- AGREGADO
	Column     int // <-- AGREGADO
}

func (bs *BlockStmt) Accept(visitor Visitor) interface{} { return visitor.VisitBlockStmt(bs) }
func (bs *BlockStmt) statementNode()                     {}
func (bs *BlockStmt) String() string                     { return "BlockStmt" }

// AssignStmt represents an assignment statement.
type AssignStmt struct {
	Left     Expression // LValue (IdentifierExpr, FieldAccessExpr, IndexAccessExpr)
	Operator string     // EQ, ADD_EQ, SUB_EQ, etc.
	Right    Expression
	Line     int // <-- AGREGAR AQUÍ
	Column   int // <-- AGREGAR AQUÍ

}

func (as *AssignStmt) Accept(visitor Visitor) interface{} { return visitor.VisitAssignStmt(as) }
func (as *AssignStmt) statementNode()                     {}
func (as *AssignStmt) String() string                     { return "AssignStmt" }

// ExpressionStmt represents an expression used as a statement.
type ExpressionStmt struct {
	Expression Expression
}

func (es *ExpressionStmt) Accept(visitor Visitor) interface{} { return visitor.VisitExpressionStmt(es) }
func (es *ExpressionStmt) statementNode()                     {}
func (es *ExpressionStmt) String() string                     { return "ExpressionStmt" }

// IfStmt represents an if-else statement. "else if" is modeled as
// an IfStmt in the Alternative field.
type IfStmt struct {
	Condition   Expression
	Consequence *BlockStmt
	Alternative Statement // Can be another IfStmt, a BlockStmt, or nil
}

func (is *IfStmt) Accept(visitor Visitor) interface{} { return visitor.VisitIfStmt(is) }
func (is *IfStmt) statementNode()                     {}
func (is *IfStmt) String() string                     { return "IfStmt" }

// SwitchStmt represents a switch statement.
type SwitchStmt struct {
	Expression Expression // Optional
	Cases      []*CaseClause
	Default    *DefaultClause // Optional
}

func (ss *SwitchStmt) Accept(visitor Visitor) interface{} { return visitor.VisitSwitchStmt(ss) }
func (ss *SwitchStmt) statementNode()                     {}
func (ss *SwitchStmt) String() string                     { return "SwitchStmt" }

// CaseClause represents a case in a switch statement.
type CaseClause struct {
	Expressions []Expression // List of expressions for this case
	Body        []Statement  // Statements to execute
}

func (cc *CaseClause) Accept(visitor Visitor) interface{} { return visitor.VisitCaseClause(cc) }
func (cc *CaseClause) String() string                     { return "CaseClause" }

// DefaultClause represents the default case in a switch statement.
type DefaultClause struct {
	Body []Statement // Statements to execute
}

func (dc *DefaultClause) Accept(visitor Visitor) interface{} { return visitor.VisitDefaultClause(dc) }
func (dc *DefaultClause) String() string                     { return "DefaultClause" }

// ForStmt represents a for loop.
type ForStmt struct {
	Init        Statement       // Optional, for C-style for or var decl in range
	Condition   Expression      // Optional, for C-style for or while-style for
	Post        Statement       // Optional, for C-style for
	RangeKey    *IdentifierExpr // Optional, for range-based for
	RangeValue  *IdentifierExpr // Optional, for range-based for (can be the only var)
	RangeSource Expression      // Optional, for range-based for
	Body        *BlockStmt
	IsRangeLoop bool // True if 'in' keyword was used
	IsWhileLoop bool // True if only condition expression is present
}

func (fs *ForStmt) Accept(visitor Visitor) interface{} { return visitor.VisitForStmt(fs) }
func (fs *ForStmt) statementNode()                     {}
func (fs *ForStmt) String() string                     { return "ForStmt" }

// ReturnStmt represents a return statement.
type ReturnStmt struct {
	Value Expression // Optional
}

func (rs *ReturnStmt) Accept(visitor Visitor) interface{} { return visitor.VisitReturnStmt(rs) }
func (rs *ReturnStmt) statementNode()                     {}
func (rs *ReturnStmt) String() string                     { return "ReturnStmt" }

// BreakStmt represents a break statement.
type BreakStmt struct{}

func (bs *BreakStmt) Accept(visitor Visitor) interface{} { return visitor.VisitBreakStmt(bs) }
func (bs *BreakStmt) statementNode()                     {}
func (bs *BreakStmt) String() string                     { return "BreakStmt" }

// ContinueStmt represents a continue statement.
type ContinueStmt struct{}

func (cs *ContinueStmt) Accept(visitor Visitor) interface{} { return visitor.VisitContinueStmt(cs) }
func (cs *ContinueStmt) statementNode()                     {}
func (cs *ContinueStmt) String() string                     { return "ContinueStmt" }

// IncDecStmt represents an increment or decrement statement (e.g., i++, i--).
type IncDecStmt struct {
	LValue   Expression // The lvalue being incremented or decremented (IdentifierExpr, FieldAccessExpr, IndexAccessExpr)
	Operator string     // "++" or "--"
}

func (ids *IncDecStmt) Accept(visitor Visitor) interface{} { return visitor.VisitIncDecStmt(ids) }
func (ids *IncDecStmt) statementNode()                     {}
func (ids *IncDecStmt) String() string {
	return fmt.Sprintf("IncDecStmt: %s %s", ids.LValue.String(), ids.Operator)
}

// --- Expressions ---

// BinaryExpr represents a binary operation.
type BinaryExpr struct {
	Left     Expression
	Operator string // e.g., +, -, *, /, ==, !=, <, <=, >, >=, AND, OR
	Right    Expression
	Line     int // <-- AGREGAR AQUÍ
	Column   int // <-- AGREGAR AQUÍ
}

func (be *BinaryExpr) Accept(visitor Visitor) interface{} { return visitor.VisitBinaryExpr(be) }
func (be *BinaryExpr) expressionNode()                    {}
func (be *BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", be.Left.String(), be.Operator, be.Right.String())
}

// ... (rest of the code remains the same)
// UnaryExpr represents a unary operation (e.g., -x, !ok).
type UnaryExpr struct {
	Operator string     // e.g., "-", "!", "&", "*"
	Right    Expression // The operand
}

func (ue *UnaryExpr) Accept(visitor Visitor) interface{} { return visitor.VisitUnaryExpr(ue) }
func (ue *UnaryExpr) expressionNode()                    {}
func (ue *UnaryExpr) String() string                     { return "UnaryExpr" }

// ParenExpr represents a parenthesized expression.
type ParenExpr struct {
	Expression Expression
}

func (pe *ParenExpr) Accept(visitor Visitor) interface{} { return visitor.VisitParenExpr(pe) }
func (pe *ParenExpr) expressionNode()                    {}
func (pe *ParenExpr) String() string                     { return "ParenExpr" }

// IdentifierExpr represents an identifier.
type IdentifierExpr struct {
	Name   string
	Line   int // <-- AGREGAR AQUÍ
	Column int // <-- AGREGAR AQUÍ
}

func (ie *IdentifierExpr) Accept(visitor Visitor) interface{} { return visitor.VisitIdentifierExpr(ie) }
func (ie *IdentifierExpr) expressionNode()                    {}
func (ie *IdentifierExpr) String() string                     { return "IdentifierExpr: " + ie.Name }

// TypeConversionExpr represents a type conversion.
type TypeConversionExpr struct {
	TargetType TypeNode
	Expression Expression
}

func (tce *TypeConversionExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitTypeConversionExpr(tce)
}
func (tce *TypeConversionExpr) expressionNode() {}
func (tce *TypeConversionExpr) String() string  { return "TypeConversionExpr" }

// CompositeLiteralExpr represents a composite literal for structs, arrays, slices, or maps with explicit types.
type CompositeLiteralExpr struct {
	Type     TypeNode // TypeName, ArrayTypeNode, SliceTypeNode, or MapTypeNode
	Elements []*CompositeElement
}

func (cle *CompositeLiteralExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitCompositeLiteralExpr(cle)
}
func (cle *CompositeLiteralExpr) expressionNode() {}
func (cle *CompositeLiteralExpr) String() string  { return "CompositeLiteralExpr" }

// CompositeElement represents an element in a composite literal.
type CompositeElement struct {
	Key   Expression // Optional, used for struct fields or map keys
	Value Expression
}

func (ce *CompositeElement) Accept(visitor Visitor) interface{} {
	return visitor.VisitCompositeElement(ce)
}
func (ce *CompositeElement) String() string { return "CompositeElement" }

// IndexAccessExpr represents an index access (e.g., arr[i]).
type IndexAccessExpr struct {
	Receiver Expression
	Index    Expression
	Line     int // <-- AGREGAR AQUÍ
	Column   int // <-- AGREGAR AQUÍ
}

func (iae *IndexAccessExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitIndexAccessExpr(iae)
}
func (iae *IndexAccessExpr) expressionNode() {}
func (iae *IndexAccessExpr) String() string  { return "IndexAccessExpr" }

// FieldAccessExpr represents a field access (e.g., obj.field).
type FieldAccessExpr struct {
	Receiver Expression
	Field    *IdentifierExpr
	Line     int // <-- AGREGAR AQUÍ
	Column   int // <-- AGREGAR AQUÍ
}

func (fae *FieldAccessExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitFieldAccessExpr(fae)
}
func (fae *FieldAccessExpr) expressionNode() {}
func (fae *FieldAccessExpr) String() string  { return "FieldAccessExpr" }

// CallExpr represents a function or method call.
type CallExpr struct {
	Function  Expression // IdentifierExpr, FieldAccessExpr, etc.
	Arguments []Expression
	Line      int // <-- AGREGAR AQUÍ
	Column    int // <-- AGREGAR AQUÍ
}

func (ce *CallExpr) Accept(visitor Visitor) interface{} { return visitor.VisitCallExpr(ce) }
func (ce *CallExpr) expressionNode()                    {}
func (ce *CallExpr) String() string                     { return "CallExpr" }

// TypeOfExpr represents a TypeOf(expression) call.
type TypeOfExpr struct {
	Expression Expression
}

func (toe *TypeOfExpr) Accept(visitor Visitor) interface{} { return visitor.VisitTypeOfExpr(toe) }
func (toe *TypeOfExpr) expressionNode()                    {}
func (toe *TypeOfExpr) String() string                     { return fmt.Sprintf("TypeOf(%s)", toe.Expression.String()) }

// --- Literals (all are Expressions) ---

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Value  string // Keep as string to handle large numbers or different bases if needed later
	Line   int    // <-- AGREGAR AQUÍ
	Column int    // <-- AGREGAR AQUÍ
}

func (il *IntegerLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitIntegerLiteral(il) }
func (il *IntegerLiteral) expressionNode()                    {}
func (il *IntegerLiteral) String() string                     { return "IntegerLiteral: " + il.Value }

// FloatLiteral represents a float literal.
type FloatLiteral struct {
	Value  string
	Line   int // <-- AGREGAR AQUÍ
	Column int // <-- AGREGAR AQUÍ
}

func (fl *FloatLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitFloatLiteral(fl) }
func (fl *FloatLiteral) expressionNode()                    {}
func (fl *FloatLiteral) String() string                     { return "FloatLiteral: " + fl.Value }

// StringLiteral represents a string literal.
type StringLiteral struct {
	Value  string
	Line   int // <-- AGREGAR AQUÍ
	Column int // <-- AGREGAR AQUÍ
}

func (sl *StringLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitStringLiteral(sl) }
func (sl *StringLiteral) expressionNode()                    {}
func (sl *StringLiteral) String() string                     { return "StringLiteral: \"" + sl.Value + "\"" }

// CharLiteral represents a character literal.
type CharLiteral struct {
	Value  string // e.g., 'a'
	Line   int    // <-- AGREGAR AQUÍ
	Column int    // <-- AGREGAR AQUÍ
}

func (cl *CharLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitCharLiteral(cl) }
func (cl *CharLiteral) expressionNode()                    {}
func (cl *CharLiteral) String() string                     { return "CharLiteral: '" + cl.Value + "'" }

// BoolLiteral represents a boolean literal (true or false).
type BoolLiteral struct {
	Value  bool
	Line   int // <-- AGREGAR AQUÍ
	Column int // <-- AGREGAR AQUÍ
}

func (bl *BoolLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitBoolLiteral(bl) }
func (bl *BoolLiteral) expressionNode()                    {}
func (bl *BoolLiteral) String() string {
	if bl.Value {
		return "BoolLiteral: true"
	} else {
		return "BoolLiteral: false"
	}
}

// NilLiteral represents a nil literal.
type NilLiteral struct{}

func (nl *NilLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitNilLiteral(nl) }
func (nl *NilLiteral) expressionNode()                    {}
func (nl *NilLiteral) String() string                     { return "NilLiteral" }

// ArrayLiteral represents an array literal (e.g., [1, 2, 3]). This is for when no explicit type is given.
type ArrayLiteral struct {
	Elements []Expression
	Line     int
	Column   int
}

func (al *ArrayLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitArrayLiteral(al) }
func (al *ArrayLiteral) expressionNode()                    {}
func (al *ArrayLiteral) String() string                     { return "ArrayLiteral" }

// MapLiteral represents a map literal (e.g., {key1: val1, key2: val2}). This is for when no explicit type is given.
type MapLiteral struct {
	Entries []*MapEntry
	Line    int
	Column  int
}

func (ml *MapLiteral) Accept(visitor Visitor) interface{} { return visitor.VisitMapLiteral(ml) }
func (ml *MapLiteral) expressionNode()                    {}
func (ml *MapLiteral) String() string                     { return "MapLiteral" }

// MapEntry represents a key-value pair in a map literal.
type MapEntry struct {
	Key   Expression
	Value Expression
}

func (me *MapEntry) Accept(visitor Visitor) interface{} { return visitor.VisitMapEntry(me) }
func (me *MapEntry) String() string                     { return "MapEntry" }

// --- Types ---

// TypeName represents a named type (e.g., int, string, MyStruct).
type TypeName struct {
	Name string // Can be qualified (e.g., package.Type)
}

func (tn *TypeName) Accept(visitor Visitor) interface{} { return visitor.VisitTypeName(tn) }
func (tn *TypeName) typeNode()                          {}
func (tn *TypeName) String() string                     { return tn.Name }

// ArrayTypeNode represents an array type (e.g., [5]int).
type ArrayTypeNode struct {
	Size        Expression // Optional, from grammar `expression?`
	ElementType TypeNode
}

func (atn *ArrayTypeNode) Accept(visitor Visitor) interface{} { return visitor.VisitArrayTypeNode(atn) }
func (atn *ArrayTypeNode) typeNode()                          {}
func (atn *ArrayTypeNode) String() string {
	s := "[]"
	if atn.Size != nil {
		s = "[" + atn.Size.String() + "]"
	}
	return s + atn.ElementType.String()
}

// SliceTypeNode represents a slice type (e.g., []int).
type SliceTypeNode struct {
	ElementType TypeNode
}

func (stn *SliceTypeNode) Accept(visitor Visitor) interface{} { return visitor.VisitSliceTypeNode(stn) }
func (stn *SliceTypeNode) typeNode()                          {}
func (stn *SliceTypeNode) String() string                     { return "[]" + stn.ElementType.String() }

// MapTypeNode represents a map type (e.g., map[string]int).
type MapTypeNode struct {
	KeyType   TypeNode
	ValueType TypeNode
}

func (mtn *MapTypeNode) Accept(visitor Visitor) interface{} { return visitor.VisitMapTypeNode(mtn) }
func (mtn *MapTypeNode) typeNode()                          {}
func (mtn *MapTypeNode) String() string {
	return "map[" + mtn.KeyType.String() + "]" + mtn.ValueType.String()
}

// PointerTypeNode represents a pointer type (e.g., *int).
type PointerTypeNode struct {
	ElementType TypeNode
}

func (ptn *PointerTypeNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitPointerTypeNode(ptn)
}
func (ptn *PointerTypeNode) typeNode()      {}
func (ptn *PointerTypeNode) String() string { return "*" + ptn.ElementType.String() }

// FunctionTypeNode represents a function type (e.g., fn(int, string) bool).
type FunctionTypeNode struct {
	ParameterTypes []TypeNode
	ReturnType     TypeNode // Optional
}

func (ftn *FunctionTypeNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitFunctionTypeNode(ftn)
}
func (ftn *FunctionTypeNode) typeNode()      {}
func (ftn *FunctionTypeNode) String() string { return "FunctionTypeNode" } // Placeholder

// OptionalTypeNode represents an optional type (e.g., ?int).
type OptionalTypeNode struct {
	ElementType TypeNode
}

func (otn *OptionalTypeNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitOptionalTypeNode(otn)
}
func (otn *OptionalTypeNode) typeNode()      {}
func (otn *OptionalTypeNode) String() string { return "?" + otn.ElementType.String() }

// AnonymousStructTypeNode represents an anonymous struct type definition.
type AnonymousStructTypeNode struct {
	Fields []*FieldDecl
}

func (astn *AnonymousStructTypeNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitAnonymousStructTypeNode(astn)
}
func (astn *AnonymousStructTypeNode) typeNode()      {}
func (astn *AnonymousStructTypeNode) String() string { return "struct { ... }" } // Placeholder

// AnonymousInterfaceTypeNode represents an anonymous interface type definition.
type AnonymousInterfaceTypeNode struct {
	Methods []*MethodSignatureNode
}

func (aitn *AnonymousInterfaceTypeNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitAnonymousInterfaceTypeNode(aitn)
}
func (aitn *AnonymousInterfaceTypeNode) typeNode()      {}
func (aitn *AnonymousInterfaceTypeNode) String() string { return "interface { ... }" } // Placeholder

// MethodSignatureNode represents a method signature in an interface type.
type MethodSignatureNode struct {
	Name       *IdentifierExpr
	Parameters []*ParameterDecl
	ReturnType TypeNode // Optional
}

func (msn *MethodSignatureNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitMethodSignatureNode(msn)
}
func (msn *MethodSignatureNode) String() string { return "MethodSignature: " + msn.Name.Name }

// PrimitiveKind defines the specific kind of a primitive type.
type PrimitiveKind int

const (
	IntKind PrimitiveKind = iota
	Float64Kind
	StringKind
	BoolKind
	RuneKind
)

func (pk PrimitiveKind) String() string {
	switch pk {
	case IntKind:
		return "int"
	case Float64Kind:
		return "float64"
	case StringKind:
		return "string"
	case BoolKind:
		return "bool"
	case RuneKind:
		return "rune"
	default:
		return "unknown_primitive_kind"
	}
}

// PrimitiveTypeNode represents a primitive type (e.g., int, bool).
type PrimitiveTypeNode struct {
	Kind PrimitiveKind
}

// NewPrimitiveType is a constructor for PrimitiveTypeNode.
func NewPrimitiveType(kind PrimitiveKind) *PrimitiveTypeNode {
	return &PrimitiveTypeNode{Kind: kind}
}

func (ptn *PrimitiveTypeNode) Accept(visitor Visitor) interface{} {
	return visitor.VisitPrimitiveTypeNode(ptn)
}
func (ptn *PrimitiveTypeNode) typeNode()      {}
func (ptn *PrimitiveTypeNode) String() string { return ptn.Kind.String() }
