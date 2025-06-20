package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// ErrorType defines the type of error (Lexical, Syntax, Semantic)
// We'll use iota to create a set of related constants. The first value will be 0, the next 1, and so on.
// We can also assign explicit values if needed, but iota is convenient for simple enumerations.
type ErrorType int

const (
	LexicalError ErrorType = iota
	SyntaxError
	SemanticError
)

// String method for ErrorType for easy printing
func (et ErrorType) String() string {
	switch et {
	case LexicalError:
		return "Lexical Error"
	case SyntaxError:
		return "Syntax Error"
	case SemanticError:
		return "Semantic Error"
	default:
		return "Unknown Error"
	}
}

// Error struct to hold error details
type Error struct {
	Line    int
	Column  int
	Message string
	Type    ErrorType
}

// String method for Error for easy printing
func (e Error) String() string {
	return fmt.Sprintf("%s: line %d:%d: %s", e.Type.String(), e.Line, e.Column, e.Message)
}

var (
	errorsMu     sync.Mutex
	globalErrors []Error
)

// AddError adds an error to the global list of errors.
// It is thread-safe.
func AddError(line, column int, message string, errorType ErrorType) {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	globalErrors = append(globalErrors, Error{Line: line, Column: column, Message: message, Type: errorType})
}

// GetErrors returns a copy of the global list of errors.
// It is thread-safe.
func GetErrors() []Error {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	// Return a copy to prevent external modification of the internal slice
	errs := make([]Error, len(globalErrors))
	copy(errs, globalErrors)
	return errs
}

// ClearErrors clears all errors from the global list.
// It is thread-safe.
func ClearErrors() {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	globalErrors = []Error{}
}

// HasErrors checks if there are any errors in the global list.
// It is thread-safe.
func HasErrors() bool {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	return len(globalErrors) > 0
}

// PrintErrorsTable prints all collected errors in a formatted table to the console.
func PrintErrorsTable() {
	errs := GetErrors()
	if len(errs) == 0 {
		fmt.Println("No errors reported.")
		return
	}

	fmt.Printf("\n--- Error Report (%d errors) ---\n", len(errs))
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-4s | %-6s | %s\n", "Type", "Line", "Column", "Message")
	fmt.Println("--------------------------------------------------------------------------------")
	for _, err := range errs {
		fmt.Printf("%-15s | %-4d | %-6d | %s\n", err.Type.String(), err.Line, err.Column, err.Message)
	}
	fmt.Println("--------------------------------------------------------------------------------")
}

// ErrorsTableString devuelve la tabla de errores en formato tabular (con tabuladores)
func ErrorsTableString() string {
	errs := GetErrors()
	if len(errs) == 0 {
		return "No errors reported.\n"
	}
	var sb strings.Builder
	sb.WriteString("Tipo\tLínea\tColumna\tMensaje\n")
	for _, err := range errs {
		sb.WriteString(fmt.Sprintf("%s\t%d\t%d\t%s\n", err.Type.String(), err.Line, err.Column, err.Message))
	}
	return sb.String()
}

// ErrorListener collects syntax errors during parsing and uses the global error system
type ErrorListener struct {
	*antlr.DefaultErrorListener
	// No longer stores errors locally, uses globalErrors
}

// NewErrorListener creates a new ErrorListener.
func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

// SyntaxError is called by ANTLR when a syntax error is encountered.
// It adds the error to our global error list.
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	// Add the syntax error to the global list
	AddError(line, column, msg, SyntaxError)
}

// Note: The ErrorListener's own HasErrors() method is no longer needed
// as we will use the global HasErrors() function.
