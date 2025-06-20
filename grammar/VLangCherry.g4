/**
 * ANTLR4 Grammar for the VLangCherry Language.
 * This grammar covers the full specification provided in the document,
 * including declarations, statements, expressions, and types.
 */
grammar VLangCherry;

// Parser Rules
// -----------------------------------------------------------------------------

// The root of the parse tree. A program is a sequence of top-level declarations.
program
    : (topLevelDeclaration)* EOF
    ;

topLevelDeclaration
    : functionDeclaration
    | structDeclaration
    | varDecl SEMI?
    ;

// -- Declarations --

structDeclaration
    : STRUCT IDENTIFIER L_BRACE (fieldDeclaration SEMI?)+ R_BRACE
    ;

fieldDeclaration
    : type IDENTIFIER
    ;

functionDeclaration
    : FN receiver? IDENTIFIER parameters returnType? block
    ;

// Method receiver, e.g., `(p Persona)` or `(p *Persona)`
receiver
    : L_PAREN IDENTIFIER type R_PAREN
    ;

parameters
    : L_PAREN (parameterDecl (COMMA parameterDecl)*)? R_PAREN
    ;

parameterDecl
    : IDENTIFIER type
    ;

returnType
    : type
    ;

// Variable declarations covering all forms like:
// mut x int; mut x int = 1; x := 1; mut x := 1;
varDecl
    : MUT IDENTIFIER type (EQ expression)?
    | MUT? IDENTIFIER COLON_EQ expression
    ;

// -- Statements --

statement
    : block
    | varDecl SEMI?
    | assignment SEMI?
    | incDecStatement SEMI?
    | ifStatement
    | switchStatement
    | forStatement
    | returnStatement SEMI?
    | breakStatement SEMI?
    | continueStatement SEMI?
    | expressionStatement SEMI?
    | SEMI // Empty statement
    ;

assignment
    : lvalue assignment_op expression
    ;

// An expression that can appear on its own as a statement (e.g., function call).
expressionStatement
    : expression
    ;

// A left-value: something that can be assigned to (identifier, field access, index access).
lvalue
    : IDENTIFIER
    | primaryExpr (DOT IDENTIFIER | L_SQUARE expression R_SQUARE)
    ;

assignment_op
    : EQ | ADD_EQ | SUB_EQ
    ;

incDecStatement
    : lvalue (INC | DEC)
    ;

block
    : L_BRACE (statement)* R_BRACE
    ;

// -- Control Flow --

ifStatement
    : IF expression block
      (ELSE IF expression block)*
      (ELSE block)?
    ;

switchStatement
    : SWITCH expression? L_BRACE (caseClause)* (defaultClause)? R_BRACE
    ;

caseClause
    : CASE expressionList COLON (statement)*
    ;

defaultClause
    : DEFAULT COLON (statement)*
    ;

forStatement
    : FOR (forClause | expression | rangeClause)? block
    ;

// C-style for: `for i := 0; i < 10; i++`
forClause
    : initStmt? SEMI expression? SEMI postStmt?
    ;

initStmt
    : varDecl | assignment | expressionStatement
    ;

postStmt
    : assignment | expressionStatement | incDecStatement
    ;

// Range-based for: `for index, value in slice`
rangeClause
    : IDENTIFIER (COMMA IDENTIFIER)? IN expression
    ;

returnStatement : RETURN expression?;
breakStatement  : BREAK;
continueStatement: CONTINUE;

// -- Expressions (ordered by precedence) --

expression
    : logicalOrExpr
    ;

logicalOrExpr
    : logicalAndExpr (OR logicalAndExpr)*
    ;

logicalAndExpr
    : equalityExpr (AND equalityExpr)*
    ;

equalityExpr
    : relationalExpr ((EQ_EQ | NOT_EQ) relationalExpr)*
    ;

relationalExpr
    : additiveExpr ((LT | LTE | GT | GTE) additiveExpr)*
    ;

additiveExpr
    : multiplicativeExpr ((ADD | SUB) multiplicativeExpr)*
    ;

multiplicativeExpr
    : unaryExpr ((MUL | DIV | MOD) unaryExpr)*
    ;

unaryExpr
    : (NOT | SUB | MUL | AMPERSAND) unaryExpr
    | primaryExpr
    ;

primaryExpr
    : L_PAREN expression R_PAREN                   #ParenExpr
    | literal                                      #LiteralExpr
    | typeConversion                               #TypeConversionExpr
    | compositeLit                                 #TypedCompositeLitExpr  // Renamed label
    | L_BRACE elementList? R_BRACE                 #InferredCompositeLitExpr // New alternative
    | IDENTIFIER                                   #IdentifierExpr
    | TYPEOF L_PAREN expression R_PAREN          #TypeOfExpr
    | L_SQUARE expressionList? R_SQUARE          #ArrayLiteralExpr
    // Left-recursive rules for accessors and calls
    | primaryExpr L_SQUARE expression R_SQUARE     #IndexAccessExpr
    | primaryExpr DOT IDENTIFIER                   #FieldAccessExpr
    | primaryExpr L_PAREN expressionList? R_PAREN  #CallExpr
    ;

expressionList
    : expression (COMMA expression)*
    ;

// -- Literals and Types --

literal
    : INT_LITERAL
    | FLOAT_LITERAL
    | STRING_LITERAL
    | RUNE_LITERAL
    | boolLiteral
    | NIL
    ;

boolLiteral
    : TRUE | FALSE
    ;

// Covers both slice and struct literals, e.g., `[]int{1,2}` or `Person{Name:"v"}`
compositeLit
    : type L_BRACE elementList? R_BRACE
    ;

elementList
    : (keyedElement (COMMA keyedElement)* | expressionList) (COMMA)?
    ;

keyedElement
    : IDENTIFIER COLON expression
    ;

typeConversion
    : primitiveType L_PAREN expression R_PAREN
    ;

type
    : primitiveType                                #TypePrimitive
    | IDENTIFIER                                   #IdentifierType
    | MUL type                                     #PointerType
    | L_SQUARE R_SQUARE type                       #SliceType
    ;

primitiveType
    : INT | FLOAT64 | STRING | BOOL | RUNE
    ;


// Lexer Rules
// -----------------------------------------------------------------------------
// Keywords
BREAK       : 'break';
CASE        : 'case';
CONTINUE    : 'continue';
DEFAULT     : 'default';
ELSE        : 'else';
FN          : 'fn';
FOR         : 'for';
IF          : 'if';
IN          : 'in';
MUT         : 'mut';
RETURN      : 'return';
STRUCT      : 'struct';
SWITCH      : 'switch';
NIL         : 'nil';
TRUE        : 'true';
FALSE       : 'false';
TYPEOF      : 'TypeOf';

// Primitive Types as Keywords
INT         : 'int';
FLOAT64     : 'float64';
STRING      : 'string';
BOOL        : 'bool';
RUNE        : 'rune';

// Operators and Punctuation
L_PAREN     : '(';
R_PAREN     : ')';
L_BRACE     : '{';
R_BRACE     : '}';
L_SQUARE    : '[';
R_SQUARE    : ']';

COMMA       : ',';
DOT         : '.';
SEMI        : ';';
COLON       : ':';

ADD         : '+';
SUB         : '-';
MUL         : '*';
DIV         : '/';
MOD         : '%';

AMPERSAND   : '&';

AND         : '&&';
OR          : '||';
NOT         : '!';

EQ_EQ       : '==';
NOT_EQ      : '!=';
LT          : '<';
LTE         : '<=';
GT          : '>';
GTE         : '>=';

EQ          : '=';
ADD_EQ      : '+=';
SUB_EQ      : '-=';
COLON_EQ    : ':=';
INC         : '++';
DEC         : '--';

// Literals
INT_LITERAL
    : [0-9]+
    ;

FLOAT_LITERAL
    : [0-9]+ ('.' [0-9]*) ([eE] [+-]? [0-9]+)?
    | '.' [0-9]+ ([eE] [+-]? [0-9]+)?
    ;

STRING_LITERAL
    : '"' ( ~["\\] | ESCAPE_SEQUENCE )* '"'
    ;

RUNE_LITERAL
    : '\'' ( ~['\\] | ESCAPE_SEQUENCE ) '\''
    ;

fragment ESCAPE_SEQUENCE
    : '\\' (["\\nrt] | 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT)
    ;

fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;

// Identifier
IDENTIFIER
    : [a-zA-Z_] [a-zA-Z0-9_]*
    ;

// Ignored tokens
WS              : [ \t\r\n]+ -> skip;
LINE_COMMENT    : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT   : '/*' .*? '*/' -> skip;