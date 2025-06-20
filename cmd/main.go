package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/xvimnt/OLC2_PROYECTO1_G15/ast"
	"github.com/xvimnt/OLC2_PROYECTO1_G15/interpreter"
	"github.com/xvimnt/OLC2_PROYECTO1_G15/parser"
)

const version = "0.1.0"

var debugMode bool

func init() {
	debugStr := strings.ToLower(os.Getenv("VLANG_DEBUG"))
	debugMode = debugStr == "true" || debugStr == "1"
}

func main() {

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Error: No file specified")
			printHelp()
			os.Exit(1)
		}
		runFile(os.Args[2], debugMode)
	case "repl":
		runRepl(debugMode)
	case "parse":
		if len(os.Args) < 3 {
			fmt.Println("Error: No file specified")
			printHelp()
			os.Exit(1)
		}
		parseFile(os.Args[2], true, debugMode)
	case "help":
		printHelp()
	case "version":
		fmt.Printf("V Language Interpreter v%s\n", version)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("V Language Interpreter - Usage:")
	fmt.Println("  OLC2_PROYECTO1_G15 run <file>       Run a V language file")
	fmt.Println("  OLC2_PROYECTO1_G15 repl             Start an interactive REPL")
	fmt.Println("  OLC2_PROYECTO1_G15 parse <file>     Parse a file and print the AST")
	fmt.Println("  OLC2_PROYECTO1_G15 help             Display this help message")
	fmt.Println("  OLC2_PROYECTO1_G15 version          Display version information")
}

func runFile(filePath string, debugMode bool) {
	ClearErrors() // Clear any previous errors before running a new file
	// Check file extension
	if !strings.HasSuffix(filePath, ".v") && !strings.HasSuffix(filePath, ".mylang") {
		fmt.Printf("Warning: File %s does not have a .v or .mylang extension\n", filePath)
	}

	// Read file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: Could not read file %s: %v\n", filePath, err)
		os.Exit(1)
	}

	// Execute the code
	result, interp, err := executeCode(string(content), filepath.Base(filePath), false, debugMode)

	if err != nil {
		// Check if the error indicates syntax errors already logged by the listener.
		// These messages are "syntax errors occurred during parsing"
		// or "syntax errors occurred during execution setup".
		isAlreadyLoggedSyntaxError := strings.HasPrefix(err.Error(), "syntax errors occurred")

		if !isAlreadyLoggedSyntaxError { // If it's not one of those, it's likely a runtime or other error
			if rErr, ok := err.(interpreter.RuntimeError); ok {
				AddError(rErr.Line, rErr.Column, rErr.Message, SemanticError)
			} else {
				// Add other unclassified errors as SemanticError.
				// Use 0,0 for line/column if not available from the error.
				AddError(0, 0, err.Error(), SemanticError)
			}
		}
	}

	if HasErrors() {
		if os.Getenv("VLANG_ERRORS") == "1" {
			fmt.Println("=== Tabla de Errores ===")
			fmt.Print(ErrorsTableString())
		} else {
			PrintErrorsTable()
		}
		os.Exit(1)
	}

	// If no global errors, proceed to print result and captured output
	if result != nil {
		valueType := result.Type()
		if valueType != interpreter.NoneType {
			fmt.Printf("Result: %s\n", result.String())
		}
	}

	// Print the captured output from the singleton
	printer := interpreter.GetNativePrintSingleton()
	output := printer.GetAndClear()
	if output != "" {
		fmt.Print(output)
	}

	// Imprimir la tabla de símbolos si se solicita
	if os.Getenv("VLANG_SYMBOLS") == "1" {
		fmt.Println("=== Tabla de Símbolos ===")
		fmt.Print(interp.GetGlobals().DumpSymbols())
	}
	// Imprimir la tabla de errores si se solicita
	if os.Getenv("VLANG_ERRORS") == "1" && HasErrors() {
		fmt.Println("=== Tabla de Errores ===")
		fmt.Print(ErrorsTableString())
	}
}

func runRepl(debugMode bool) {
	fmt.Printf("V Language Interpreter v%s - REPL Mode\n", version)
	fmt.Println("Type 'exit' or 'quit' to exit, 'help' for help")

	scanner := bufio.NewScanner(os.Stdin)
	astBuilder := NewAstBuilder(debugMode)                                  // Create AstBuilder
	exprParserFn := func(expressionString string) (ast.Expression, error) { // Define the parser function
		return astBuilder.ParseExpressionFromString(expressionString)
	}
	interp := interpreter.NewInterpreter(debugMode, exprParserFn) // Pass it to NewInterpreter
	lineNumber := 1

	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "exit" || input == "quit" {
			break
		} else if input == "help" {
			fmt.Println("Available commands:")
			fmt.Println("  exit, quit - Exit the REPL")
			fmt.Println("  help       - Display this help message")
			continue
		} else if input == "" {
			continue
		}

		// Execute the code
		result, err := executeCodeWithInterpreter(interp, input, fmt.Sprintf("<repl:%d>", lineNumber), true, debugMode)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else if result != nil {
			valueType := result.Type()
			if valueType != interpreter.NoneType {
				fmt.Printf("%s\n", result.String())
			}
		}

		lineNumber++
	}

	if scanner.Err() != nil {
		fmt.Printf("Error reading input: %v\n", scanner.Err())
	}
}

func parseFile(filePath string, printAst bool, debugMode bool) (*ast.Program, error) {
	// Read file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: Could not read file %s: %v\n", filePath, err)
		os.Exit(1) // Consider returning error instead of os.Exit(1) for better testability/library use
	}

	// Create lexer and parser
	inputStream := antlr.NewInputStream(string(content))
	lexer := parser.NewVLangCherryLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewVLangCherryParser(tokenStream)

	// Add error listener
	errorListener := NewErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the input
	programContext := p.Program() // This is parser.IProgramContext

	// Check for syntax errors
	if HasErrors() { // Use global HasErrors
		// Errors are already logged globally by the ErrorListener.
		// PrintErrorsTable() will be called by the top-level command handler (runFile) if needed.
		return nil, fmt.Errorf("syntax errors occurred during parsing")
	}

	// Build the AST
	var visitor parser.VLangCherryVisitor = NewAstBuilder(debugMode) // Explicitly use interface type

	// Visit the program context. The result is interface{}.
	programAstNode := visitor.VisitProgram(programContext.(*parser.ProgramContext)) // Cast to concrete context type

	// Type assert the result to *ast.Program
	program, ok := programAstNode.(*ast.Program)
	if !ok {
		// It's helpful to know what type was actually returned if the assertion fails.
		actualType := "<nil>"
		if programAstNode != nil {
			actualType = fmt.Sprintf("%T", programAstNode)
		}
		fmt.Fprintf(os.Stderr, "Error: VisitProgram did not return an *ast.Program. Got %s\n", actualType)
		return nil, fmt.Errorf("AST building failed: incorrect program node type. Expected *ast.Program, got %s", actualType)
	}

	// Print the AST if requested
	if printAst {
		fmt.Println("Abstract Syntax Tree:")
		fmt.Println(program.String()) // Assumes program has a String() method
	}

	return program, nil
}

func executeCode(code string, source string, isRepl bool, debugMode bool) (interpreter.Value, *interpreter.Interpreter, error) {
	astBuilder := NewAstBuilder(debugMode)
	exprParserFn := func(expressionString string) (ast.Expression, error) {
		return astBuilder.ParseExpressionFromString(expressionString)
	}
	interp := interpreter.NewInterpreter(debugMode, exprParserFn)
	val, err := executeCodeWithInterpreter(interp, code, source, isRepl, debugMode)
	return val, interp, err
}

func executeCodeWithInterpreter(interp *interpreter.Interpreter, code string, source string, isRepl bool, debugMode bool) (interpreter.Value, error) {
	// Create lexer and parser
	inputStream := antlr.NewInputStream(code)
	lexer := parser.NewVLangCherryLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewVLangCherryParser(tokenStream)

	// Add error listener
	errorListener := NewErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the input
	programContext := p.Program()

	// Check for syntax errors
	if HasErrors() { // Use global HasErrors
		// Errors are already logged globally by the ErrorListener.
		// PrintErrorsTable() will be called by the top-level command handler (runFile) if needed.
		return nil, fmt.Errorf("syntax errors occurred during execution setup")
	}

	// Build the AST
	visitor := NewAstBuilder(debugMode)
	// Use type assertion to convert the interface to the concrete type expected by the visitor
	program := visitor.VisitProgram(programContext.(*parser.ProgramContext)).(*ast.Program)

	// Execute the program
	result, err := interp.Interpret(program)
	if err != nil {
		return nil, err
	}

	return result, nil
}
