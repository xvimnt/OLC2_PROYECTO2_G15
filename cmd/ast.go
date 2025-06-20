package main

import (
	"fmt"
	"os"
	"strconv"

	antlr "github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/xvimnt/OLC2_PROYECTO1_G15/ast"
	"github.com/xvimnt/OLC2_PROYECTO1_G15/parser"
)

// getNodeString is a helper to get a string representation of an AST node for logging.
func getNodeString(node interface{}) string {
	if node == nil {
		return "<nil_node>"
	}
	if expr, ok := node.(ast.Expression); ok {
		return expr.String()
	}
	// Add more types as needed, e.g., ast.Statement, ast.Declaration
	// For now, just return the type for unhandled ast node types.
	switch node.(type) {
	case ast.Statement:
		// If statements have a String() method, call it.
		// For now, just using type.
		return fmt.Sprintf("(%T)", node)
	case ast.Declaration:
		// If declarations have a String() method, call it.
		return fmt.Sprintf("(%T)", node)
	}
	return fmt.Sprintf("(%T_unknown_ast_type)", node)
}

// AstBuilder implements the ANTLR visitor pattern to build an AST
type AstBuilder struct {
	parser.BaseVLangCherryVisitor
	DebugMode bool
}

func NewAstBuilder(debugMode bool) *AstBuilder {
	return &AstBuilder{DebugMode: debugMode}
}

// ParseExpressionFromString takes a string, parses it as a VLang expression,
// and returns the corresponding ast.Expression node or an error.
func (v *AstBuilder) ParseExpressionFromString(expressionString string) (ast.Expression, error) {
	if v.DebugMode {
		fmt.Printf("AstBuilder.ParseExpressionFromString: Parsing '%s'\n", expressionString)
	}

	// Create an ANTLR input stream from the string
	input := antlr.NewInputStream(expressionString)
	lexer := parser.NewVLangCherryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewVLangCherryParser(stream)

	// TODO: Consider adding a more robust error listener to capture syntax errors specifically from this parse.
	// For now, relies on default console error listener or panics from ANTLR.
	p.BuildParseTrees = true // Enable parse tree construction

	// Call the parser rule for a single expression.
	// The 'expression' rule is typically the entry point for parsing expressions.
	exprCtx := p.Expression()

	// Use the existing AstBuilder (v) to visit the parsed expression context.
	// The Accept method will dispatch to the correct Visit... method in AstBuilder.
	result := exprCtx.Accept(v)

	if expr, ok := result.(ast.Expression); ok {
		if v.DebugMode {
			fmt.Printf("AstBuilder.ParseExpressionFromString: Successfully parsed to AST node: %s (%T)\n", expr.String(), expr)
		}
		return expr, nil
	}

	// If the result is not an ast.Expression, it's an unexpected parsing outcome or error.
	errMsg := fmt.Sprintf("failed to parse expression string '%s' into an ast.Expression. Got type %T, value: %+v", expressionString, result, result)
	if v.DebugMode {
		fmt.Println(errMsg)
	}
	if err, isErr := result.(error); isErr { // Check if the result itself is an error from a visitor method
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}
	return nil, fmt.Errorf(errMsg)
}

func (v *AstBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting Program")
	}
	program := &ast.Program{
		Declarations: []ast.Declaration{},
	}

	for _, declCtx := range ctx.AllTopLevelDeclaration() {
		declNode := declCtx.Accept(v)
		if declNode != nil {
			if decl, ok := declNode.(ast.Declaration); ok {
				program.Declarations = append(program.Declarations, decl)
			} else {
				if v.DebugMode {
					fmt.Fprintf(os.Stderr, "Warning: Visiting a top-level declaration did not return an ast.Declaration. Got %T\n", declNode)
				}
			}
		}
	}

	return program
}

func (v *AstBuilder) VisitTopLevelDeclaration(ctx *parser.TopLevelDeclarationContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting TopLevelDeclarationContext")
	}

	if fnCtx := ctx.FunctionDeclaration(); fnCtx != nil {
		return fnCtx.Accept(v)
	} else if structCtx := ctx.StructDeclaration(); structCtx != nil {
		return structCtx.Accept(v)
	} else if varCtx := ctx.VarDecl(); varCtx != nil { // Global variable declarations
		return varCtx.Accept(v)
	}

	// Only FunctionDeclaration, StructDeclaration, and VarDecl are handled here as top-level.
	// Other statement types (like IncDecStatement, ContinueStatement, standalone Blocks)
	// are handled via VisitStatement, which is called from within function bodies, loops, etc.

	if v.DebugMode {
		fmt.Fprintf(os.Stderr, "Warning: Unhandled TopLevelDeclarationContext: %s. Expected Function, Struct, or VarDecl.\n", ctx.GetText())
	}
	return nil
}

// VisitExpression handles the dispatch from a generic ExpressionContext to its specific alternative (LogicalOrExpr).
func (v *AstBuilder) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ExpressionContext")
	}

	// The 'expression' rule in the grammar is: 'expression: logicalOrExpr;'
	// So, ExpressionContext should have a LogicalOrExpr() method.
	logicalOrCtx := ctx.LogicalOrExpr()
	if logicalOrCtx == nil {
		if v.DebugMode { // Keep this specific debug for nil case, it's useful
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitExpression: LogicalOrExprContext is nil for ExpressionContext '%s'. Returning nil.\n", ctx.GetText())
		}
		return nil
	}

	// Call Accept on the specific LogicalOrExprContext to ensure VisitLogicalOrExpr is called.
	result := logicalOrCtx.Accept(v)
	return result
}

// --- Type Visitor Methods ---

func (v *AstBuilder) VisitTypePrimitive(ctx *parser.TypePrimitiveContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting TypePrimitive (type rule alternative): %s\n", ctx.GetText())
	}
	if ctx.PrimitiveType() == nil {
		fmt.Fprintf(os.Stderr, "Error: PrimitiveTypeContext is nil in TypePrimitiveContext: %s\n", ctx.GetText())
		return nil
	}
	return ctx.PrimitiveType().Accept(v) // Delegate to VisitPrimitiveType
}

func (v *AstBuilder) VisitIdentifierType(ctx *parser.IdentifierTypeContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting IdentifierType (type rule alternative): %s\n", ctx.GetText())
	}
	if ctx.IDENTIFIER() == nil {
		fmt.Fprintf(os.Stderr, "Error: IDENTIFIER is nil in IdentifierTypeContext: %s\n", ctx.GetText())
		return nil
	}
	return &ast.TypeName{Name: ctx.IDENTIFIER().GetText()}
}

func (v *AstBuilder) VisitPointerType(ctx *parser.PointerTypeContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting PointerType (type rule alternative): %s\n", ctx.GetText())
	}
	if ctx.Type_() == nil {
		fmt.Fprintf(os.Stderr, "Error: Type_ is nil in PointerTypeContext: %s\n", ctx.GetText())
		return nil
	}
	elemTypeResult := ctx.Type_().Accept(v)
	if elemType, ok := elemTypeResult.(ast.TypeNode); ok {
		return &ast.PointerTypeNode{ElementType: elemType}
	} else {
		fmt.Fprintf(os.Stderr, "Error: Visiting element type in PointerType did not return ast.TypeNode. Got %T for '%s'\n", elemTypeResult, ctx.Type_().GetText())
		return nil
	}
}

func (v *AstBuilder) VisitSliceType(ctx *parser.SliceTypeContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting SliceType (type rule alternative): %s\n", ctx.GetText())
	}
	if ctx.Type_() == nil {
		fmt.Fprintf(os.Stderr, "Error: Type_ is nil in SliceTypeContext: %s\n", ctx.GetText())
		return nil
	}
	elemTypeResult := ctx.Type_().Accept(v)
	if elemType, ok := elemTypeResult.(ast.TypeNode); ok {
		return &ast.SliceTypeNode{ElementType: elemType}
	} else {
		fmt.Fprintf(os.Stderr, "Error: Visiting element type in SliceType did not return ast.TypeNode. Got %T for '%s'\n", elemTypeResult, ctx.Type_().GetText())
		return nil
	}
}

// VisitPrimitiveType handles the primitiveType rule from the grammar.
// primitiveType: INT | FLOAT64 | STRING | BOOL | RUNE;
func (v *AstBuilder) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting PrimitiveTypeContext: %s\n", ctx.GetText())
	}

	var kind ast.PrimitiveKind
	var typeName string

	if ctx.INT() != nil {
		kind = ast.IntKind
		typeName = "int"
	} else if ctx.FLOAT64() != nil {
		kind = ast.Float64Kind
		typeName = "float64"
	} else if ctx.STRING() != nil {
		kind = ast.StringKind
		typeName = "string"
	} else if ctx.BOOL() != nil {
		kind = ast.BoolKind
		typeName = "bool"
	} else if ctx.RUNE() != nil {
		kind = ast.RuneKind
		typeName = "rune"
	} else {
		fmt.Fprintf(os.Stderr, "Error: Unknown primitive type in PrimitiveTypeContext: %s\n", ctx.GetText())
		return nil
	}

	if v.DebugMode {
		fmt.Printf("AstBuilder.VisitPrimitiveType: Determined kind %v for type %s\n", kind, typeName)
	}
	// Assuming ast.NewPrimitiveType exists as per memory 02162820-4e24-4f1b-bf4e-fd028ac80a1c
	return ast.NewPrimitiveType(kind)
}

func (v *AstBuilder) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting StructDeclarationContext: %s\n", ctx.GetText())
	}

	name := &ast.IdentifierExpr{Name: ctx.IDENTIFIER().GetText()}

	var fields []*ast.FieldDecl
	for _, fieldCtx := range ctx.AllFieldDeclaration() {
		if concreteFieldCtx, ok := fieldCtx.(*parser.FieldDeclarationContext); ok {
			fieldNode := v.VisitFieldDeclaration(concreteFieldCtx)
			if f, ok := fieldNode.(*ast.FieldDecl); ok {
				fields = append(fields, f)
			}
		}
	}

	return &ast.StructDecl{
		Name:   name,
		Fields: fields,
	}
}

func (v *AstBuilder) VisitReceiver(ctx *parser.ReceiverContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting ReceiverContext: %s\n", ctx.GetText())
	}

	if ctx.IDENTIFIER() == nil {
		fmt.Fprintf(os.Stderr, "Error: missing receiver name in function declaration. Found '%s'.\n", ctx.GetText())
		return nil
	}
	receiverName := &ast.IdentifierExpr{Name: ctx.IDENTIFIER().GetText()}

	if ctx.Type_() == nil {
		fmt.Fprintf(os.Stderr, "Error: missing type for receiver '%s' in function declaration.\n", receiverName.Name)
		return nil
	}

	var receiverType ast.TypeNode
	typeResult := ctx.Type_().Accept(v)
	if typeNode, ok := typeResult.(ast.TypeNode); ok {
		receiverType = typeNode
	} else {
		fmt.Fprintf(os.Stderr, "Error: could not parse type for receiver '%s'. Parser returned %T for type '%s'.\n", receiverName.Name, typeResult, ctx.Type_().GetText())
		return nil
	}

	return &ast.ReceiverDecl{
		Name: receiverName,
		Type: receiverType, // The type node itself will indicate if it's a pointer
	}
}

func (v *AstBuilder) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting FieldDeclarationContext: %s\n", ctx.GetText())
	}

	// V syntax is `fieldName type`. Grammar should be `fieldDeclaration: IDENTIFIER type_`.
	if ctx.IDENTIFIER() == nil {
		fmt.Fprintf(os.Stderr, "Error: missing field name in struct declaration. Found '%s'.\n", ctx.GetText())
		return nil
	}
	fieldName := &ast.IdentifierExpr{Name: ctx.IDENTIFIER().GetText()}

	if ctx.Type_() == nil {
		fmt.Fprintf(os.Stderr, "Error: missing type for field '%s' in struct declaration.\n", fieldName.Name)
		return nil
	}

	var fieldType ast.TypeNode
	typeResult := ctx.Type_().Accept(v) // Use Accept for dispatch
	if typeNode, ok := typeResult.(ast.TypeNode); ok {
		fieldType = typeNode
	} else if typeResult != nil {
		fmt.Fprintf(os.Stderr, "Error: could not parse type for struct field '%s'. Parser returned %T for type '%s'.\n", fieldName.Name, typeResult, ctx.Type_().GetText())
		return nil
	}

	return &ast.FieldDecl{
		Name: fieldName,
		Type: fieldType,
	}
}

func (v *AstBuilder) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting FunctionDeclarationContext") // Corresponds to ast.FunctionDecl
	}

	var funcNameNode *ast.IdentifierExpr
	if idNode := ctx.IDENTIFIER(); idNode != nil {
		funcNameNode = &ast.IdentifierExpr{Name: idNode.GetText()}
	} else {
		// This should ideally not happen if the grammar enforces a function name.

		// Return a placeholder or error node to avoid propagating nil further up.
		return &ast.FunctionDecl{Name: &ast.IdentifierExpr{Name: "<ERROR_FUNC_NO_NAME>"}}
	}

	var receiverDecl *ast.ReceiverDecl
	if receiverCtx := ctx.Receiver(); receiverCtx != nil {
		// Explicitly call VisitReceiver if it's a ReceiverContext
		if concreteReceiverCtx, ok := receiverCtx.(*parser.ReceiverContext); ok {
			receiverNodeRet := v.VisitReceiver(concreteReceiverCtx) // Call the specific visitor
			if rec, ok := receiverNodeRet.(*ast.ReceiverDecl); ok {
				receiverDecl = rec
			} else if receiverNodeRet != nil {
				fmt.Fprintf(os.Stderr, "Warning: VisitReceiver did not return *ast.ReceiverDecl, got %T\n", receiverNodeRet)
			}
		} else {
			// Fallback or error if it's not the expected concrete type, though ANTLR usually provides it.
			fmt.Fprintf(os.Stderr, "Warning: ctx.Receiver() returned an unexpected type %T, attempting generic visit.\n", receiverCtx)
			receiverNodeRet := v.Visit(receiverCtx) // Fallback to generic visit if type assertion fails (should not happen)
			if rec, ok := receiverNodeRet.(*ast.ReceiverDecl); ok {
				receiverDecl = rec
			}
		}
	}

	var parameters []*ast.ParameterDecl
	if paramsCtx := ctx.Parameters(); paramsCtx != nil {
		var paramsNodeRet interface{}
		if concreteParamsCtx, ok := paramsCtx.(*parser.ParametersContext); ok {
			paramsNodeRet = v.VisitParameters(concreteParamsCtx)
		} else {
			paramsNodeRet = v.Visit(paramsCtx) // Fallback
		}

		// VisitParameters is expected to return []*ast.ParameterDecl
		if paramsSlice, ok := paramsNodeRet.([]*ast.ParameterDecl); ok {
			parameters = paramsSlice
		} else if paramsNodeRet != nil { // Non-nil, but not []*ast.ParameterDecl
			// TODO: Log if paramsNodeRet is not nil but not []*ast.ParameterDecl
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitFunctionDeclaration: Expected []*ast.ParameterDecl from VisitParameters, got %T for '%s'\n", paramsNodeRet, paramsCtx.GetText())
			}
		} else { // paramsNodeRet is nil
			// TODO: Log if paramsNodeRet is nil (meaning VisitParameters returned nil)
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitFunctionDeclaration: VisitParameters returned nil for '%s'\n", paramsCtx.GetText())
			}
		}
	} else { // No ParametersContext (no parameters provided in source)
		parameters = []*ast.ParameterDecl{} // Ensure parameters is empty
	}

	var returnType ast.TypeNode
	if returnTypeCtx := ctx.ReturnType(); returnTypeCtx != nil {
		returnTypeNodeRet := v.Visit(returnTypeCtx)
		if rt, ok := returnTypeNodeRet.(ast.TypeNode); ok {
			returnType = rt
		} else if returnTypeNodeRet != nil {
			// TODO: Log if returnTypeNodeRet is not nil but not ast.TypeNode
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitFunctionDeclaration: Expected ast.TypeNode from VisitReturnType, got %T for '%s'\n", returnTypeNodeRet, returnTypeCtx.GetText())
			}
		} else {
			// TODO: Log if returnTypeNodeRet is nil (meaning VisitReturnType returned nil)
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitFunctionDeclaration: VisitReturnType returned nil for '%s'\n", returnTypeCtx.GetText())
			}
		}
	}

	var body *ast.BlockStmt
	if blockCtx := ctx.Block(); blockCtx != nil {

		var bodyNodeRet interface{}
		// Attempt to type assert to the concrete *parser.BlockContext and call VisitBlock directly
		if concreteBlockCtx, ok := blockCtx.(*parser.BlockContext); ok {

			bodyNodeRet = v.VisitBlock(concreteBlockCtx)
		} else { // If type assertion fails (blockCtx is known to be non-nil here)
			bodyNodeRet = v.Visit(blockCtx) // Fallback to original generic visit
		}

		if blk, ok := bodyNodeRet.(*ast.BlockStmt); ok {
			body = blk
		} else if bodyNodeRet != nil {
			// TODO: Log if bodyNodeRet is not nil but not *ast.BlockStmt
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitFunctionDeclaration: Expected *ast.BlockStmt from VisitBlock, got %T for '%s'\n", bodyNodeRet, blockCtx.GetText())
			}
		} else {
			// TODO: Log if bodyNodeRet is nil (meaning VisitBlock returned nil)
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitFunctionDeclaration: VisitBlock returned nil for '%s'\n", blockCtx.GetText())
			}
		}
	} else {

		// This can happen for function declarations without bodies (e.g. interface methods or forward declarations if supported)

	}

	if v.DebugMode {
		fmt.Printf("AstBuilder.VisitFunctionDeclaration: Populated %d parameters for function %s\n", len(parameters), funcNameNode.Name)
		for i, p := range parameters {
			if p != nil && p.Name != nil {
				typeName := "<nil_type>"
				if p.Type != nil {
					typeName = p.Type.String()
				}
				fmt.Printf("  Param %d: Name=%s, Type=%s\n", i, p.Name.Name, typeName)
			} else {
				fmt.Printf("  Param %d: <nil_param_or_name>\n", i)
			}
		}
	}

	return &ast.FunctionDecl{
		Receiver:   receiverDecl,
		Name:       funcNameNode,
		Parameters: parameters,
		ReturnType: returnType,
		Body:       body,
	}
}

func (v *AstBuilder) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting VarDeclContext") // Corresponds to ast.VarDecl
	}

	var varNameNode *ast.IdentifierExpr
	// Assumes IDENTIFIER() returns the TerminalNode for the variable's name.
	// This needs to match your VLangCherryParser.g4 rule for varDecl.
	// Example: if your rule is `varDecl: 'let' name=IDENTIFIER ...;` then `ctx.GetName()` or `ctx.name` might be used.
	// If IDENTIFIER is a direct child token: `ctx.IDENTIFIER()`
	if idTerminal := ctx.IDENTIFIER(); idTerminal != nil {
		nameToken := idTerminal.GetText()
		varNameNode = &ast.IdentifierExpr{Name: nameToken}
	} else {
		// This indicates a potential issue if an identifier is always expected.

		// Return a placeholder or error node to avoid propagating nil, though this node is invalid.
		return &ast.VarDecl{Name: &ast.IdentifierExpr{Name: "<ERROR_VAR_NO_NAME>"}}
	}

	// Assumes MUT() returns the TerminalNode for the 'mut' keyword, if present.
	isMutable := ctx.MUT() != nil

	var explicitType ast.TypeNode
	// Assumes Type_() returns the context for the explicit type specification (e.g., rule `type_`).
	if typeCtx := ctx.Type_(); typeCtx != nil { // This is ITypeContext
		if v.DebugMode {
			fmt.Printf("  VisitVarDecl: typeCtx for variable '%s' is NOT nil. Text: '%s'. Visiting it now.\n", varNameNode.Name, typeCtx.GetText())
		}
		typeNodeRet := typeCtx.Accept(v) // Use ANTLR visitor dispatch directly on the interface
		if tn, ok := typeNodeRet.(ast.TypeNode); ok {
			explicitType = tn
			if v.DebugMode {
				fmt.Printf("  VisitVarDecl: Visiting typeCtx for variable '%s' returned: (%T) %s\n", varNameNode.Name, explicitType, getNodeString(explicitType))
			}
		} else if typeNodeRet != nil {
			if v.DebugMode {
				fmt.Printf("  VisitVarDecl: Visiting typeCtx for variable '%s' returned non-TypeNode: (%T) %v\n", varNameNode.Name, typeNodeRet, typeNodeRet)
			}
		} else {
			if v.DebugMode {
				fmt.Printf("  VisitVarDecl: Visiting typeCtx for variable '%s' returned nil. Type context text: '%s'\n", varNameNode.Name, typeCtx.GetText())
			}
		}
	} else {
		if v.DebugMode {
			fmt.Printf("  VisitVarDecl: typeCtx for variable '%s' is nil.\n", varNameNode.Name)
		}
	}

	var initializer ast.Expression
	// Assumes Expression() returns the context for the initializer expression.
	// This needs to match your VLangCherryParser.g4 rule.
	// Example: `varDecl: ... '=' init=expression ;` then `ctx.GetInit()` or `ctx.init`
	// If it's just `Expression()`: `ctx.Expression()`
	if exprCtx := ctx.Expression(); exprCtx != nil {
		exprNodeRet := exprCtx.Accept(v) // Use Accept for correct dispatch to overridden visitor methods
		if v.DebugMode {
			fmt.Printf("  VisitVarDecl: exprCtx.Accept(v) for '%s' returned: (%T) %v\n", exprCtx.GetText(), exprNodeRet, exprNodeRet)
		}
		if expr, ok := exprNodeRet.(ast.Expression); ok {
			initializer = expr
		} else if exprNodeRet != nil {
			// TODO: Log if exprNodeRet is not nil but not ast.Expression
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  VisitVarDecl Warning: Initializer expression '%s' (type %T) was visited, but result (%T) %v is not ast.Expression\n", exprCtx.GetText(), exprCtx, exprNodeRet, exprNodeRet)
			}
		} else {
			// TODO: Log if exprNodeRet is nil (meaning Accept(v) returned nil)
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  VisitVarDecl Error: Initializer expression '%s' (type %T) was visited, but Accept(v) returned nil\n", exprCtx.GetText(), exprCtx)
			}
		}
	} else { // exprCtx IS nil
	}

	var isShortHand bool
	// Check for short-hand assignment (:=). The ANTLR-generated parser uses COLON_EQ for this token.
	isShortHand = ctx.COLON_EQ() != nil

	return &ast.VarDecl{
		Name:         varNameNode,
		IsMutable:    isMutable,
		ExplicitType: explicitType,
		Initializer:  initializer,
		IsShortHand:  isShortHand,
		Line:         ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
		Column:       ctx.GetStart().GetColumn(), // <-- Y ESTO
	}
}

// VisitType is removed as it's superseded by specific labeled alternative visitors.
// The ANTLR dispatch mechanism will call VisitTypePrimitive, VisitIdentifierType, etc.
// based on the #label in the grammar for the 'type' rule.

func (v *AstBuilder) VisitParameters(ctx *parser.ParametersContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ParametersContext")
	}
	var paramDecls []*ast.ParameterDecl
	// Assuming ParametersContext has a list of ParameterDeclContexts.
	// The exact method to get these depends on your ANTLR grammar for 'parameters'.
	// Common ways are ctx.AllParameterDecl() or iterating ctx.GetChildren().
	// Let's assume ctx.AllParameterDecl() exists and returns []parser.IParameterDeclContext
	for _, iParamDeclCtx := range ctx.AllParameterDecl() {
		var paramNodeRet interface{}
		if concreteParamDeclCtx, ok := iParamDeclCtx.(*parser.ParameterDeclContext); ok {
			paramNodeRet = v.VisitParameterDecl(concreteParamDeclCtx)
		} else {
			paramNodeRet = v.Visit(iParamDeclCtx) // Fallback
		}

		if paramNode, ok := paramNodeRet.(*ast.ParameterDecl); ok {
			paramDecls = append(paramDecls, paramNode)
		} else if paramNodeRet != nil {
			// Log that we got something, but it wasn't an ast.ParameterDecl
			if v.DebugMode {
				fmt.Printf("AstBuilder.VisitParameters: Expected *ast.ParameterDecl, but got %T for context %s\n", paramNodeRet, iParamDeclCtx.GetText())
			}
		} else {
			// Log that we got nil from visiting the parameter declaration context
			if v.DebugMode {
				fmt.Printf("AstBuilder.VisitParameters: Got nil from visiting parameter declaration context %s\n", iParamDeclCtx.GetText())
			}
		}
	}
	return paramDecls
}

func (v *AstBuilder) VisitParameterDecl(ctx *parser.ParameterDeclContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting ParameterDeclContext: %s\n", ctx.GetText())
	}

	var paramNameNode *ast.IdentifierExpr
	// Assuming ParameterDeclContext has IDENTIFIER() for name and Type_() for type context.
	// These might need to be ctx.Name().GetText() or ctx.Id().GetText() etc. depending on grammar labels.
	if idTerminalNode := ctx.IDENTIFIER(); idTerminalNode != nil {
		paramNameNode = &ast.IdentifierExpr{Name: idTerminalNode.GetText()}
	} else {
		fmt.Fprintf(os.Stderr, "Error: Parameter declaration is missing a name in context: %s\n", ctx.GetText())
		return &ast.ParameterDecl{Name: &ast.IdentifierExpr{Name: "<ERROR_PARAM_NO_NAME>"}}
	}

	var paramType ast.TypeNode
	// Assuming ParameterDeclContext has a Type_() method returning a parser.ITypeContext
	if typeCtx := ctx.Type_(); typeCtx != nil { // Common ANTLR naming if rule is 'type'
		if v.DebugMode {
			fmt.Printf("AstBuilder.VisitParameterDecl: Visiting type for param '%s' using Accept. Type context is %T: %s\n", paramNameNode.Name, typeCtx, typeCtx.GetText())
		}
		paramTypeResult := typeCtx.Accept(v)
		if pt, ok := paramTypeResult.(ast.TypeNode); ok {
			paramType = pt
		} else {
			if v.DebugMode {
				fmt.Printf("AstBuilder.VisitParameterDecl: Visiting type context for param '%s' did not return an ast.TypeNode, but %T\n", paramNameNode.Name, paramTypeResult)
			}
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: Parameter declaration is missing a type specification for param %s in context: %s\n", paramNameNode.Name, ctx.GetText())
	}

	return &ast.ParameterDecl{
		Name:   paramNameNode,
		Type:   paramType,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (v *AstBuilder) VisitLogicalOrExpr(ctx *parser.LogicalOrExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting LogicalOrExprContext")
	}

	leftResult := ctx.LogicalAndExpr(0).Accept(v)
	left, ok := leftResult.(ast.Expression)
	if !ok || left == nil {

		return nil
	}

	// Iterate over subsequent operands and operators
	// There will be one less OR operator than LogicalAndExpr operands
	for i := 1; i < len(ctx.AllLogicalAndExpr()); i++ {
		operatorNode := ctx.OR(i - 1) // OR token for the current operation
		if operatorNode == nil {

			return nil
		}
		operator := operatorNode.GetText()

		rightResult := ctx.LogicalAndExpr(i).Accept(v)
		right, ok := rightResult.(ast.Expression)
		if !ok || right == nil {

			return nil
		}
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
		if v.DebugMode {
			fmt.Printf("  Built BinaryExpr: %s\n", left.String())
		}
	}

	return left
}

func (v *AstBuilder) VisitLogicalAndExpr(ctx *parser.LogicalAndExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting LogicalAndExprContext")
	}

	leftResult := ctx.EqualityExpr(0).Accept(v)
	left, ok := leftResult.(ast.Expression)
	if !ok || left == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  LogicalAndExpr: Left operand '%s' did not evaluate to ast.Expression or was nil. Got %T\n", ctx.EqualityExpr(0).GetText(), leftResult)
		}
		return nil
	}

	for i := 1; i < len(ctx.AllEqualityExpr()); i++ {
		operatorNode := ctx.AND(i - 1)
		if operatorNode == nil {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  LogicalAndExpr: Operator AND(%d) is nil. Text: '%s'\n", i-1, ctx.GetText())
			}
			return nil
		}
		operator := operatorNode.GetText()

		rightResult := ctx.EqualityExpr(i).Accept(v)
		right, ok := rightResult.(ast.Expression)
		if !ok || right == nil {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  LogicalAndExpr: Right operand '%s' for operator '%s' did not evaluate to ast.Expression or was nil. Got %T\n",
					ctx.EqualityExpr(i).GetText(), operator, rightResult)
			}
			return nil
		}
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}

	return left
}

func (v *AstBuilder) VisitEqualityExpr(ctx *parser.EqualityExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting EqualityExprContext")
	}

	leftResult := ctx.RelationalExpr(0).Accept(v)
	left, ok := leftResult.(ast.Expression)
	if !ok || left == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  EqualityExpr: Left operand '%s' did not evaluate to ast.Expression or was nil. Got %T\n", ctx.RelationalExpr(0).GetText(), leftResult)
		}
		return nil
	}

	// The grammar is equalityExpr: relationalExpr ((EQ | NE) relationalExpr)*;
	// So, we iterate based on the number of relationalExpr beyond the first one.
	for i := 1; i < len(ctx.AllRelationalExpr()); i++ {
		var operator string
		// The operator (EQ or NE) is the child at index 2*(i-1) + 1
		opNodeIndex := 2*(i-1) + 1
		if opNodeIndex < 0 || opNodeIndex >= len(ctx.GetChildren()) {
			return nil
		}
		opNode := ctx.GetChild(opNodeIndex)
		if terminalNode, ok := opNode.(antlr.TerminalNode); ok {
			operator = terminalNode.GetText()
			// Optionally, verify token type:
			// tokenType := terminalNode.GetSymbol().GetType()
			// if tokenType != parser.VLangCherryLexerEQ && tokenType != parser.VLangCherryLexerNE {
			// 	if v.DebugMode {
			// 		fmt.Fprintf(os.Stderr, "  EqualityExpr: Operator token at index %d has unexpected type %d. Text: '%s'\n", opNodeIndex, tokenType, ctx.GetText())
			// 	}
			// 	return nil
			// }
		} else {
			return nil
		}

		rightResult := ctx.RelationalExpr(i).Accept(v)
		right, ok := rightResult.(ast.Expression)
		if !ok || right == nil {
			return nil
		}
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
		if v.DebugMode {
			fmt.Printf("  Built BinaryExpr: %s\n", left.String())
		}
	}

	return left
}

func (v *AstBuilder) VisitRelationalExpr(ctx *parser.RelationalExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting RelationalExprContext")
	}

	leftResult := ctx.AdditiveExpr(0).Accept(v)
	left, ok := leftResult.(ast.Expression)
	if !ok || left == nil {
		return nil
	}

	// Grammar: relationalExpr: additiveExpr ((LT | LE | GT | GE) additiveExpr)*;
	for i := 1; i < len(ctx.AllAdditiveExpr()); i++ {
		var operator string
		// The operator (LT, LE, GT, GE) is the child at index 2*(i-1) + 1
		opNodeIndex := 2*(i-1) + 1
		if opNodeIndex < 0 || opNodeIndex >= len(ctx.GetChildren()) {
			return nil
		}
		opNode := ctx.GetChild(opNodeIndex)
		if terminalNode, ok := opNode.(antlr.TerminalNode); ok {
			operator = terminalNode.GetText()
		} else {
			return nil
		}

		rightResult := ctx.AdditiveExpr(i).Accept(v)
		right, ok := rightResult.(ast.Expression)
		if !ok || right == nil {
			return nil
		}
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
		if v.DebugMode {
			fmt.Printf("  Built BinaryExpr: %s\n", left.String())
		}
	}

	return left
}

func (v *AstBuilder) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting AdditiveExprContext")
	}

	leftResult := ctx.MultiplicativeExpr(0).Accept(v)
	currentLeft, ok := leftResult.(ast.Expression)
	if !ok || currentLeft == nil {
		return nil
	}

	// Loop for operators and subsequent operands
	// Assumes grammar like: multiplicativeExpr ( (PLUS | MINUS) multiplicativeExpr )*
	for i := 1; i < len(ctx.AllMultiplicativeExpr()); i++ {
		// Operator is child at index 2*i - 1 relative to the start of the AdditiveExpr rule's children sequence
		opNodeIndex := 2*i - 1
		opNode := ctx.GetChild(opNodeIndex)
		operator := ""
		if terminalNode, ok := opNode.(antlr.TerminalNode); ok {
			operator = terminalNode.GetText()
		} else {
			return nil
		}

		rightResult := ctx.MultiplicativeExpr(i).Accept(v)
		right, ok := rightResult.(ast.Expression)
		if !ok || right == nil {
			return nil
		}
		currentLeft = &ast.BinaryExpr{
			Left:     currentLeft,
			Operator: operator,
			Right:    right,
			Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}

	return currentLeft
}

func (v *AstBuilder) VisitMultiplicativeExpr(ctx *parser.MultiplicativeExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting MultiplicativeExprContext")
	}

	leftResult := ctx.UnaryExpr(0).Accept(v)
	left, ok := leftResult.(ast.Expression)
	if !ok || left == nil {
		return nil
	}

	// Grammar: multiplicativeExpr: unaryExpr ((MUL | DIV | MOD) unaryExpr)*;
	for i := 1; i < len(ctx.AllUnaryExpr()); i++ {
		var operator string
		if ctx.MUL(i-1) != nil {
			operator = ctx.MUL(i - 1).GetText()
		} else if ctx.DIV(i-1) != nil {
			operator = ctx.DIV(i - 1).GetText()
		} else if ctx.MOD(i-1) != nil {
			operator = ctx.MOD(i - 1).GetText()
		} else {
			return nil
		}
		rightResult := ctx.UnaryExpr(i).Accept(v)
		right, ok := rightResult.(ast.Expression)
		if !ok || right == nil {
			return nil
		}
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}

	return left
}

func (v *AstBuilder) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting UnaryExprContext: '%s'\n", ctx.GetText())
	}

	// Alternative 1: unaryExpr: primaryExpr
	if primCtx := ctx.PrimaryExpr(); primCtx != nil {
		if v.DebugMode {
			typeName := fmt.Sprintf("%T", primCtx)
			// Check if it's a composite literal and log specific type if so
			if _, okIsTyped := primCtx.(*parser.TypedCompositeLitExprContext); okIsTyped {
				typeName = "TypedCompositeLitExprContext"
			} else if _, okIsInferred := primCtx.(*parser.InferredCompositeLitExprContext); okIsInferred {
				typeName = "InferredCompositeLitExprContext"
			} else {
				// Log that it's not one of the composite literal types, and show actual type
				// Ensure this specific log is indented if it's inside the primCtx block's DebugMode check
				fmt.Printf("    primCtx IS NOT a TypedCompositeLitExprContext or InferredCompositeLitExprContext. Actual type: %s\n", typeName)
			}
			fmt.Printf("  VisitUnaryExpr -> PrimaryExpr branch. primCtx type: %s, text: '%s'\n", typeName, primCtx.GetText())
		}

		// DIAGNOSTIC: Explicitly call VisitParenExpr if context is ParenExprContext
		if parenCtx, ok := primCtx.(*parser.ParenExprContext); ok {
			if v.DebugMode {
				fmt.Printf("  VisitUnaryExpr: primCtx is ParenExprContext, EXPLICITLY calling VisitParenExpr for '%s'\n", parenCtx.GetText())
			}
			// Directly call VisitParenExpr and return its result
			explicitResult := v.VisitParenExpr(parenCtx)
			if v.DebugMode {
				fmt.Printf("    v.VisitParenExpr(parenCtx) EXPLICIT result type: %T, value: %s\n", explicitResult, getNodeString(explicitResult))
			}
			return explicitResult
		}

		// Original path if not ParenExprContext
		result := primCtx.Accept(v)
		if v.DebugMode {
			fmt.Printf("    primCtx.Accept(v) result type: %T, value: %s\n", result, getNodeString(result))
			fmt.Printf("    primCtx.Accept(v) result type: %T, value: %+v\n", result, result)
		}
		return result
	}

	// Alternative 2: unaryExpr: (NOT | SUB | MUL | AMPERSAND) unaryExpr
	if v.DebugMode {
		fmt.Println("  VisitUnaryExpr -> Operator branch")
	}
	var operator string
	if ctx.NOT() != nil {
		operator = "!"
	} else if ctx.SUB() != nil {
		operator = "-"
	} else if ctx.MUL() != nil {
		operator = "*"
	} else if ctx.AMPERSAND() != nil {
		operator = "&"
	} else {
		if v.DebugMode {
			fmt.Println("  VisitUnaryExpr: No PrimaryExpr and no known operator. Returning nil.")
		}
		return nil // Should not happen if grammar is well-defined for UnaryExpr
	}

	operandRecursiveCtx := ctx.UnaryExpr() // This gets the nested UnaryExprContext
	if operandRecursiveCtx == nil {
		if v.DebugMode {
			fmt.Printf("  VisitUnaryExpr: Operator '%s' present, but recursive UnaryExpr is nil. Returning nil.\n", operator)
		}
		return nil
	}

	if v.DebugMode {
		fmt.Printf("  VisitUnaryExpr: Operator '%s', dispatching to Accept for recursive UnaryExpr: '%s'\n", operator, operandRecursiveCtx.GetText())
	}
	operandResult := operandRecursiveCtx.Accept(v) // Visit the nested UnaryExpr
	if v.DebugMode {
		fmt.Printf("    operandRecursiveCtx.Accept(v) result type: %T, value: %+v\n", operandResult, operandResult)
	}

	opExpr, ok := operandResult.(ast.Expression)

	if !ok || opExpr == nil {
		if v.DebugMode {
			fmt.Printf("  VisitUnaryExpr: Recursive operand result is not ast.Expression or is nil. Type: %T, Value: %+v. Returning nil.\n", operandResult, operandResult)
		}
		return nil
	}

	unaryNode := &ast.UnaryExpr{Operator: operator, Right: opExpr}
	if v.DebugMode {
		fmt.Printf("  VisitUnaryExpr: Constructed UnaryNode: %+v\n", unaryNode)
	}
	return unaryNode
}

func (v *AstBuilder) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting CallExprContext")
	}

	// Visit the function part (which is a PrimaryExpr itself, e.g., an identifier)
	// Use Accept for proper dispatch to specific primary expression visitors (e.g. VisitIdentifierExpr)
	functionResult := ctx.PrimaryExpr().Accept(v)
	fnExpr, ok := functionResult.(ast.Expression)
	if !ok || fnExpr == nil {
		return nil
	}

	var args []ast.Expression
	// Visit the arguments if ExpressionList is present
	if exprListCtx := ctx.ExpressionList(); exprListCtx != nil {
		argsResult := exprListCtx.Accept(v) // Dispatch to VisitExpressionList
		if argsRes, okArgs := argsResult.([]ast.Expression); okArgs {
			args = argsRes
		} else if argsResult != nil { // Non-nil, but not []ast.Expression
			args = []ast.Expression{} // Ensure args is empty
		} else { // argsResult is nil
			args = []ast.Expression{} // Ensure args is empty
		}
	} else { // No ExpressionList context (no arguments provided in source)
		args = []ast.Expression{} // Ensure args is empty
	}

	return &ast.CallExpr{
		Function:  fnExpr,
		Arguments: args,
		Line:      ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
		Column:    ctx.GetStart().GetColumn(), // <-- Y ESTO
	}
}

// VisitTypeOfExpr handles the '#TypeOfExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: TYPEOF L_PAREN expression R_PAREN #TypeOfExpr ;
func (v *AstBuilder) VisitTypeOfExpr(ctx *parser.TypeOfExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting TypeOfExpr: %s\n", ctx.GetText())
	}

	if ctx.Expression() == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitTypeOfExpr: ExpressionContext is nil for TypeOfExpr '%s'. Returning nil.\n", ctx.GetText())
		}
		return nil
	}

	exprNode := ctx.Expression().Accept(v)
	if expr, ok := exprNode.(ast.Expression); ok {
		return &ast.TypeOfExpr{Expression: expr}
	}

	if v.DebugMode {
		fmt.Fprintf(os.Stderr, "AstBuilder.VisitTypeOfExpr: Visiting expression in TypeOf did not return ast.Expression. Got %T for '%s'\n", exprNode, ctx.Expression().GetText())
	}
	return nil // Or an error node
}

// VisitArrayLiteralExpr handles the '#ArrayLiteralExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: L_SQUARE expressionList? R_SQUARE #ArrayLiteralExpr ;
func (v *AstBuilder) VisitArrayLiteralExpr(ctx *parser.ArrayLiteralExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting ArrayLiteralExpr: %s\n", ctx.GetText())
	}

	var elements []ast.Expression
	if exprListCtx := ctx.ExpressionList(); exprListCtx != nil {
		exprListResult := exprListCtx.Accept(v) // This should return []ast.Expression
		if exprSlice, ok := exprListResult.([]ast.Expression); ok {
			elements = exprSlice
		} else if exprListResult != nil {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitArrayLiteralExpr: ExpressionList did not return []ast.Expression. Got %T for '%s'\n", exprListResult, exprListCtx.GetText())
			}
			elements = []ast.Expression{}
		} else {
			// exprListResult is nil (e.g. VisitExpressionList returned nil)
			elements = []ast.Expression{}
		}
	} else {
		// No expressionList, so it's an empty array literal like `[]`
		elements = []ast.Expression{}
	}

	return &ast.ArrayLiteral{
		Elements: elements,
		Line:     ctx.GetStart().GetLine(),
		Column:   ctx.GetStart().GetColumn(),
	}
}

// Specific visitor for labeled alternative #IdentifierExpr in primaryExpr
func (v *AstBuilder) VisitIdentifierExpr(ctx *parser.IdentifierExprContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting IdentifierExprContext")
	}
	return &ast.IdentifierExpr{
		Name:   ctx.IDENTIFIER().GetText(),
		Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
		Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
	}
}

// VisitCompositeLitExpr handles the '#CompositeLitExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: ... | compositeLit #CompositeLitExpr ;
// compositeLit: type_ L_CURLY literalValue? R_CURLY;
func (v *AstBuilder) VisitTypedCompositeLitExpr(ctx *parser.TypedCompositeLitExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting TypedCompositeLitExprContext: '%s'\n", ctx.GetText())
	}

	actualCompositeLitCtx := ctx.CompositeLit()
	if actualCompositeLitCtx == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  VisitTypedCompositeLitExpr: actualCompositeLitCtx (from ctx.CompositeLit()) is nil for '%s'. Returning nil.\n", ctx.GetText())
		}
		return nil
	}

	// 1. Visit the type of the literal
	typeCtx := actualCompositeLitCtx.Type_()
	if typeCtx == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  VisitTypedCompositeLitExpr: Type_ context is nil for '%s'. Returning nil.\n", ctx.GetText())
		}
		return nil
	}

	typeNodeResult := typeCtx.Accept(v)
	astTypeNode, ok := typeNodeResult.(ast.TypeNode)
	if !ok || astTypeNode == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  VisitTypedCompositeLitExpr: Visiting Type_ context for '%s' did not return ast.TypeNode. Got %T: %+v. Returning nil.\n", typeCtx.GetText(), typeNodeResult, typeNodeResult)
		}
		return nil
	}
	if v.DebugMode {
		fmt.Printf("    VisitTypedCompositeLitExpr: Parsed TypeNode: %s\n", astTypeNode.String())
	}

	// 2. Visit the elements of the literal
	var elements []*ast.CompositeElement
	elementListRuleCtx := actualCompositeLitCtx.ElementList()

	if elementListRuleCtx == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  VisitTypedCompositeLitExpr: ElementList context is nil for '%s'. Creating empty elements slice.\n", ctx.GetText())
		}
		elements = []*ast.CompositeElement{}
	} else {
		// Logic for parsing elements is now robust for both keyed and non-keyed elements.
		exprListCtx := elementListRuleCtx.ExpressionList()
		keyedElementCtxs := elementListRuleCtx.AllKeyedElement()

		if exprListCtx != nil {
			for _, exprCtx := range exprListCtx.AllExpression() {
				elementResult := exprCtx.Accept(v)
				astElement, castOk := elementResult.(ast.Expression)
				if !castOk || astElement == nil {
					if v.DebugMode {
						fmt.Fprintf(os.Stderr, "    VisitTypedCompositeLitExpr: Element expression '%s' did not evaluate to ast.Expression. Got %T: %+v. Returning nil.\n", exprCtx.GetText(), elementResult, elementResult)
					}
					return nil
				}
				compositeElement := &ast.CompositeElement{Value: astElement, Key: nil}
				elements = append(elements, compositeElement)
			}
		} else if len(keyedElementCtxs) > 0 {
			for _, keyedCtx := range keyedElementCtxs {
				keyNode := &ast.IdentifierExpr{Name: keyedCtx.IDENTIFIER().GetText()}
				valueResult := keyedCtx.Expression().Accept(v)
				astValue, castOk := valueResult.(ast.Expression)
				if !castOk || astValue == nil {
					if v.DebugMode {
						fmt.Fprintf(os.Stderr, "    VisitTypedCompositeLitExpr: Keyed element value '%s' did not evaluate to ast.Expression. Got %T: %+v. Returning nil.\n", keyedCtx.Expression().GetText(), valueResult, valueResult)
					}
					return nil
				}
				compositeElement := &ast.CompositeElement{Key: keyNode, Value: astValue}
				elements = append(elements, compositeElement)
			}
		} else {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  VisitTypedCompositeLitExpr: ElementList was present but neither ExpressionList nor KeyedElements found for '%s'. Assuming empty elements.\n", elementListRuleCtx.GetText())
			}
			elements = []*ast.CompositeElement{}
		}
	}

	compositeLitNode := &ast.CompositeLiteralExpr{
		Type:     astTypeNode,
		Elements: elements,
	}

	if v.DebugMode {
		fmt.Printf("  VisitTypedCompositeLitExpr: Constructed CompositeLiteralExpr: %+v with %d elements. Returning node.\n", compositeLitNode, len(elements))
	}
	return compositeLitNode
}

// VisitInferredCompositeLitExpr handles the '#InferredCompositeLitExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: ... | L_BRACE elementList? R_BRACE #InferredCompositeLitExpr ;
// This is for literals like {1, 2, 3} or {key: "value"} where the type is inferred.
func (v *AstBuilder) VisitInferredCompositeLitExpr(ctx *parser.InferredCompositeLitExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting InferredCompositeLitExprContext: '%s'\n", ctx.GetText())
	}

	// For an inferred composite literal, there is no explicit type node.
	// The ast.CompositeLiteralExpr.Type field will be nil.
	// The type will be determined later during semantic analysis/type checking.

	var elements []*ast.CompositeElement
	elementListRuleCtx := ctx.ElementList() // Get the ElementListContext directly from the InferredCompositeLitExprContext

	if elementListRuleCtx == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  VisitInferredCompositeLitExpr: ElementList context is nil for '%s'. Creating empty elements slice.\n", ctx.GetText())
		}
		elements = []*ast.CompositeElement{}
	} else {
		// Logic for parsing elements is similar to VisitTypedCompositeLitExpr
		// For slice/array literals, we expect an ExpressionList within the ElementList.
		// For map/struct literals, we would look for KeyedElements.
		exprListCtx := elementListRuleCtx.ExpressionList()
		keyedElementCtxs := elementListRuleCtx.AllKeyedElement() // Check for keyed elements

		if exprListCtx != nil {
			for _, exprCtx := range exprListCtx.AllExpression() { // exprCtx is parser.IExpressionContext
				elementResult := exprCtx.Accept(v)
				astElement, castOk := elementResult.(ast.Expression)
				if !castOk || astElement == nil {
					if v.DebugMode {
						fmt.Fprintf(os.Stderr, "    VisitInferredCompositeLitExpr: Element expression '%s' did not evaluate to ast.Expression. Got %T: %+v. Returning nil.\n", exprCtx.GetText(), elementResult, elementResult)
					}
					return nil // Be strict
				}
				compositeElement := &ast.CompositeElement{Value: astElement, Key: nil}
				elements = append(elements, compositeElement)
				if v.DebugMode {
					fmt.Printf("      VisitInferredCompositeLitExpr: Added value element: %s\n", astElement.String())
				}
			}
		} else if len(keyedElementCtxs) > 0 {
			for _, keyedCtx := range keyedElementCtxs { // keyedCtx is parser.IKeyedElementContext
				keyNode := &ast.IdentifierExpr{Name: keyedCtx.IDENTIFIER().GetText()}
				valueResult := keyedCtx.Expression().Accept(v)
				astValue, castOk := valueResult.(ast.Expression)
				if !castOk || astValue == nil {
					if v.DebugMode {
						fmt.Fprintf(os.Stderr, "    VisitInferredCompositeLitExpr: Keyed element value '%s' did not evaluate to ast.Expression. Got %T: %+v. Returning nil.\n", keyedCtx.Expression().GetText(), valueResult, valueResult)
					}
					return nil // Be strict
				}
				compositeElement := &ast.CompositeElement{Key: keyNode, Value: astValue}
				elements = append(elements, compositeElement)
				if v.DebugMode {
					fmt.Printf("      VisitInferredCompositeLitExpr: Added keyed element: %s: %s\n", keyNode.String(), astValue.String())
				}
			}
		} else {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  VisitInferredCompositeLitExpr: ElementList was present but neither ExpressionList nor KeyedElements found for '%s'. Assuming empty elements.\n", elementListRuleCtx.GetText())
			}
			elements = []*ast.CompositeElement{}
		}
	}

	compositeLitNode := &ast.CompositeLiteralExpr{
		Type:     nil, // Type is inferred, so set to nil in AST initially
		Elements: elements,
	}

	if v.DebugMode {
		fmt.Printf("  VisitInferredCompositeLitExpr: Constructed CompositeLiteralExpr (inferred type): %+v with %d elements. Returning node.\n", compositeLitNode, len(elements))
	}
	return compositeLitNode
}

// VisitLiteralExpr handles the 'primaryExpr: literal #LiteralExpr' rule.
func (v *AstBuilder) VisitLiteralExpr(ctx *parser.LiteralExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting LiteralExprContext: %s\n", ctx.GetText())
	}
	if ctx.Literal() == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitLiteralExpr: LiteralContext is nil for LiteralExpr '%s'. Returning nil.\n", ctx.GetText())
		}
		return nil
	}
	// Delegate to VisitLiteral, which handles all kinds of literals.
	// VisitLiteral is expected to return an ast.Expression (e.g., *ast.IntegerLiteral, *ast.StringLiteral).
	litResult := ctx.Literal().Accept(v)
	if expr, ok := litResult.(ast.Expression); ok {
		if v.DebugMode {
			// getNodeString is a helper defined in ast.go (or should be)
			fmt.Printf("AstBuilder.VisitLiteralExpr: Literal().Accept(v) result type: %T, value: %s\n", expr, getNodeString(expr))
		}
		return expr
	}
	if v.DebugMode {
		fmt.Fprintf(os.Stderr, "AstBuilder.VisitLiteralExpr: Visiting literal did not return an ast.Expression. Got %T for '%s'. Returning nil.\n", litResult, ctx.Literal().GetText())
	}
	return nil
}

// VisitParenExpr handles the '#ParenExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: L_PAREN expression R_PAREN #ParenExpr ;
func (v *AstBuilder) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting ParenExprContext: %s\n", ctx.GetText())
	}

	if ctx.Expression() == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitParenExpr: Inner ExpressionContext is nil for ParenExpr '%s'. Returning nil.\n", ctx.GetText())
		}
		return nil
	}

	// Visit the inner expression
	innerExprResult := ctx.Expression().Accept(v)

	if innerExpr, ok := innerExprResult.(ast.Expression); ok {
		if v.DebugMode {
			fmt.Printf("AstBuilder.VisitParenExpr: Successfully parsed inner expression to AST node: %s (%T)\n", innerExpr.String(), innerExpr)
		}
		return innerExpr
	}

	if v.DebugMode {
		errMsg := fmt.Sprintf("AstBuilder.VisitParenExpr: Inner expression of '%s' did not evaluate to ast.Expression. Got type %T, value: %+v. Returning nil.", ctx.GetText(), innerExprResult, innerExprResult)
		fmt.Fprintf(os.Stderr, "%s\n", errMsg)
	}
	return nil
}

// VisitExpressionList handles the 'expressionList' rule.
// It should return a slice of ast.Expression ([]ast.Expression).
func (v *AstBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ExpressionListContext")
	}
	var exprs []ast.Expression
	for _, exprCtx := range ctx.AllExpression() { // exprCtx is parser.IExpressionContext
		exprNode := exprCtx.Accept(v) // Use Accept for robust dispatch
		if expr, ok := exprNode.(ast.Expression); ok {
			exprs = append(exprs, expr)
		} else {
			// Consider returning an error or a placeholder if an expression is not valid
		}
	}
	return exprs
}

// TODO: Implement visitors for other PrimaryExpr alternatives as needed:
// VisitIndexAccessExpr handles the '#IndexAccessExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: primaryExpr L_SQUARE expression R_SQUARE #IndexAccessExpr ;
func (v *AstBuilder) VisitIndexAccessExpr(ctx *parser.IndexAccessExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting IndexAccessExprContext: %s\n", ctx.GetText())
	}

	var receiver ast.Expression
	// In 'primaryExpr L_SQUARE expression R_SQUARE', ctx.PrimaryExpr() is the part to the left of '['
	if receiverCtx := ctx.PrimaryExpr(); receiverCtx != nil {
		receiverNode := receiverCtx.Accept(v)
		if r, ok := receiverNode.(ast.Expression); ok {
			receiver = r
		} else {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  IndexAccessExpr: Receiver expression '%s' did not evaluate to ast.Expression. Got %T: %s\n", receiverCtx.GetText(), receiverNode, getNodeString(receiverNode))
			}
			return nil
		}
	} else {
		// This case should ideally not be reached if grammar enforces a primaryExpr before '['
		if v.DebugMode {
			fmt.Fprintln(os.Stderr, "  IndexAccessExpr: Critical error - Receiver (PrimaryExpr) context is nil.")
		}
		return nil
	}

	var index ast.Expression
	// ctx.Expression() is the part between '[' and ']'
	if indexCtx := ctx.Expression(); indexCtx != nil {
		indexNode := indexCtx.Accept(v)
		if i, ok := indexNode.(ast.Expression); ok {
			index = i
		} else {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "  IndexAccessExpr: Index expression '%s' did not evaluate to ast.Expression. Got %T: %s\n", indexCtx.GetText(), indexNode, getNodeString(indexNode))
			}
			return nil
		}
	} else {
		// This case should ideally not be reached if grammar enforces an expression within '[]'
		if v.DebugMode {
			fmt.Fprintln(os.Stderr, "  IndexAccessExpr: Critical error - Index (Expression) context is nil.")
		}
		return nil
	}

	// Ensure both parts were successfully parsed
	if receiver == nil || index == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "  IndexAccessExpr: Failed to build receiver (%s) or index (%s) for '%s'. Returning nil.\n", getNodeString(receiver), getNodeString(index), ctx.GetText())
		}
		return nil
	}

	node := &ast.IndexAccessExpr{
		Receiver: receiver,
		Index:    index,
		Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
		Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
	}
	if v.DebugMode {
		fmt.Printf("  IndexAccessExpr: Constructed node for '%s': %s (Receiver: %s [%T], Index: %s [%T])\n", ctx.GetText(), node.String(), receiver.String(), receiver, index.String(), index)
	}
	return node
}

// VisitFieldAccessExpr handles the '#FieldAccessExpr' labeled alternative for 'primaryExpr'.
// Grammar: primaryExpr: primaryExpr DOT IDENTIFIER #FieldAccessExpr ;
func (v *AstBuilder) VisitFieldAccessExpr(ctx *parser.FieldAccessExprContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting FieldAccessExprContext: %s\n", ctx.GetText())
	}

	// The first PrimaryExpr() call gives the receiver part of the field access.
	receiverResult := ctx.PrimaryExpr().Accept(v)
	receiver, ok := receiverResult.(ast.Expression)
	if !ok {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitFieldAccessExpr: Receiver expression '%s' did not evaluate to ast.Expression. Got %T: %v\n", ctx.PrimaryExpr().GetText(), receiverResult, receiverResult)
		}
		return nil
	}

	if ctx.IDENTIFIER() == nil {
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitFieldAccessExpr: IDENTIFIER is nil for FieldAccessExprContext '%s'.\n", ctx.GetText())
		}
		return nil
	}
	fieldName := ctx.IDENTIFIER().GetText()
	fieldIdent := &ast.IdentifierExpr{Name: fieldName}

	node := &ast.FieldAccessExpr{
		Receiver: receiver,
		Field:    fieldIdent,
		Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
		Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
	}

	if v.DebugMode {
		// Use getNodeString for potentially complex receiver expressions
		fmt.Printf("AstBuilder.VisitFieldAccessExpr: Created FieldAccessExpr: %s.%s (AST Node: %s)\n", getNodeString(receiver), fieldIdent.Name, node.String())
	}
	return node
}

// VisitBlock is called when a block of statements is encountered.
func (v *AstBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting BlockContext")
	}
	var stmts []ast.Statement

	for _, stmtCtxInterface := range ctx.AllStatement() { // stmtCtxInterface is parser.IStatementContext
		var visitedNode interface{}

		stmtCtx, ok := stmtCtxInterface.(*parser.StatementContext)
		if !ok {

			visitedNode = v.Visit(stmtCtxInterface) // Fallback to generic visit
		} else {
			// Type switch based on the alternatives in the 'statement' rule
			if blockCtx := stmtCtx.Block(); blockCtx != nil {
				visitedNode = v.VisitBlock(blockCtx.(*parser.BlockContext))
			} else if varDeclCtx := stmtCtx.VarDecl(); varDeclCtx != nil {
				visitedNode = v.VisitVarDecl(varDeclCtx.(*parser.VarDeclContext))
			} else if assignmentCtx := stmtCtx.Assignment(); assignmentCtx != nil {
				visitedNode = v.VisitAssignment(assignmentCtx.(*parser.AssignmentContext))
			} else if ifStmtCtx := stmtCtx.IfStatement(); ifStmtCtx != nil {
				visitedNode = v.VisitIfStatement(ifStmtCtx.(*parser.IfStatementContext))
			} else if switchStmtCtx := stmtCtx.SwitchStatement(); switchStmtCtx != nil {
				visitedNode = v.VisitSwitchStatement(switchStmtCtx.(*parser.SwitchStatementContext))
			} else if forStmtCtx := stmtCtx.ForStatement(); forStmtCtx != nil {
				visitedNode = v.VisitForStatement(forStmtCtx.(*parser.ForStatementContext))
			} else if returnStmtCtx := stmtCtx.ReturnStatement(); returnStmtCtx != nil {
				visitedNode = v.VisitReturnStatement(returnStmtCtx.(*parser.ReturnStatementContext))
			} else if breakStmtCtx := stmtCtx.BreakStatement(); breakStmtCtx != nil {
				visitedNode = v.VisitBreakStatement(breakStmtCtx.(*parser.BreakStatementContext))
			} else if continueStmtCtx := stmtCtx.ContinueStatement(); continueStmtCtx != nil {
				visitedNode = v.VisitContinueStatement(continueStmtCtx.(*parser.ContinueStatementContext))
			} else if exprStmtCtx := stmtCtx.ExpressionStatement(); exprStmtCtx != nil {
				visitedNode = v.VisitExpressionStatement(exprStmtCtx.(*parser.ExpressionStatementContext))
			} else if incDecStmtCtx := stmtCtx.IncDecStatement(); incDecStmtCtx != nil {
				visitedNode = v.VisitIncDecStatement(incDecStmtCtx.(*parser.IncDecStatementContext))
			} else if stmtCtx.SEMI() != nil {
				// No AST node for a lone semicolon, so visitedNode remains nil
			} else {
				visitedNode = v.VisitChildren(stmtCtx) // Fallback
			}
		}

		if stmt, ok := visitedNode.(ast.Statement); ok && stmt != nil {
			stmts = append(stmts, stmt)
		} else if visitedNode != nil {
			// This branch handles cases where a visited node is not an ast.Statement but is also not nil.
			// This might occur if a visitor returns a raw value or a non-statement AST node unexpectedly.
			// It's important for debugging to understand what kind of node or value is being processed.
		} else { // This 'else' corresponds to the 'if stmt, ok := visitedNode.(ast.Statement)' and 'else if visitedNode != nil'
			// This branch handles cases where the visited node is nil.
			// This can happen for empty statements (like a lone semicolon),
			// or if a visitor for a particular statement type explicitly returns nil (e.g., unhandled type or error during visit).
		}
	} // Closes the for loop in VisitBlock ( iterating over blockCtx.AllStatement() )
	return &ast.BlockStmt{
		Statements: stmts,
		Line:       ctx.GetStart().GetLine(),
		Column:     ctx.GetStart().GetColumn(),
	}
} // Closes VisitBlock function

// VisitForStatement handles 'forStatement' rule, creating an ast.ForStmt node.
func (v *AstBuilder) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ForStatementContext")
	}
	forStmtNode := &ast.ForStmt{}

	// Process the loop body (BlockContext)
	if blockCtx := ctx.Block(); blockCtx != nil { // This is the main block for the for loop
		var bodyNode interface{}
		if concreteBlock, ok := blockCtx.(*parser.BlockContext); ok {
			bodyNode = concreteBlock.Accept(v) // Explicitly call Accept on concrete type
		} else {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement: Failed to assert blockCtx to *parser.BlockContext. Type was %T. Falling back to v.Visit().\n", blockCtx)
			bodyNode = v.Visit(blockCtx) // Fallback, though this is the problematic path
		}
		if b, ok := bodyNode.(*ast.BlockStmt); ok {
			forStmtNode.Body = b
		} else if bodyNode != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement: Expected *ast.BlockStmt for loop body, got %T for text '%s'\n", bodyNode, blockCtx.GetText())
			forStmtNode.Body = &ast.BlockStmt{} // Assign empty block on error
		} else {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement: Loop body visit returned nil. Text: '%s'\n", blockCtx.GetText())
			forStmtNode.Body = &ast.BlockStmt{} // Assign empty block
		}
	} else {
		fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement: Loop body (BlockContext) is nil. Text: '%s'\n", ctx.GetText())
		forStmtNode.Body = &ast.BlockStmt{} // Assign empty block if grammar allows, or this is an error state
	}

	// Check for C-style for loop: FOR forClause block
	if forClauseCtxRule := ctx.ForClause(); forClauseCtxRule != nil {
		if concreteForClauseCtx, ok := forClauseCtxRule.(*parser.ForClauseContext); ok {
			// Init statement
			if initStmtRule := concreteForClauseCtx.InitStmt(); initStmtRule != nil { // initStmtRule is parser.IInitStmtContext
				if initStmtCtx, ok := initStmtRule.(*parser.InitStmtContext); ok { // Type assertion
					initNodeRet := v.VisitInitStmt(initStmtCtx) // Now initStmtCtx is *parser.InitStmtContext
					if initS, okInit := initNodeRet.(ast.Statement); okInit {
						forStmtNode.Init = initS
					} else if initNodeRet != nil {
						// Error if initNodeRet is not nil AND not ast.Statement
						fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (ForClause): Expected ast.Statement for init, got %T for text '%s'\n", initNodeRet, initStmtCtx.GetText())
					}
					// If initNodeRet is nil, forStmtNode.Init remains nil, which is fine.
				} else {
					// Handle the case where the type assertion fails
					fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (ForClause): InitStmt is not *parser.InitStmtContext, got %T for text '%s'\n", initStmtRule, initStmtRule.GetText())
				}
			}
			// Condition expression
			if condExprCtx := concreteForClauseCtx.Expression(); condExprCtx != nil {
				condNodeRet := condExprCtx.Accept(v)
				if condE, okCond := condNodeRet.(ast.Expression); okCond {
					forStmtNode.Condition = condE
				} else if condNodeRet != nil {
					fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (ForClause): Expected ast.Expression for condition, got %T for text '%s'\n", condNodeRet, condExprCtx.GetText())
				} // else: condNodeRet is nil, so forStmtNode.Condition remains nil
			}
			// Post statement
			if postStmtRule := concreteForClauseCtx.PostStmt(); postStmtRule != nil { // postStmtRule is parser.IPostStmtContext
				if postStmtCtx, ok := postStmtRule.(*parser.PostStmtContext); ok { // Type assertion
					postNodeRet := v.VisitPostStmt(postStmtCtx) // Now postStmtCtx is *parser.PostStmtContext
					if postS, okPost := postNodeRet.(ast.Statement); okPost {
						forStmtNode.Post = postS
						if v.DebugMode {
							fmt.Printf("  VisitForStatement: Assigned Post statement. Type: %T, Value: %s\n", forStmtNode.Post, getNodeString(forStmtNode.Post))
						}
					} else if postNodeRet != nil && v.DebugMode {
						fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (ForClause): Expected ast.Statement for post, got %T for text '%s'\n", postNodeRet, postStmtCtx.GetText())
					}
					// If postNodeRet is nil, forStmtNode.Post remains nil, which is fine.
				} else {
					// Handle the case where the type assertion fails
					fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (ForClause): PostStmt is not *parser.PostStmtContext, got %T for text '%s'\n", postStmtRule, postStmtRule.GetText())
				}
			}
		} else {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement: ForClause is not *parser.ForClauseContext, got %T for text '%s'\n", forClauseCtxRule, forClauseCtxRule.GetText())
		}
	} else if condExprCtx := ctx.Expression(); condExprCtx != nil {
		// This is a condition-only loop (e.g., 'for i < 5 {}')
		// The ForClause was nil, so we check the direct Expression on ForStatementContext.
		condNodeRet := condExprCtx.Accept(v)
		if condE, okCond := condNodeRet.(ast.Expression); okCond {
			forStmtNode.Condition = condE
		} else if condNodeRet != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (Condition-only): Expected ast.Expression for condition, got %T for text '%s'\n", condNodeRet, condExprCtx.GetText())
		}
		// For this type of loop, Init and Post statements are nil, which is the correct default for ast.ForStmt.
	}
	// If both ForClause and direct Expression are nil, it's an infinite loop.
	// Init, Condition, and Post will remain nil in forStmtNode, which is correct.

	if rangeClauseCtxRule := ctx.RangeClause(); rangeClauseCtxRule != nil {
		forStmtNode.IsRangeLoop = true
		if concreteRangeClauseCtx, ok := rangeClauseCtxRule.(*parser.RangeClauseContext); ok {
			// Handle range variables (IDENTIFIER tokens)
			identifiers := concreteRangeClauseCtx.AllIDENTIFIER() // Returns []antlr.TerminalNode
			if len(identifiers) > 0 {
				keyName := identifiers[0].GetText()
				forStmtNode.RangeKey = &ast.IdentifierExpr{Name: keyName}

			}
			if len(identifiers) > 1 {
				valueName := identifiers[1].GetText()
				forStmtNode.RangeValue = &ast.IdentifierExpr{Name: valueName}

			}
			if len(identifiers) > 2 {
				// This would be unusual for a typical range clause, log a warning.
				fmt.Fprintf(os.Stderr, "[WARNING] AstBuilder.VisitForStatement (RangeClause): Found %d identifiers in range clause, expected 1 or 2. Text: '%s'\n", len(identifiers), concreteRangeClauseCtx.GetText())
			}
			// Range source expression
			if rangeSourceExprCtx := concreteRangeClauseCtx.Expression(); rangeSourceExprCtx != nil {
				rangeSourceNodeRet := rangeSourceExprCtx.Accept(v)
				if rangeSourceE, okRangeSource := rangeSourceNodeRet.(ast.Expression); okRangeSource {
					forStmtNode.RangeSource = rangeSourceE
				} else if rangeSourceNodeRet != nil {
					fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (RangeClause): Expected ast.Expression for range source, got %T for text '%s'\n", rangeSourceNodeRet, rangeSourceExprCtx.GetText())
				}
			} else {
				fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement (RangeClause): Range source expression is nil. Text: '%s'\n", concreteRangeClauseCtx.GetText())
			}
		} else {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitForStatement: RangeClause is not *parser.RangeClauseContext, got %T for text '%s'\n", rangeClauseCtxRule, rangeClauseCtxRule.GetText())
		}
	}
	// else: it's a simple FOR block (while-style loop if no ForClause or RangeClause), only body is processed, which was done above.

	return forStmtNode
} // Closes VisitForStatement function

// The VisitStatement function should correctly follow now.
func (v *AstBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting StatementContext")
	}
	if varDeclCtx := ctx.VarDecl(); varDeclCtx != nil {
		return varDeclCtx.Accept(v)
	} else if exprStmtCtx := ctx.ExpressionStatement(); exprStmtCtx != nil {
		return exprStmtCtx.Accept(v)
	} else if ctx.IfStatement() != nil {
		return ctx.IfStatement().Accept(v)
	} else if ctx.ForStatement() != nil {
		return ctx.ForStatement().Accept(v)
	} else if ctx.ReturnStatement() != nil {
		return ctx.ReturnStatement().Accept(v)
	} else if ctx.SwitchStatement() != nil {
		return ctx.SwitchStatement().Accept(v)
	} else if ctx.Assignment() != nil {
		return ctx.Assignment().Accept(v) // Assuming there's a VisitAssignment method
	} else if breakCtx := ctx.BreakStatement(); breakCtx != nil {
		return breakCtx.Accept(v)
	} else if continueCtx := ctx.ContinueStatement(); continueCtx != nil {
		return continueCtx.Accept(v)
	} else if incDecCtx := ctx.IncDecStatement(); incDecCtx != nil {
		return incDecCtx.Accept(v)
	}
	return nil
}

func (v *AstBuilder) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.Visiting IfStatement: %s\n", ctx.GetText())
	}

	// The grammar is: IF expression block (ELSE IF expression block)* (ELSE block)?
	// This creates flat lists of expressions and blocks in the context.

	// The first expression and block always belong to the initial 'if'.
	condition := ctx.Expression(0).Accept(v).(ast.Expression)
	consequence := ctx.Block(0).Accept(v).(*ast.BlockStmt)

	// Create the root of our recursive IfStmt structure.
	rootIfStmt := &ast.IfStmt{
		Condition:   condition,
		Consequence: consequence,
	}

	// This pointer will track the last 'if' or 'else if' statement,
	// so we can attach the next 'else if' or 'else' to its Alternative field.
	currentIfStmt := rootIfStmt

	// Iterate through the 'else if' clauses. There is one expression for each 'if' and 'else if'.
	numExpr := len(ctx.AllExpression())
	for i := 1; i < numExpr; i++ {
		// The i-th expression and i-th block correspond to the i-th 'else if'
		// (since i starts at 1).
		elseIfCondition := ctx.Expression(i).Accept(v).(ast.Expression)
		elseIfBlock := ctx.Block(i).Accept(v).(*ast.BlockStmt)

		// Create the new IfStmt for this 'else if' clause.
		newIfStmt := &ast.IfStmt{
			Condition:   elseIfCondition,
			Consequence: elseIfBlock,
		}

		// Link it to the previous statement's Alternative.
		currentIfStmt.Alternative = newIfStmt

		// Move the pointer forward to the newly created statement.
		currentIfStmt = newIfStmt
	}

	// Check for a final 'else' block. This exists if the number of blocks
	// is greater than the number of expressions.
	if len(ctx.AllBlock()) > numExpr {
		// The 'else' block is the last one in the list. Its index is `numExpr`.
		elseBlock := ctx.Block(numExpr).Accept(v).(*ast.BlockStmt)
		currentIfStmt.Alternative = elseBlock
	}

	return rootIfStmt
}

func (v *AstBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting AssignmentContext")
	}

	leftNode := ctx.Lvalue().Accept(v)
	leftExpr, ok := leftNode.(ast.Expression)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: VisitLvalue did not return ast.Expression. Got %T\n", leftNode)
		return &ast.AssignStmt{} // Return an empty or error node
	}

	operator := ctx.Assignment_op().GetText()

	rightNode := ctx.Expression().Accept(v)
	rightExpr, ok := rightNode.(ast.Expression)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: VisitExpression for RHS of assignment did not return ast.Expression. Got %T\n", rightNode)
		return &ast.AssignStmt{Left: leftExpr, Operator: operator} // Partial error node
	}

	return &ast.AssignStmt{
		Left:     leftExpr,
		Operator: operator,
		Right:    rightExpr,
		Line:     ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
		Column:   ctx.GetStart().GetColumn(), // <-- Y ESTO
	}
}

func (v *AstBuilder) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ExpressionStatementContext")
	}
	exprCtx := ctx.Expression()
	if exprCtx == nil {
		return nil
	}

	visitedExprNode := exprCtx.Accept(v) // This should dispatch to VisitExpression -> ... -> VisitLiteralExpr for literals

	if visitedExprNode == nil {
		return nil
	}

	astExpr, ok := visitedExprNode.(ast.Expression)
	if !ok {
		return nil
	}

	return &ast.ExpressionStmt{Expression: astExpr}
}

// VisitIncDecStatement creates an ast.IncDecStmt node.
// incDecStatement: lvalue (INC | DEC);
func (v *AstBuilder) VisitIncDecStatement(ctx *parser.IncDecStatementContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.VisitIncDecStatement: Visiting IncDecStatementContext. Text: '%s'\n", ctx.GetText())
	}

	var lval ast.Expression
	if lvalueCtx := ctx.Lvalue(); lvalueCtx != nil {
		// Ensure we are calling VisitLvalue with the concrete type if possible,
		// or rely on the generic Accept if LvalueContext is an interface itself.
		// Assuming Lvalue() returns a specific context type like *parser.LvalueContext.
		var lvalNodeRet interface{}
		if concreteLvalueCtx, ok := lvalueCtx.(*parser.LvalueContext); ok {
			lvalNodeRet = v.VisitLvalue(concreteLvalueCtx)
		} else {
			lvalNodeRet = v.Visit(lvalueCtx) // Fallback or if Lvalue() returns an interface type
		}

		if lv, ok := lvalNodeRet.(ast.Expression); ok {
			lval = lv
		} else {
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitIncDecStatement: Expected ast.Expression from VisitLvalue, got %T for '%s'\n", lvalNodeRet, lvalueCtx.GetText())
			}
			return nil // Or an error node
		}
	} else {
		if v.DebugMode {
			fmt.Fprintln(os.Stderr, "AstBuilder.VisitIncDecStatement: LvalueContext is nil")
		}
		return nil // Or an error node
	}

	op := ""
	if ctx.INC() != nil {
		op = "++"
	} else if ctx.DEC() != nil {
		op = "--"
	} else {
		if v.DebugMode {
			fmt.Fprintln(os.Stderr, "AstBuilder.VisitIncDecStatement: No INC or DEC token found")
		}
		return nil // Or an error node
	}

	incDecNode := &ast.IncDecStmt{
		LValue:   lval,
		Operator: op,
	}
	if v.DebugMode {
		fmt.Printf("  VisitIncDecStatement: Created ast.IncDecStmt: LValue=%s, Op=%s\n", getNodeString(lval), op)
	}
	return incDecNode
}

// Stubs for other statement types, inserted before VisitLvalue

// VisitBreakStatement creates an ast.BreakStmt node from a parser.BreakStatementContext.
func (v *AstBuilder) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting BreakStatementContext")
	}
	return &ast.BreakStmt{}
}

// VisitContinueStatement creates an ast.ContinueStmt node from a parser.ContinueStatementContext.
func (v *AstBuilder) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ContinueStatementContext")
	}
	return &ast.ContinueStmt{}
}

func (v *AstBuilder) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting SwitchStatementContext")
	}

	switchASTNode := &ast.SwitchStmt{
		Cases:   []*ast.CaseClause{},
		Default: nil,
	}

	// 1. Handle the optional switch expression
	if exprCtx := ctx.Expression(); exprCtx != nil {
		exprNodeRet := exprCtx.Accept(v) // Use Accept for robust dispatch
		if expr, ok := exprNodeRet.(ast.Expression); ok {
			switchASTNode.Expression = expr
		} else if exprNodeRet != nil {
			fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited switch expression is not ast.Expression. Got %T for '%s'\n", exprNodeRet, exprCtx.GetText())
		} else {
			fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited switch expression is nil for '%s'\n", exprCtx.GetText())
		}
	}

	// 2. Handle case clauses
	for _, caseClauseCtx := range ctx.AllCaseClause() { // caseClauseCtx is parser.ICaseClauseContext
		currentCase := &ast.CaseClause{
			Expressions: []ast.Expression{},
			Body:        []ast.Statement{},
		}

		// Expressions for the case
		if exprListCtx := caseClauseCtx.ExpressionList(); exprListCtx != nil {
			exprNodesObj := exprListCtx.Accept(v) // Use Accept for robust dispatch
			if exprs, ok := exprNodesObj.([]ast.Expression); ok {
				currentCase.Expressions = exprs
			} else if exprNodesObj != nil {
				fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited case expressions list is not []ast.Expression. Got %T for '%s'\n", exprNodesObj, exprListCtx.GetText())
			} else {
				fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited case expressions list is nil for '%s'\n", exprListCtx.GetText())
			}
		}

		// Body for the case
		for _, stmtCtx := range caseClauseCtx.AllStatement() { // stmtCtx is parser.IStatementContext
			stmtNodeObj := stmtCtx.Accept(v) // Use Accept for robust dispatch
			if stmt, ok := stmtNodeObj.(ast.Statement); ok {
				currentCase.Body = append(currentCase.Body, stmt)
			} else if stmtNodeObj != nil {
				fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited case statement is not ast.Statement. Got %T for '%s'\n", stmtNodeObj, stmtCtx.GetText())
			} else {
				fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited case statement is nil for '%s'\n", stmtCtx.GetText())
			}
		}
		switchASTNode.Cases = append(switchASTNode.Cases, currentCase)
	}

	// 3. Handle default clause
	if defaultClauseCtx := ctx.DefaultClause(); defaultClauseCtx != nil { // defaultClauseCtx is parser.IDefaultClauseContext
		currentDefault := &ast.DefaultClause{
			Body: []ast.Statement{},
		}
		for _, stmtCtx := range defaultClauseCtx.AllStatement() { // stmtCtx is parser.IStatementContext
			stmtNodeObj := stmtCtx.Accept(v) // Use Accept for robust dispatch
			if stmt, ok := stmtNodeObj.(ast.Statement); ok {
				currentDefault.Body = append(currentDefault.Body, stmt)
			} else if stmtNodeObj != nil {
				fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited default statement is not ast.Statement. Got %T for '%s'\n", stmtNodeObj, stmtCtx.GetText())
			} else {
				fmt.Fprintf(os.Stderr, "[AstBuilder.VisitSwitchStatement] Warning: Visited default statement is nil for '%s'\n", stmtCtx.GetText())
			}
		}
		switchASTNode.Default = currentDefault
	}

	return switchASTNode
}

func (v *AstBuilder) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ReturnStatementContext")
	}
	var returnValue ast.Expression
	if exprCtx := ctx.Expression(); exprCtx != nil {
		// Use Accept for expression alternatives, as per MEMORY[6787a305-9aff-4a30-8be8-27787ff29659]
		exprNodeRet := exprCtx.Accept(v)
		if v.DebugMode {
			fmt.Fprintf(os.Stderr, "[WARNING] VisitReturnStatement: Expected ast.Expression from visiting ExpressionContext, got %T\n", exprNodeRet)
		}
		if exprNode, ok := exprNodeRet.(ast.Expression); ok {
			returnValue = exprNode
		} else if exprNodeRet != nil {
			// TODO: Log if exprNodeRet is not nil but not ast.Expression
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "[WARNING] VisitReturnStatement: Expected ast.Expression from visiting ExpressionContext, got %T\n", exprNodeRet)
			}
		} else {
			// TODO: Log if exprNodeRet is nil (meaning Accept(v) returned nil)
			if v.DebugMode {
				fmt.Fprintf(os.Stderr, "[WARNING] VisitReturnStatement: Visiting ExpressionContext resulted in nil\n")
			}
		}
	}
	// If returnValue is nil here, it means either there was no expression in the return statement,
	// or visiting it resulted in nil (which might be an error or intended for some expression types).
	return &ast.ReturnStmt{Value: returnValue}
}

func (v *AstBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting LiteralContext")
	}
	if ctx.STRING_LITERAL() != nil {
		rawStr := ctx.STRING_LITERAL().GetText()
		val, err := strconv.Unquote(rawStr) // Use strconv.Unquote for proper handling of escapes
		if err != nil {
			fmt.Printf("AstBuilder.Visiting Literal: Warning: could not unquote string literal: %s, err: %v. Using raw value.\n", rawStr, err)
			// Fallback for basic unquoting if strconv.Unquote fails (e.g. already unquoted or malformed)
			if len(rawStr) >= 2 && rawStr[0] == '"' && rawStr[len(rawStr)-1] == '"' {
				val = rawStr[1 : len(rawStr)-1]
			} else {
				val = rawStr // Use as is if not clearly quoted
			}
		}
		return &ast.StringLiteral{
			Value:  val,
			Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}
	if ctx.INT_LITERAL() != nil {
		// ast.IntegerLiteral has field `Value string`
		return &ast.IntegerLiteral{
			Value:  ctx.INT_LITERAL().GetText(),
			Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}
	if ctx.FLOAT_LITERAL() != nil {
		// ast.FloatLiteral has field `Value string`
		return &ast.FloatLiteral{
			Value:  ctx.FLOAT_LITERAL().GetText(),
			Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}
	if boolLitCtx := ctx.BoolLiteral(); boolLitCtx != nil { // Check for the boolLiteral sub-rule
		if boolLitCtx.TRUE() != nil {
			return &ast.BoolLiteral{
				Value:  true,
				Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
				Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
			}
		} else if boolLitCtx.FALSE() != nil {
			return &ast.BoolLiteral{
				Value:  false,
				Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
				Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
			}
		}
		fmt.Printf("AstBuilder.Visiting Literal: Warning: boolLiteral context found but neither TRUE nor FALSE token: %s\n", boolLitCtx.GetText())
		fmt.Printf("AstBuilder.Visiting Literal: returning nil due to invalid bool literal\n")
		return nil
	}
	if ctx.RUNE_LITERAL() != nil {
		// TODO: Implement RuneLiteral handling if ast.RuneLiteral exists
		// For now, similar to string, but need to unquote rune quotes (e.g. 'a')
		rawRune := ctx.RUNE_LITERAL().GetText()
		val := rawRune
		if len(rawRune) >= 2 && rawRune[0] == '\'' && rawRune[len(rawRune)-1] == '\'' {
			val = rawRune[1 : len(rawRune)-1]
			// Further unescaping might be needed for runes if supported (e.g. '\n', '\uFFFF')
		}
		// Assuming ast.CharLiteral or similar exists and takes a string for the char value
		return &ast.CharLiteral{
			Value:  val,
			Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
		} // Or ast.RuneLiteral
		// fmt.Printf("Warning: RUNE_LITERAL found but not fully handled: %s\n", rawRune)
		// fmt.Printf("AstBuilder.Visiting Literal: returning nil due to unhandled rune literal\n")
		// return nil // Placeholder
	}
	if ctx.NIL() != nil {
		return &ast.NilLiteral{}
	}

	fmt.Printf("Warning: Unknown or unhandled literal type in VisitLiteral: %s\n", ctx.GetText())
	return nil
}

// Note: Removed VisitStringLiteral as specific literal types are handled in VisitLiteral directly.
func (v *AstBuilder) VisitLvalue(ctx *parser.LvalueContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting LvalueContext")
	}

	if ctx.PrimaryExpr() == nil && ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		return &ast.IdentifierExpr{Name: name}
	}

	if pExprCtx := ctx.PrimaryExpr(); pExprCtx != nil {
		baseValue := pExprCtx.Accept(v)
		if baseValue == nil {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitLvalue: Visiting PrimaryExpr for lvalue returned nil for: %s\n", pExprCtx.GetText())
			return nil
		}
		baseExpr, ok := baseValue.(ast.Expression)
		if !ok {
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitLvalue: Visiting PrimaryExpr for lvalue did not return ast.Expression for: %s, got %T\n", pExprCtx.GetText(), baseValue)
			return nil
		}

		if ctx.DOT() != nil && ctx.IDENTIFIER() != nil {
			fieldName := ctx.IDENTIFIER().GetText()
			return &ast.FieldAccessExpr{Receiver: baseExpr, Field: &ast.IdentifierExpr{Name: fieldName}}

		} else if ctx.L_SQUARE() != nil && ctx.Expression() != nil {
			indexValue := ctx.Expression().Accept(v)
			if indexValue == nil {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitLvalue: Visiting index Expression for lvalue returned nil for: %s\n", ctx.Expression().GetText())
				return nil
			}
			indexExpr, ok := indexValue.(ast.Expression)
			if !ok {
				fmt.Fprintf(os.Stderr, "AstBuilder.VisitLvalue: Visiting index Expression for lvalue did not return ast.Expression for: %s, got %T\n", ctx.Expression().GetText(), indexValue)
				return nil
			}

			return &ast.IndexAccessExpr{Receiver: baseExpr, Index: indexExpr}
		} else {
			// The grammar `lvalue: IDENTIFIER | primaryExpr (DOT IDENTIFIER | L_SQUARE expression R_SQUARE)` means primaryExpr alone is not an lvalue unless it's an IDENTIFIER handled by the first case.
			// So, if pExprCtx is not nil, one of the accessor patterns (DOT or L_SQUARE) must also be present.
			fmt.Fprintf(os.Stderr, "AstBuilder.VisitLvalue: Lvalue has PrimaryExpr but no valid accessor (DOT/IDENTIFIER or L_SQUARE/Expression). Text: '%s'\n", ctx.GetText())
			return nil
		}
	}

	fmt.Fprintf(os.Stderr, "AstBuilder.VisitLvalue: Unhandled lvalue structure. Text: '%s'\n", ctx.GetText())
	return nil
}

func (v *AstBuilder) VisitAssignment_op(ctx *parser.Assignment_opContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting Assignment_opContext")
	} // Represents assignment operators like =, +=, etc.
	return nil
}

func (v *AstBuilder) VisitCaseClause(ctx *parser.CaseClauseContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting CaseClauseContext") // Corresponds to ast.CaseClause
	}
	// TODO: Implement actual AST node creation
	return nil
}

func (v *AstBuilder) VisitDefaultClause(ctx *parser.DefaultClauseContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting DefaultClauseContext") // Corresponds to ast.DefaultClause
	}
	// TODO: Implement actual AST node creation
	return nil
}

func (v *AstBuilder) VisitForClause(ctx *parser.ForClauseContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting ForClauseContext") // Part of ForStmt for C-style loops
	}
	// TODO: Implement actual AST node creation
	return nil
}

func (v *AstBuilder) VisitInitStmt(ctx *parser.InitStmtContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting InitStmtContext")
	}
	// Assuming InitStmtContext has a 'statement()' child that returns parser.IStatementContext
	// This needs to match your VLangCherryParser.g4 rule for initStmt.
	// e.g., initStmt: statement;
	// For a rule like 'initStmt: simpleStmt;', the SimpleStmtContext is the first child.
	// Access it generically via GetChild(0) and call Accept on it.
	if ctx.GetChildCount() > 0 {
		childParseTree := ctx.GetChild(0) // Returns antlr.ParseTree
		if childParseTree != nil {
			// The childParseTree should be an antlr.RuleContext (specifically ISimpleStmtContext or one of its implementations).
			// We need to type-assert it to antlr.RuleContext to call Accept.
			if ruleCtx, ok := childParseTree.(antlr.RuleContext); ok {
				stmtAstNode := ruleCtx.Accept(v)
				return stmtAstNode
			} else {
				fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitInitStmt: Child node is not an antlr.RuleContext. Type: %T, Text: '%s'\n", childParseTree, ctx.GetText())
			}
		} else {
			fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitInitStmt: Child node is nil in InitStmtContext. Text: '%s'\n", ctx.GetText())
		}
	} else {
		fmt.Fprintf(os.Stderr, "[ERROR] AstBuilder.VisitInitStmt: No child node found in InitStmtContext. Text: '%s'\n", ctx.GetText())
	}
	return nil
}

// VisitPostStmt handles the 'postStmt' rule.
// postStmt: assignment | expressionStatement | incDecStatement;
func (v *AstBuilder) VisitPostStmt(ctx *parser.PostStmtContext) interface{} {
	if v.DebugMode {
		fmt.Printf("AstBuilder.VisitPostStmt: Visiting PostStmtContext. Text: '%s', ChildCount: %d\n", ctx.GetText(), ctx.GetChildCount())
	}

	// The PostStmtContext will have one child which is either an AssignmentContext,
	// an ExpressionStatementContext, or an IncDecStatementContext.
	// We need to visit that child to get the actual AST node for the statement.
	if ctx.GetChildCount() > 0 {
		childParseTree := ctx.GetChild(0) // This should be the actual statement context
		if ruleCtx, ok := childParseTree.(antlr.RuleContext); ok {
			if v.DebugMode {
				fmt.Printf("  VisitPostStmt: Child is a RuleContext. Type: %T, Text: '%s'. Visiting it now.\n", ruleCtx, ruleCtx.GetText())
			}
			stmtAstNode := ruleCtx.Accept(v) // This should dispatch to VisitAssignment, VisitExpressionStatement, or VisitIncDecStatement
			if stmtAstNode == nil {
				if v.DebugMode {
					fmt.Printf("  VisitPostStmt: Visiting child '%s' (type %T) returned nil.\n", ruleCtx.GetText(), ruleCtx)
				}
				return nil
			}
			if v.DebugMode {
				fmt.Printf("  VisitPostStmt: Visiting child '%s' (type %T) returned AST node: %T (%s)\n", ruleCtx.GetText(), ruleCtx, stmtAstNode, getNodeString(stmtAstNode))
			}
			return stmtAstNode
		} else if v.DebugMode {
			childText := "<unknown_child_text>"
			if tn, okTn := childParseTree.(antlr.TerminalNode); okTn {
				childText = tn.GetText()
			}
			fmt.Printf("  VisitPostStmt: Child is not a RuleContext. Type: %T, Text: '%s'\n", childParseTree, childText)
		}
	} else if v.DebugMode {
		fmt.Println("  VisitPostStmt: No children found in PostStmtContext.")
	}
	return nil
}

func (v *AstBuilder) VisitRangeClause(ctx *parser.RangeClauseContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting RangeClauseContext") // Part of ForStmt for range loops
	}
	// TODO: Implement actual AST node creation
	return nil
}

func (v *AstBuilder) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	if v.DebugMode {
		fmt.Println("AstBuilder.Visiting BoolLiteralContext")
	} // Corresponds to ast.BoolLiteral
	// This method is called when ANTLR encounters a 'boolLiteral' rule directly.
	// It's also used by VisitLiteral if that method dispatches to specific literal types.
	if ctx.TRUE() != nil {
		return &ast.BoolLiteral{
			Value:  true,
			Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	} else if ctx.FALSE() != nil {
		return &ast.BoolLiteral{
			Value:  false,
			Line:   ctx.GetStart().GetLine(),   // <-- AGREGA ESTO
			Column: ctx.GetStart().GetColumn(), // <-- Y ESTO
		}
	}
	// Should not be reached if grammar ensures TRUE or FALSE for BoolLiteralContext
	fmt.Fprintf(os.Stderr, "AstBuilder.VisitBoolLiteral: BoolLiteralContext does not contain TRUE or FALSE. Text: %s\n", ctx.GetText())
	return nil
}
