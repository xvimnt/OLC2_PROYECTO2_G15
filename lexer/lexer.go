package lexer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/xvimnt/OLC2_PROYECTO2_G15/parser"
)

// Lexer wraps ANTLR-generated lexer to provide additional functionality
type Lexer struct {
	lexer         *parser.VLangCherryLexer
	errorListener *ErrorListener
}

// NewLexer creates a new lexer for the given input
func NewLexer(input string) *Lexer {
	// Create an input stream from the input string
	inputStream := antlr.NewInputStream(input)

	// Create the ANTLR lexer
	lexer := parser.NewVLangCherryLexer(inputStream)

	// Create and set a custom error listener
	errorListener := NewErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	return &Lexer{
		lexer:         lexer,
		errorListener: errorListener,
	}
}

// GetErrors returns any lexical errors that occurred during tokenization
func (l *Lexer) GetErrors() []string {
	return l.errorListener.Errors
}

// HasErrors checks if any lexical errors occurred
func (l *Lexer) HasErrors() bool {
	return len(l.errorListener.Errors) > 0
}

// NextToken forwards to the embedded lexer's NextToken method
func (l *Lexer) NextToken() antlr.Token {
	return l.lexer.NextToken()
}

// Reset resets the lexer to the beginning of the input
func (l *Lexer) Reset() {
	l.lexer.SetInputStream(l.lexer.GetInputStream())
}

// TokenizeAll processes all tokens and returns them as a slice
func (l *Lexer) TokenizeAll() []Token {
	var tokens []Token

	for {
		token := l.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		tokens = append(tokens, Token{
			Type:    token.GetTokenType(),
			Text:    token.GetText(),
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Channel: token.GetChannel(),
		})
	}

	l.Reset()
	return tokens
}

// Token represents a lexical token with position information
type Token struct {
	Type    int
	Text    string
	Line    int
	Column  int
	Channel int
}

// String provides a string representation of the token
func (t Token) String() string {
	return fmt.Sprintf("Token{type=%d, text='%s', line=%d, column=%d}",
		t.Type, t.Text, t.Line, t.Column)
}

// ErrorListener collects lexical errors during tokenization
type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

// NewErrorListener creates a new error listener
func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]string, 0),
	}
}

// SyntaxError is called when a syntax error is encountered
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

// LexicalError represents an error that occurred during lexical analysis
type LexicalError struct {
	Line    int
	Column  int
	Message string
}

// String provides a string representation of the error
func (e LexicalError) String() string {
	return fmt.Sprintf("Lexical error at %d:%d: %s", e.Line, e.Column, e.Message)
}

func (e LexicalError) Error() string {
	return e.String()
}
