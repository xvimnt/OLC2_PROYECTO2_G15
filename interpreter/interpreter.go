package interpreter

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings" // Used for strings.Join in native functions like print, println

	"github.com/xvimnt/OLC2_PROYECTO2_G15/ast"
	// "github.com/xvimnt/OLC2_PROYECTO2_G15/parser" // Not directly used here, but ast depends on it
)

// getTypeName converts a Value to its string type representation.
// It handles basic types and recursively determines array element types.
// For empty arrays, it returns "[]any" as element type cannot be inferred.
func getTypeName(val Value) string {
	if val == nil {
		// This case handles if a raw nil is passed, though typically values are wrapped, e.g., NilValue
		return "nil"
	}
	switch v := val.(type) {
	case *IntegerValue:
		return "int"
	case *FloatValue:
		return "f64"
	case *BoolValue:
		return "bool"
	case *StringValue:
		return "string"
	case *CharValue:
		return "byte" // V uses byte for char
	case *ArrayValue:
		if len(v.Elements) > 0 {
			// Recursively call getTypeName for the element type
			// This will correctly form "[]int", "[]string", etc.
			// And for nested, it would be "[]" + getTypeName(v.Elements[0]) -> "[][]int"
			return "[]" + getTypeName(v.Elements[0])
		}
		return "[]any" // For empty arrays, element type unknown by pure runtime inference
	case *MapValue:
		// TODO: Enhance for map[K]V if type info (key/value types) becomes available in MapValue
		return "map"
	case *VFunction, *NativeFunction: // Grouping function types
		return "fn"
	// Removed *NilValue case, nil/none handled by val.Type() or initial val == nil check
	case *StructValue:
		// TODO: Enhance to return struct name if available, e.g., v.StructName
		// For now, just "struct". If v.Name is available: return v.Name
		return "struct"
	default:
		// Fallback using the ValueType enum if direct type switch misses something
		// or for types not having a distinct struct (like internal ones if they leak)
		// This part is more of a safeguard.
		switch val.Type() {
		case IntegerType:
			return "int"
		case FloatType:
			return "f64"
		case BoolType:
			return "bool"
		case StringType:
			return "string"
		case CharType:
			return "byte"
		case ArrayType: // Should have been caught by *ArrayValue case
			if arrVal, ok := val.(*ArrayValue); ok {
				if len(arrVal.Elements) > 0 {
					return "[]" + getTypeName(arrVal.Elements[0])
				}
				return "[]any"
			}
			return "array" // Fallback if cast fails somehow
		case MapType:
			return "map"
		case FunctionType:
			return "fn"
		case NoneType: // Assuming NoneType is the correct enum for nil/none values
			return "none" // Or "nil" if preferred for output consistency
		case StructType:
			return "struct"
		default:
			return "unknown" // For any other types like ErrorType, ReturnValueType, etc.
		}
	}
}

// ExpressionParser defines the function signature for parsing an expression string into an AST node.
// This allows the main package (which has AstBuilder) to provide the parsing capability to the interpreter.
type ExpressionParser func(expressionString string) (ast.Expression, error)

// Control Signals
type BreakSignal struct{}
type ContinueSignal struct{}
type ReturnSignal struct{ Value Value }

var (
	BreakSignalValue    = BreakSignal{}
	ContinueSignalValue = ContinueSignal{}
)

func IsBreakSignal(v interface{}) bool {
	_, ok := v.(BreakSignal)
	return ok
}

func IsContinueSignal(v interface{}) bool {
	_, ok := v.(ContinueSignal)
	return ok
}

func IsReturnSignal(v interface{}) (Value, bool) {
	rs, ok := v.(ReturnSignal)
	if ok {
		return rs.Value, true
	}
	return nil, false
}

// RuntimeError represents an error that occurs during program execution
type RuntimeError struct {
	Message string
	Line    int
	Column  int
}

func (e RuntimeError) Error() string {
	return fmt.Sprintf("[line %d:%d] Runtime Error: %s", e.Line, e.Column, e.Message)
}

func IsRuntimeError(v interface{}) bool {
	_, ok := v.(RuntimeError)
	return ok
}

// StructDef stores information about a struct definition
type StructDef struct {
	Name    string
	Fields  map[string]bool // Just field names for now, type info can be added later
	Methods map[string]*VFunction
}

// Interpreter implements the visitor pattern to interpret V AST
type Interpreter struct {
	DebugMode   bool
	globals     *Environment // Global environment
	environment *Environment // Current environment
	// callStack                         *CallStack            // Manages call stack (temporarily commented out)
	expectedCompositeLiteralTypeStack []ast.TypeNode        // Stack to hold expected types for composite literals
	locals                            map[ast.Node]int      // Map for resolved variables (placeholder for resolver)
	structs                           map[string]*StructDef // Defined structs
	exprParser                        ExpressionParser      // Function to parse expression strings
}

// NewInterpreter creates a new interpreter instance
func NewInterpreter(debugMode bool, exprParser ExpressionParser) *Interpreter {
	globals := NewEnvironment(nil)

	interpreter := &Interpreter{
		DebugMode:   debugMode,
		environment: globals, // Current environment starts as global
		globals:     globals, // Global environment
		// callStack:                         NewCallStack(), // (temporarily commented out)
		expectedCompositeLiteralTypeStack: make([]ast.TypeNode, 0), // Initialize the stack
		locals:                            make(map[ast.Node]int),
		structs:                           make(map[string]*StructDef),
		exprParser:                        exprParser, // Store the provided expression parser
	}

	// Add native functions to global environment
	interpreter.defineNativeFunctions()

	return interpreter
}

// defineNativeFunctions adds built-in functions to the global environment
// Assuming Value, NewNativeFunction, NewString, None, etc. are defined in values.go or similar
func (i *Interpreter) defineNativeFunctions() {
	// Print function
	i.globals.Define("print", NewNativeFunction("print", nil, // Parameter names are illustrative for varargs
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			printer := GetNativePrintSingleton()
			if len(arguments) > 0 {
				var parts []string
				for _, arg := range arguments {
					parts = append(parts, arg.String())
				}
				printer.Print(strings.Join(parts, " ")) // Join with space and print
			}
			return None, nil // Assuming None is a defined Value
		}))

	// Println function
	i.globals.Define("println", NewNativeFunction("println", nil, // Parameter names are illustrative for varargs
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			printer := GetNativePrintSingleton()
			if len(arguments) == 0 {
				printer.Println("")
			} else {
				var parts []string
				for _, arg := range arguments {
					parts = append(parts, arg.String())
				}
				printer.Println(strings.Join(parts, " ")) // Join with space and print
			}
			return None, nil // Assuming None is a defined Value
		}))

	// typeOf function
	// Uses the package-level getTypeName helper
	i.globals.Define("typeOf", NewNativeFunction("typeOf", []string{"value"},
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) != 1 {
				return nil, fmt.Errorf("typeOf expects 1 argument, got %d", len(arguments))
			}
			// Use the getTypeName helper to determine the type string
			typeName := getTypeName(arguments[0])
			return NewString(typeName), nil
		}))

	// len function
	i.globals.Define("len", NewNativeFunction("len", []string{"collection"},
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) != 1 {
				return nil, fmt.Errorf("len expects 1 argument, got %d", len(arguments))
			}
			arg := arguments[0]
			switch value := arg.(type) {
			case *ArrayValue:
				return NewIntegerValue(int64(len(value.Elements))), nil
			case *StringValue:
				return NewIntegerValue(int64(len(value.Value))), nil // V's len on string is byte length
			case *MapValue:
				return NewIntegerValue(int64(len(value.Elements))), nil
			default:
				return nil, fmt.Errorf("len cannot be applied to type %s", arg.Type().String())
			}
		}))

	// append function
	i.globals.Define("append", NewNativeFunction("append", []string{"slice", "elements..."}, // "elements..." indicates variadic
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) < 1 { // Must have at least the slice
				return nil, fmt.Errorf("append expects at least 1 argument (the slice), got %d", len(arguments))
			}

			sliceVal, ok := arguments[0].(*ArrayValue)
			if !ok {
				return nil, fmt.Errorf("first argument to append must be a slice/array, got %s", arguments[0].Type().String())
			}

			// Create a new slice for the elements.
			// Start with the original elements.
			// The capacity can be len(sliceVal.Elements) + len(arguments) - 1 for slight optimization
			newElements := make([]Value, len(sliceVal.Elements), len(sliceVal.Elements)+len(arguments)-1)
			copy(newElements, sliceVal.Elements)

			// Add the new elements to append
			if len(arguments) > 1 {
				elementsToAppend := arguments[1:]
				newElements = append(newElements, elementsToAppend...)
			}

			// Create a new ArrayValue with the combined elements.
			// The type of the new array should ideally match the original array's declared type,
			// but our ArrayValue doesn't explicitly store this beyond its elements.
			return NewArray(newElements), nil
		}))

	// join function
	i.globals.Define("join", NewNativeFunction("join", []string{"slice", "separator"},
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) != 2 {
				return nil, fmt.Errorf("join expects 2 arguments (slice, separator), got %d", len(arguments))
			}

			sliceVal, ok := arguments[0].(*ArrayValue)
			if !ok {
				return nil, fmt.Errorf("first argument to join must be a slice/array, got %s", arguments[0].Type().String())
			}

			separatorVal, ok := arguments[1].(*StringValue)
			if !ok {
				return nil, fmt.Errorf("second argument to join must be a string, got %s", arguments[1].Type().String())
			}

			if len(sliceVal.Elements) == 0 {
				return NewString(""), nil // Return empty string if slice is empty
			}

			var stringElements []string
			for idx, el := range sliceVal.Elements {
				strEl, ok := el.(*StringValue)
				if !ok {
					return nil, fmt.Errorf("all elements in the slice for join must be strings, found %s at index %d", el.Type().String(), idx)
				}
				stringElements = append(stringElements, strEl.Value)
			}

			joinedString := strings.Join(stringElements, separatorVal.Value)
			return NewString(joinedString), nil
		}))

	// indexOf function
	i.globals.Define("indexOf", NewNativeFunction("indexOf", []string{"slice", "element"},
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) != 2 {
				return nil, fmt.Errorf("indexOf expects 2 arguments (slice, element), got %d", len(arguments))
			}

			sliceVal, ok := arguments[0].(*ArrayValue)
			if !ok {
				return nil, fmt.Errorf("first argument to indexOf must be a slice/array, got %s", arguments[0].Type().String())
			}

			elementToFind := arguments[1]

			for idx, el := range sliceVal.Elements {
				if el.Equals(elementToFind) { // Corrected call to Equals
					return NewIntegerValue(int64(idx)), nil // Corrected constructor
				}
			}

			return NewIntegerValue(-1), nil // Corrected constructor: Element not found
		}))

	// atoi function
	i.globals.Define("atoi", NewNativeFunction("atoi", []string{"s"},
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) != 1 {
				return nil, fmt.Errorf("atoi expects 1 argument, got %d", len(arguments))
			}

			strVal, ok := arguments[0].(*StringValue)
			if !ok {
				return nil, fmt.Errorf("atoi expects a string argument, got %s", arguments[0].Type().String())
			}

			// Use strconv.ParseInt to convert string to int64
			intValue, err := strconv.ParseInt(strVal.Value, 10, 64)
			if err != nil {
				// V's atoi returns an optional. For now, we can return an error or a default value.
				// Returning an error is more informative for debugging.
				// A more complete implementation might return a V-level optional/result type.
				return nil, fmt.Errorf("invalid string for atoi: '%s' cannot be converted to an integer", strVal.Value)
			}

			return NewIntegerValue(intValue), nil
		}))

	// parseFloat function
	i.globals.Define("parseFloat", NewNativeFunction("parseFloat", []string{"s"},
		func(interpreter *Interpreter, arguments []Value) (Value, error) {
			if len(arguments) != 1 {
				return nil, fmt.Errorf("parseFloat expects 1 argument, got %d", len(arguments))
			}

			strVal, ok := arguments[0].(*StringValue)
			if !ok {
				return nil, fmt.Errorf("parseFloat expects a string argument, got %s", arguments[0].Type().String())
			}

			// Use strconv.ParseFloat to convert string to float64
			floatValue, err := strconv.ParseFloat(strVal.Value, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid string for parseFloat: '%s' cannot be converted to a float", strVal.Value)
			}

			return NewFloatValue(floatValue), nil
		}))

	// Alias Atoi to atoi for convenience, as some languages use this convention.
	atoiFunc, err := i.globals.Get("atoi")
	if err == nil {
		i.globals.Define("Atoi", atoiFunc)
	}
}

// Interpret executes a V program represented as an AST
func (i *Interpreter) Interpret(program *ast.Program) (Value, error) {
	if i.DebugMode {
		fmt.Println("Interpreting program...")
	}

	// --- Pass 1: Register all struct definitions ---
	for _, decl := range program.Declarations {
		if structDecl, ok := decl.(*ast.StructDecl); ok {
			result := structDecl.Accept(i)
			if err, isErr := result.(RuntimeError); isErr {
				return nil, err // Propagate any errors during struct registration
			}
		}
	}

	// --- Pass 2: Process all global variable declarations ---
	for _, decl := range program.Declarations {
		if varDecl, ok := decl.(*ast.VarDecl); ok {
			result := varDecl.Accept(i)
			if err, isErr := result.(RuntimeError); isErr {
				return nil, err // Propagate any errors during var declaration
			}
		}
	}

	// --- Pass 3: Register all function definitions ---
	for _, decl := range program.Declarations {
		if fnDecl, ok := decl.(*ast.FunctionDecl); ok {
			result := fnDecl.Accept(i)
			if err, isErr := result.(RuntimeError); isErr {
				return nil, err // Propagate any errors during function registration
			}
		}
	}

	// --- Pass 4: Find and execute the main function ---
	mainFuncVal, err := i.globals.Get("main")
	if err != nil {
		return nil, RuntimeError{Message: "main function not found"}
	}

	mainFunc, ok := mainFuncVal.(*VFunction)
	if !ok {
		return nil, RuntimeError{Message: "'main' is not a function"}
	}

	// Call the main function
	result, callErr := mainFunc.Call(i, []Value{})
	if callErr != nil {
		// Check if the error from Call is already a RuntimeError
		if re, ok := callErr.(RuntimeError); ok {
			return nil, re
		}
		return nil, RuntimeError{Message: fmt.Sprintf("error executing main function: %s", callErr.Error())}
	}

	return result, nil
}

// execute runs a statement or declaration (stubbed)
func (i *Interpreter) execute(node ast.Node) (Value, error) {
	result := node.Accept(i)
	if val, ok := result.(Value); ok {
		return val, nil
	}
	if err, ok := result.(error); ok {
		return nil, err
	}
	return None, nil // Default return
}

// evaluate resolves an expression to a value (stubbed)
func (i *Interpreter) evaluate(expr ast.Expression) (Value, error) {
	if expr == nil {
		return None, nil
	}
	result := expr.Accept(i)
	if val, ok := result.(Value); ok {
		return val, nil
	}
	if err, ok := result.(error); ok {
		return nil, err
	}
	return None, nil // Default return
}

// executeBlock executes a list of statements in a given environment.
// It handles 'return' statements by returning a ReturnValue.
func (i *Interpreter) executeBlock(statements []ast.Statement, environment *Environment) interface{} {
	previous := i.environment
	i.environment = environment
	defer func() { i.environment = previous }()

	for _, stmt := range statements {
		result := stmt.Accept(i)
		if i.DebugMode {
			fmt.Fprintf(os.Stderr, "Interpreter.executeBlock: Statement %T executed, result type: %T, result value: %v\n", stmt, result, result)
		}
		// Check for control signals or runtime errors
		if _, isRet := IsReturnSignal(result); isRet {
			return result // Propagate return signal (which is the ReturnSignal struct)
		}
		if IsBreakSignal(result) {
			return result // Propagate break signal
		}
		if IsContinueSignal(result) {
			return result // Propagate continue signal
		}
		if IsRuntimeError(result) {
			return result // Propagate runtime error
		}
	}
	return nil // Default if no return, break, or continue encountered
}

// isTruthy evaluates if a value is considered true in a boolean context.
// V's truthiness:
// - false is false
// - none is false
// - 0 is true (unlike some languages)
// - empty string is true (unlike some languages)
// - empty arrays/maps are true
// Essentially, only explicit `false` and `none` are falsy.
func (i *Interpreter) isTruthy(value Value) bool {
	if value == nil {
		// This case should ideally be prevented by evaluate always returning a Value or an error.
		// If it does happen, treat it as falsy to be safe.
		if i.DebugMode {
			fmt.Println("isTruthy received a nil Value interface")
		}
		return false
	}
	switch v := value.(type) {
	case *BoolValue:
		return v.Value
	case *NoneValue:
		return false
	default:
		// All other concrete value types (Integer, Float, String, Array, Map, Function, Struct, etc.)
		// are considered truthy.
		return true
	}
}

// Visitor methods implementation with println

func (i *Interpreter) VisitProgram(node *ast.Program) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting Program")
	}
	for _, decl := range node.Declarations {
		decl.Accept(i)
	}
	return nil
}

// Declarations
func (i *Interpreter) VisitStructDecl(node *ast.StructDecl) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting StructDecl: %s\n", node.Name.Name)
	}

	// Create a map of field names for quick lookup.
	fields := make(map[string]bool)
	for _, field := range node.Fields {
		fields[field.Name.Name] = true
	}

	// Create the struct definition.
	structDef := &StructDef{
		Name:    node.Name.Name,
		Fields:  fields,
		Methods: make(map[string]*VFunction), // Initialize methods map
	}

	// Store the struct definition in the interpreter's global map of structs.
	i.structs[structDef.Name] = structDef

	if i.DebugMode {
		// Create a slice of field names for printing.
		fieldNames := make([]string, 0, len(fields))
		for name := range fields {
			fieldNames = append(fieldNames, name)
		}
		fmt.Printf("  Defined struct '%s' with fields: %v\n", structDef.Name, fieldNames)
	}

	// A declaration does not produce a value.
	return nil
}

func (i *Interpreter) VisitFieldDecl(node *ast.FieldDecl) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting FieldDecl: %s of type %s\n", node.Name.Name, node.Type.String())
	}
	return nil
}

func (i *Interpreter) VisitFunctionDecl(node *ast.FunctionDecl) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting FunctionDecl: %s\n", node.Name.Name)
	}

	paramNames := make([]string, len(node.Parameters))
	for idx, param := range node.Parameters {
		paramNames[idx] = param.Name.Name // Assuming IdentifierExpr has a Name string field
	}

	// The closure for a function is the environment where it's defined.
	// For top-level functions, this will be the global environment.
	// For nested functions (if supported later), it would be the lexical scope.
	closure := i.environment

	// The VFunction stores the AST node of the declaration and its closure.
	// The Name and ParamNames are also stored for easier access/debugging.
	if i.DebugMode {
		fmt.Printf("[DEBUG] VisitFunctionDecl: Capturing closure for '%s' in env %p\n", node.Name.Name, closure)
	}
	vFunc := NewVFunction(node, closure, node.Name.Name, paramNames)

	// Define the function in the current environment.
	// For 'main' and other top-level functions, this should be the global environment.
	if node.Receiver != nil {
		// This is a method, associate it with its struct definition.
		var receiverTypeName string
		switch rt := node.Receiver.Type.(type) {
		case *ast.TypeName:
			receiverTypeName = rt.Name
		case *ast.PointerTypeNode:
			switch et := rt.ElementType.(type) {
			case *ast.TypeName:
				receiverTypeName = et.Name
			default:
				// TODO: Handle or return error for other element types if necessary
				// For now, assume direct or pointer to TypeName for receivers
				return RuntimeError{Message: fmt.Sprintf("unsupported receiver element type: %T for method %s", et, node.Name.Name), Line: 0, Column: 0} // AST node missing Pos()
			}
		default:
			return RuntimeError{Message: fmt.Sprintf("unsupported receiver type: %T for method %s", rt, node.Name.Name), Line: 0, Column: 0} // AST node missing Pos()
		}

		if receiverTypeName == "" {
			return RuntimeError{Message: fmt.Sprintf("could not determine receiver type name for method %s", node.Name.Name), Line: 0, Column: 0} // AST node missing Pos()
		}

		structDef, ok := i.structs[receiverTypeName]
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("struct %s not defined for method %s", receiverTypeName, node.Name.Name), Line: 0, Column: 0} // AST node missing Pos()
		}

		structDef.Methods[node.Name.Name] = vFunc
		if i.DebugMode {
			fmt.Printf("[DEBUG] VisitFunctionDecl: Registered method '%s' for struct '%s'\n", node.Name.Name, receiverTypeName)
		}
	} else {
		// This is a regular function, define it in the current environment.
		i.environment.Define(node.Name.Name, vFunc)
		if i.DebugMode {
			fmt.Printf("[DEBUG] VisitFunctionDecl: Registered function '%s' in env %p\n", node.Name.Name, i.environment)
		}
	}

	return None // Function definition itself doesn't produce a runtime value.
}

func (i *Interpreter) VisitReceiverDecl(node *ast.ReceiverDecl) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting ReceiverDecl: %s of type %s\n", node.Name, node.Type.String())
	}
	return nil
}

func (i *Interpreter) VisitParameterDecl(node *ast.ParameterDecl) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting ParameterDecl: %s\n", node.Name.Name)
	}
	return nil
}

func (i *Interpreter) VisitVarDecl(node *ast.VarDecl) interface{} {
	varName := node.Name.Name
	if i.DebugMode {
		var explicitTypeStr string
		if node.ExplicitType != nil {
			explicitTypeStr = node.ExplicitType.String()
		} else {
			explicitTypeStr = "<inferred>"
		}
		fmt.Fprintf(os.Stderr, "Interpreter.VisitVarDecl (%s): AST ExplicitType: %s\n", varName, explicitTypeStr)
	}

	var value Value

	if node.Initializer != nil {
		if i.DebugMode {
			fmt.Fprintf(os.Stderr, "Interpreter.Visiting VarDecl: %s. AST Initializer: %s\n", varName, node.Initializer.String())
		}

		// If there's an explicit type, push it onto the stack for the composite literal evaluator to use.
		if node.ExplicitType != nil {
			i.expectedCompositeLiteralTypeStack = append(i.expectedCompositeLiteralTypeStack, node.ExplicitType)
			defer func() {
				// Ensure the stack is popped after evaluation.
				if len(i.expectedCompositeLiteralTypeStack) > 0 {
					i.expectedCompositeLiteralTypeStack = i.expectedCompositeLiteralTypeStack[:len(i.expectedCompositeLiteralTypeStack)-1]
				}
			}()
		}

		initVal, err := i.evaluate(node.Initializer)
		if err != nil {
			if i.DebugMode {
				fmt.Fprintf(os.Stderr, "Interpreter.VisitVarDecl (%s): Error evaluating initializer: %v\n", varName, err)
			}
			return err // Propagate error
		}
		value = initVal

		if i.DebugMode {
			fmt.Fprintf(os.Stderr, "Interpreter.VisitVarDecl (%s): Evaluated initializer to: %s (Type: %s)\n", varName, value.String(), value.Type())
		}
	} else {
		// No initializer, assign zero value based on declared type.
		value = None // Default to None
		if node.ExplicitType != nil {
			// This is a simplified zero-value logic. A more robust implementation
			// would probably involve a helper function on the types themselves.
			switch t := node.ExplicitType.(type) {
			case *ast.PrimitiveTypeNode:
				switch t.Kind {
				case ast.IntKind:
					value = NewIntegerValue(0)
				case ast.Float64Kind:
					value = NewFloatValue(0.0)
				case ast.StringKind:
					value = NewString("")
				case ast.BoolKind:
					value = NewBoolValue(false)
				}
			case *ast.TypeName:
				switch t.Name {
				case "int":
					value = NewIntegerValue(0)
				case "float", "float64":
					value = NewFloatValue(0.0)
				case "string":
					value = NewString("")
				case "bool":
					value = NewBoolValue(false)
				}
			case *ast.SliceTypeNode, *ast.ArrayTypeNode:
				// The zero value for a slice or array is `nil`. We'll represent this
				// with an empty ArrayValue for now.
				value = NewArray([]Value{})
			}
		}
		if i.DebugMode {
			fmt.Fprintf(os.Stderr, "Interpreter.VisitVarDecl (%s): No initializer, assigned zero value: %s\n", varName, value.String())
		}
	}

	i.environment.Define(varName, value)
	if i.DebugMode {
		fmt.Fprintf(os.Stderr, "Interpreter.Defined variable in environment: %s = %s (Type: %s, Declared Mutable: %t)\n", varName, value.String(), value.Type(), node.IsMutable)
	}
	return value
}

func (i *Interpreter) VisitBlockStmt(node *ast.BlockStmt) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting BlockStmt with %d statements\n", len(node.Statements))
	}
	// Create a new environment for the block, inheriting from the current one.
	return i.executeBlock(node.Statements, NewEnvironment(i.environment))
}

func (i *Interpreter) VisitAssignStmt(node *ast.AssignStmt) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting AssignStmt: LValue: %s, Op: %s, RValue: %s\n", node.Left.String(), node.Operator, node.Right.String())
	}

	rhsValue, err := i.evaluate(node.Right)
	if err != nil {
		return err // Propagate error from RHS evaluation
	}

	switch lvalue := node.Left.(type) {
	case *ast.IdentifierExpr:
		varName := lvalue.Name
		if i.DebugMode {
			fmt.Printf("  Assigning to identifier: %s\n", varName)
		}

		switch node.Operator {
		case "=":
			err := i.environment.Assign(varName, rhsValue)
			if err != nil {
				return RuntimeError{
					Message: fmt.Sprintf("Error assigning to '%s': %v", varName, err),
					Line:    node.Line,
					Column:  node.Column,
				}
			}
			return nil // Assignment successful
		case "+=":
			currentValue, getErr := i.environment.Get(varName)
			if getErr != nil {
				return RuntimeError{
					Message: fmt.Sprintf("Error getting current value of '%s' for '+=': %v", varName, getErr),
					Line:    node.Line,
					Column:  node.Column,
				}
			}

			// Handle Integer += Integer or Integer += Float
			if currentInt, okC := currentValue.(*IntegerValue); okC {
				if rhsInt, okR := rhsValue.(*IntegerValue); okR { // Integer += Integer
					resultValue := NewIntegerValue(currentInt.Value + rhsInt.Value)
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '+=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				} else if rhsFloat, okR := rhsValue.(*FloatValue); okR { // Integer += Float
					resultValue := NewFloatValue(float64(currentInt.Value) + rhsFloat.Value)
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '+=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				}
			} else if currentFloat, okC := currentValue.(*FloatValue); okC { // Handle Float += Integer or Float += Float
				if rhsFloat, okR := rhsValue.(*FloatValue); okR { // Float += Float
					resultValue := NewFloatValue(currentFloat.Value + rhsFloat.Value)
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '+=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				} else if rhsInt, okR := rhsValue.(*IntegerValue); okR { // Float += Integer
					resultValue := NewFloatValue(currentFloat.Value + float64(rhsInt.Value))
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '+=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				}
			} else if currentString, okC := currentValue.(*StringValue); okC { // Handle String += String (concatenation)
				if rhsString, okR := rhsValue.(*StringValue); okR {
					resultValue := NewString(currentString.Value + rhsString.Value) // Corrected constructor name
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '+=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				}
			}
			// If no specific type combination matched, return an error
			return RuntimeError{
				Message: fmt.Sprintf("Operands for '+=' have incompatible types: %s and %s.", currentValue.Type(), rhsValue.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		case "-=":
			currentValue, getErr := i.environment.Get(varName)
			if getErr != nil {
				return RuntimeError{
					Message: fmt.Sprintf("Error getting current value of '%s' for '-=': %v", varName, getErr),
					Line:    node.Line,
					Column:  node.Column,
				}
			}

			// Handle Integer -= Integer
			if currentInt, okC := currentValue.(*IntegerValue); okC {
				if rhsInt, okR := rhsValue.(*IntegerValue); okR {
					resultValue := NewIntegerValue(currentInt.Value - rhsInt.Value)
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '-=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				} else if rhsFloat, okR := rhsValue.(*FloatValue); okR { // Handle Integer -= Float
					resultValue := NewFloatValue(float64(currentInt.Value) - rhsFloat.Value)
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '-=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				}
			} else if currentFloat, okC := currentValue.(*FloatValue); okC { // Handle Float -= ...
				if rhsFloat, okR := rhsValue.(*FloatValue); okR { // Handle Float -= Float
					resultValue := NewFloatValue(currentFloat.Value - rhsFloat.Value)
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '-=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				} else if rhsInt, okR := rhsValue.(*IntegerValue); okR { // Handle Float -= Integer
					resultValue := NewFloatValue(currentFloat.Value - float64(rhsInt.Value))
					err := i.environment.Assign(varName, resultValue)
					if err != nil {
						return RuntimeError{
							Message: fmt.Sprintf("Error assigning result of '-=' to '%s': %v", varName, err),
							Line:    node.Line,
							Column:  node.Column,
						}
					}
					return nil
				}
			}
			return RuntimeError{
				Message: fmt.Sprintf("Operands for '-=' must be numbers. Got %s and %s.", currentValue.Type(), rhsValue.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}

		// TODO: Implement other compound assignment operators like *=, /= etc.
		default:
			// NodeWithPosition was problematic, using 0,0 for now
			return RuntimeError{
				Message: fmt.Sprintf("Unsupported assignment operator '%s' for identifier '%s'.", node.Operator, varName),
				Line:    node.Line,
				Column:  node.Column,
			}
		}

	case *ast.FieldAccessExpr:
		fieldAccessExpr := lvalue // lvalue is *ast.FieldAccessExpr
		// rhsValue is already evaluated at the top of VisitAssignStmt

		if i.DebugMode {
			fmt.Printf("VisitAssignStmt: FieldAccessExpr. Receiver: %s, Field: %s, Operator: %s, RHS: %s\n",
				fieldAccessExpr.Receiver.String(),
				fieldAccessExpr.Field.Name,
				node.Operator,
				rhsValue.String())
		}

		// 1. Evaluate the receiver of the field access (e.g., 'p' in p.age)
		receiverResultRaw := fieldAccessExpr.Receiver.Accept(i)
		if err, isErr := receiverResultRaw.(RuntimeError); isErr {
			return err
		}
		receiverValue, ok := receiverResultRaw.(Value)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("receiver expression for field assignment '%s.%s' did not evaluate to a usable Value, got %T: %v",
					fieldAccessExpr.Receiver.String(), fieldAccessExpr.Field.Name, receiverResultRaw, receiverResultRaw),
				Line:   node.Line,
				Column: node.Column, // TODO: Get proper node position
			}
		}

		// Automatically dereference pointers to structs for assignment
		if ptr, isPtr := receiverValue.(*PointerValue); isPtr {
			dereferencedValue, derefErr := ptr.Dereference()
			if derefErr != nil {
				return RuntimeError{
					Message: derefErr.Error(),
					Line:    node.Line,
					Column:  node.Column,
				}
			}
			receiverValue = dereferencedValue // Use the dereferenced value for field access
		}

		structVal, ok := receiverValue.(*StructValue)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("Cannot assign to field '%s' of non-struct type '%s'. Receiver evaluated to: %s",
					fieldAccessExpr.Field.Name, receiverValue.Type(), receiverValue.String()),
				Line:   node.Line,
				Column: node.Column, // TODO: Get proper node position
			}
		}

		fieldName := fieldAccessExpr.Field.Name

		// 2. Handle assignment based on operator
		switch node.Operator {
		case "=":
			oldValue, _ := structVal.Fields[fieldName] // For debug logging
			structVal.Fields[fieldName] = rhsValue
			if i.DebugMode {
				fmt.Printf("  Assigned to struct field: %s.%s = %s (was %s)\n",
					structVal.TypeName, fieldName, rhsValue.String(), oldValue.String())
			}
			return rhsValue // Assignment expression evaluates to the assigned value

		case "+=", "-=", "*=", "/=", "%= ":
			currentFieldValue, exists := structVal.Fields[fieldName]
			if !exists {
				return RuntimeError{
					Message: fmt.Sprintf("Field '%s' not found in struct '%s' for compound assignment '%s'",
						fieldName, structVal.TypeName, node.Operator),
					Line:   node.Line,
					Column: node.Column, // TODO: Get proper node position
				}
			}

			var operationResult Value
			var opErr error
			opSymbol := node.Operator[:len(node.Operator)-1] // e.g., "+" from "+="

			switch opSymbol {
			case "+":
				if cvInt, okCv := currentFieldValue.(*IntegerValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						operationResult = NewIntegerValue(cvInt.Value + rvInt.Value)
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						operationResult = NewFloatValue(float64(cvInt.Value) + rvFloat.Value)
					} else {
						opErr = fmt.Errorf("cannot add %s to integer field '%s'", rhsValue.Type(), fieldName)
					}
				} else if cvFloat, okCv := currentFieldValue.(*FloatValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						operationResult = NewFloatValue(cvFloat.Value + float64(rvInt.Value))
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						operationResult = NewFloatValue(cvFloat.Value + rvFloat.Value)
					} else {
						opErr = fmt.Errorf("cannot add %s to float field '%s'", rhsValue.Type(), fieldName)
					}
				} else if cvString, okCv := currentFieldValue.(*StringValue); okCv {
					if rvString, okRv := rhsValue.(*StringValue); okRv {
						operationResult = NewString(cvString.Value + rvString.Value)
					} else {
						opErr = fmt.Errorf("cannot concatenate %s with string field '%s'", rhsValue.Type(), fieldName)
					}
				} else {
					opErr = fmt.Errorf("field '%s' of type %s does not support '+' operator", fieldName, currentFieldValue.Type())
				}
			case "-":
				if cvInt, okCv := currentFieldValue.(*IntegerValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						operationResult = NewIntegerValue(cvInt.Value - rvInt.Value)
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						operationResult = NewFloatValue(float64(cvInt.Value) - rvFloat.Value)
					} else {
						opErr = fmt.Errorf("cannot subtract %s from integer field '%s'", rhsValue.Type(), fieldName)
					}
				} else if cvFloat, okCv := currentFieldValue.(*FloatValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						operationResult = NewFloatValue(cvFloat.Value - float64(rvInt.Value))
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						operationResult = NewFloatValue(cvFloat.Value - rvFloat.Value)
					} else {
						opErr = fmt.Errorf("cannot subtract %s from float field '%s'", rhsValue.Type(), fieldName)
					}
				} else {
					opErr = fmt.Errorf("field '%s' of type %s does not support '-' operator", fieldName, currentFieldValue.Type())
				}
			case "*":
				if cvInt, okCv := currentFieldValue.(*IntegerValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						operationResult = NewIntegerValue(cvInt.Value * rvInt.Value)
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						operationResult = NewFloatValue(float64(cvInt.Value) * rvFloat.Value)
					} else {
						opErr = fmt.Errorf("cannot multiply integer field '%s' by %s", fieldName, rhsValue.Type())
					}
				} else if cvFloat, okCv := currentFieldValue.(*FloatValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						operationResult = NewFloatValue(cvFloat.Value * float64(rvInt.Value))
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						operationResult = NewFloatValue(cvFloat.Value * rvFloat.Value)
					} else {
						opErr = fmt.Errorf("cannot multiply float field '%s' by %s", fieldName, rhsValue.Type())
					}
				} else {
					opErr = fmt.Errorf("field '%s' of type %s does not support '*' operator", fieldName, currentFieldValue.Type())
				}
			case "/":
				if cvInt, okCv := currentFieldValue.(*IntegerValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						if rvInt.Value == 0 {
							opErr = fmt.Errorf("division by zero")
						} else {
							operationResult = NewIntegerValue(cvInt.Value / rvInt.Value) // Integer division
						}
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						if rvFloat.Value == 0.0 {
							opErr = fmt.Errorf("division by zero")
						} else {
							operationResult = NewFloatValue(float64(cvInt.Value) / rvFloat.Value)
						}
					} else {
						opErr = fmt.Errorf("cannot divide integer field '%s' by %s", fieldName, rhsValue.Type())
					}
				} else if cvFloat, okCv := currentFieldValue.(*FloatValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						if rvInt.Value == 0 {
							opErr = fmt.Errorf("division by zero")
						} else {
							operationResult = NewFloatValue(cvFloat.Value / float64(rvInt.Value))
						}
					} else if rvFloat, okRv := rhsValue.(*FloatValue); okRv {
						if rvFloat.Value == 0.0 {
							opErr = fmt.Errorf("division by zero")
						} else {
							operationResult = NewFloatValue(cvFloat.Value / rvFloat.Value)
						}
					} else {
						opErr = fmt.Errorf("cannot divide float field '%s' by %s", fieldName, rhsValue.Type())
					}
				} else {
					opErr = fmt.Errorf("field '%s' of type %s does not support '/' operator", fieldName, currentFieldValue.Type())
				}
			case "%":
				if cvInt, okCv := currentFieldValue.(*IntegerValue); okCv {
					if rvInt, okRv := rhsValue.(*IntegerValue); okRv {
						if rvInt.Value == 0 {
							opErr = fmt.Errorf("modulo by zero")
						} else {
							operationResult = NewIntegerValue(cvInt.Value % rvInt.Value)
						}
					} else {
						opErr = fmt.Errorf("cannot apply modulo to integer field '%s' with %s", fieldName, rhsValue.Type())
					}
				} else {
					opErr = fmt.Errorf("field '%s' of type %s does not support '%%' operator", fieldName, currentFieldValue.Type())
				}
			default:
				opErr = fmt.Errorf("unsupported compound assignment operator symbol '%s' for struct field", opSymbol)
			}

			if opErr != nil {
				return RuntimeError{
					Message: fmt.Sprintf("Error during compound assignment '%s' on field '%s.%s': %v",
						node.Operator, structVal.TypeName, fieldName, opErr),
					Line:   node.Line,
					Column: node.Column,
				}
			}

			structVal.Fields[fieldName] = operationResult
			if i.DebugMode {
				fmt.Printf("  Compound assigned to struct field: %s.%s %s %s -> %s (was %s)\n",
					structVal.TypeName, fieldName, node.Operator, rhsValue.String(), operationResult.String(), currentFieldValue.String())
			}
			return operationResult // Compound assignment also returns the resulting value

		default:
			return RuntimeError{
				Message: fmt.Sprintf("Unsupported assignment operator '%s' for struct field '%s.%s'",
					node.Operator, structVal.TypeName, fieldName),
				Line:   node.Line,
				Column: node.Column,
			}
		}

	case *ast.IndexAccessExpr:
		idxAccessExpr := lvalue // lvalue is *ast.IndexAccessExpr
		if i.DebugMode {
			fmt.Printf("  Assigning to index: %s\n", idxAccessExpr.String())
		}

		// 1. Evaluate the receiver of the index expression (e.g., 'numbers' in numbers[0])
		receiverValRaw := idxAccessExpr.Receiver.Accept(i)
		if err, ok := receiverValRaw.(RuntimeError); ok {
			return err
		}
		receiverValue, ok := receiverValRaw.(Value)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("receiver for index assignment did not evaluate to a Value, got %T", receiverValRaw),
				Line:    node.Line,
				Column:  node.Column,
			}
		}

		// 2. Evaluate the index expression (e.g., '0' in numbers[0])
		indexValRaw := idxAccessExpr.Index.Accept(i)
		if err, ok := indexValRaw.(RuntimeError); ok {
			return err
		}
		indexValue, ok := indexValRaw.(Value)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("index for assignment did not evaluate to a Value, got %T", indexValRaw),
				Line:    node.Line,
				Column:  node.Column,
			}
		}

		// 3. Perform assignment based on receiver type
		switch collection := receiverValue.(type) {
		case *ArrayValue:
			intIndex, ok := indexValue.(*IntegerValue)
			if !ok {
				return RuntimeError{
					Message: fmt.Sprintf("array index for assignment must be an integer, got %s", indexValue.Type()),
					Line:    node.Line,
					Column:  node.Column,
				}
			}
			idx := int(intIndex.Value)

			// Check mutability: The ArrayValue itself doesn't store mutability,
			// but the original variable declaration does.
			// We need to ensure the variable itself was declared mutable.
			// This check is implicitly handled by Environment.Assign if we were re-assigning the whole array.
			// For element assignment, we rely on the fact that if we have a reference to ArrayValue,
			// it's generally modifiable unless V introduces more complex ownership/borrowing.
			// For now, we assume if we can get the ArrayValue, we can modify its elements.
			// A more robust check might involve looking up the original variable's mutability flag
			// if the receiver is a simple IdentifierExpr.

			if idx < 0 || idx >= len(collection.Elements) {
				return RuntimeError{
					Message: fmt.Sprintf("array index out of bounds for assignment: %d with length %d", idx, len(collection.Elements)),
					Line:    node.Line,
					Column:  node.Column,
				}
			}

			// V is statically typed, but our interpreter is dynamic.
			// We might want to check if rhsValue.Type() is compatible with the array's element type.
			// For now, we'll allow assigning any type, like in some dynamic languages.
			// A more advanced interpreter would get the array's declared element type.

			if node.Operator == "=" {
				collection.Elements[idx] = rhsValue
				if i.DebugMode {
					fmt.Printf("    Assigned value %s to array index %d. New array: %s\n", rhsValue.String(), idx, collection.String())
				}
				return rhsValue // V assignments can be expressions, return the assigned value
			} else {
				// TODO: Handle compound assignments like arr[idx] += value
				return RuntimeError{
					Message: fmt.Sprintf("Unsupported operator '%s' for array index assignment.", node.Operator),
					Line:    node.Line,
					Column:  node.Column,
				}
			}

		// case *MapValue:
		// TODO: Implement map index assignment map[key] = value
		// return RuntimeError{Message: "Map index assignment not yet implemented."}

		default:
			return RuntimeError{
				Message: fmt.Sprintf("type %s is not assignable via index", receiverValue.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}

	default:
		return RuntimeError{
			Message: fmt.Sprintf("Invalid lvalue type '%T' in assignment.", node.Left),
			Line:    node.Line,
			Column:  node.Column,
		}
	}
}

func (i *Interpreter) VisitExpressionStmt(node *ast.ExpressionStmt) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting ExpressionStmt")
	}
	if node.Expression == nil {
		return nil
	}
	result := node.Expression.Accept(i)
	return result
}

func (i *Interpreter) VisitIfStmt(node *ast.IfStmt) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting IfStmt. Condition: %s\n", node.Condition.String())
	}

	conditionValue, err := i.evaluate(node.Condition)
	if err != nil {
		return err
	}

	// If the condition is true, execute the consequence and we're done.
	if i.isTruthy(conditionValue) {
		return node.Consequence.Accept(i)
	}

	// If the condition is false, check the Alternative.
	// The Alternative can be another IfStmt (for 'else if') or a BlockStmt (for 'else').
	if node.Alternative != nil {
		// By simply calling Accept on the Alternative, we recursively handle
		// the entire 'else if'/'else' chain without needing a loop or special cases.
		// The visitor dispatch mechanism will call either VisitIfStmt or VisitBlockStmt as appropriate.
		return node.Alternative.Accept(i)
	}

	// If no branch is taken (e.g., an 'if' without an 'else' that evaluates to false).
	return nil
}

func (i *Interpreter) VisitSwitchStmt(node *ast.SwitchStmt) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting SwitchStmt. Expression: %s\n", node.Expression.String())
	}

	switchValue, errEval := i.evaluate(node.Expression)
	if errEval != nil {
		return errEval // Propagate evaluation error
	}

	var matched bool
	for _, caseClause := range node.Cases {
		for _, caseExpr := range caseClause.Expressions {
			caseValue, errEvalCase := i.evaluate(caseExpr)
			if errEvalCase != nil {
				return errEvalCase // Propagate evaluation error
			}

			if i.isEqual(switchValue, caseValue) { // Use i.isEqual
				matched = true
				// Execute the block of statements for the matched case.
				// The environment for the case block is the same as the switch statement's environment.
				// executeBlock will handle nested environments if necessary for the statements within.
				result := i.executeBlock(caseClause.Body, i.environment)

				if IsBreakSignal(result) { // Check for BreakSignal using the helper
					return nil // Break from switch, effectively consuming the signal
				}
				// If it's a ReturnValue or another error/signal, propagate it.
				if result != nil {
					return result
				}
				// If result is nil (block executed normally), and V needs explicit fallthrough (which it does), we are done with this switch.
				return nil
			}
		}
	}

	if !matched && node.Default != nil {
		// Execute the default block if no cases matched.
		return i.executeBlock(node.Default.Body, i.environment)
	}

	return nil // No case matched, and no default, or default executed and returned nil.
}

// VisitCaseClause is generally not called directly; handled by VisitSwitchStmt.
func (i *Interpreter) VisitCaseClause(node *ast.CaseClause) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting CaseClause (should be handled by SwitchStmt)")
	}
	return RuntimeError{Message: "VisitCaseClause should not be called directly."}
}

// VisitDefaultClause is generally not called directly; handled by VisitSwitchStmt.
func (i *Interpreter) VisitDefaultClause(node *ast.DefaultClause) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting DefaultClause (should be handled by SwitchStmt)")
	}
	return RuntimeError{Message: "VisitDefaultClause should not be called directly."}
}

func (i *Interpreter) VisitForStmt(stmt *ast.ForStmt) interface{} {
	if i.DebugMode {
		initStatus := "nil"
		if stmt.Init != nil {
			initStatus = "present"
		}
		condStatus := "nil"
		if stmt.Condition != nil {
			condStatus = "present"
		}
		postStatus := "nil"
		if stmt.Post != nil {
			postStatus = "present"
		}
		fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt. Init: %s, Condition: %s, Post: %s, IsRange: %v\n",
			initStatus, condStatus, postStatus, stmt.IsRangeLoop)
	}

	if stmt.IsRangeLoop {
		if i.DebugMode {
			rangeKeyStr := "<nil>"
			if stmt.RangeKey != nil {
				rangeKeyStr = stmt.RangeKey.Name
			}
			rangeValueStr := "<nil>"
			if stmt.RangeValue != nil {
				rangeValueStr = stmt.RangeValue.Name
			}
			fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt: Executing Range loop. Source: %s, Key: %s, Value: %s\n",
				stmt.RangeSource.String(),
				rangeKeyStr,
				rangeValueStr,
			)
		}
		sourceValRaw := stmt.RangeSource.Accept(i)
		if err, isErr := sourceValRaw.(RuntimeError); isErr {
			return err
		}
		if sig, isRet := IsReturnSignal(sourceValRaw); isRet {
			return sig // Propagate if expression somehow returned a signal
		}
		sourceVal, ok := sourceValRaw.(Value)
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("range source expression '%s' did not evaluate to a Value, got %T", stmt.RangeSource.String(), sourceValRaw)}
		}

		switch iterable := sourceVal.(type) {
		case *ArrayValue:
			for idx, elem := range iterable.Elements {
				loopEnv := NewEnvironment(i.environment) // New scope for each iteration

				if stmt.RangeKey != nil && stmt.RangeKey.Name != "_" {
					if i.DebugMode {
						fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt (Range): Defining key '%s' = %d\n", stmt.RangeKey.Name, idx)
					}
					loopEnv.Define(stmt.RangeKey.Name, NewIntegerValue(int64(idx))) // true for immutable
				}

				if stmt.RangeValue != nil && stmt.RangeValue.Name != "_" {
					if i.DebugMode {
						fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt (Range): Defining value '%s' = %s\n", stmt.RangeValue.Name, elem.String())
					}
					loopEnv.Define(stmt.RangeValue.Name, elem) // true for immutable
				}

				originalEnv := i.environment
				i.environment = loopEnv
				bodyResult := stmt.Body.Accept(i)
				i.environment = originalEnv // Restore environment

				if IsBreakSignal(bodyResult) {
					if i.DebugMode {
						fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt (Range): Break signal received.\n")
					}
					break // Break from Go's range loop
				}
				if IsContinueSignal(bodyResult) {
					if i.DebugMode {
						fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt (Range): Continue signal received.\n")
					}
					continue // Continue to next iteration of Go's range loop
				}
				if retVal, isRet := IsReturnSignal(bodyResult); isRet {
					if i.DebugMode {
						fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt (Range): Return signal received.\n")
					}
					return retVal // Propagate return signal
				}
				if err, isErr := bodyResult.(RuntimeError); isErr {
					if i.DebugMode {
						fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt (Range): Runtime error: %s\n", err.Error())
					}
					return err // Propagate runtime error
				}
			}
			return None // Normal completion of range loop over array
		// TODO: Add case for *StringValue (iterate over characters)
		// TODO: Add case for *MapValue (iterate over key-value pairs)
		default:
			return RuntimeError{Message: fmt.Sprintf("cannot range over type %s (value: %s)", sourceVal.Type(), sourceVal.String())}
		}
	}

	// Case 1: Loop with Init or Post (typically C-style for loop)
	// These loops get their own environment for variables declared in Init.
	if stmt.Init != nil || stmt.Post != nil {
		if i.DebugMode {
			fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt: Executing C-style/Init/Post loop.\n")
		}
		originalEnv := i.environment
		loopEnv := NewEnvironment(originalEnv) // Create new scope for loop
		i.environment = loopEnv
		defer func() { i.environment = originalEnv }() // Restore environment after loop

		if stmt.Init != nil {
			initResult := stmt.Init.Accept(i)
			if _, isRet := IsReturnSignal(initResult); isRet {
				return initResult
			}
			if err, isErr := initResult.(RuntimeError); isErr {
				return err
			}
		}

		for { // Infinite loop construct, condition checked inside
			conditionHolds := true // Default for loops without a condition (e.g. for ;; {})
			if stmt.Condition != nil {
				condValRaw := stmt.Condition.Accept(i)
				if _, isRet := IsReturnSignal(condValRaw); isRet {
					return condValRaw
				}
				if err, isErr := condValRaw.(RuntimeError); isErr {
					return err
				}

				condVal, ok := condValRaw.(Value)
				if !ok {
					return RuntimeError{Message: fmt.Sprintf("loop condition evaluated to non-Value type %T", condValRaw)}
				}
				conditionHolds = i.isTruthy(condVal)
			}

			if !conditionHolds {
				break // Exit C-style for loop
			}

			bodyResult := stmt.Body.Accept(i)

			if IsBreakSignal(bodyResult) {
				break
			}
			if IsContinueSignal(bodyResult) {
				if stmt.Post != nil { // Execute Post before continuing
					postResult := stmt.Post.Accept(i)
					if _, isRet := IsReturnSignal(postResult); isRet {
						return postResult
					}
					if err, isErr := postResult.(RuntimeError); isErr {
						return err
					}
				}
				continue
			}
			if retVal, isRet := IsReturnSignal(bodyResult); isRet {
				return retVal
			}
			if err, isErr := bodyResult.(RuntimeError); isErr {
				return err
			}

			if stmt.Post != nil {
				postResult := stmt.Post.Accept(i)
				if _, isRet := IsReturnSignal(postResult); isRet {
					return postResult
				}
				if err, isErr := postResult.(RuntimeError); isErr {
					return err
				}
			}
		}
		return None // Normal completion of C-style loop
	}

	// Case 2: Condition-only loop or Infinite loop (no Init, no Post)
	// These loops operate in the current environment.
	if i.DebugMode {
		fmt.Fprintf(os.Stderr, "Interpreter.VisitForStmt: Executing Condition-only or Infinite loop.\n")
	}
	for { // Infinite loop construct, condition checked inside
		conditionHolds := true     // Default for infinite loop (for {})
		if stmt.Condition != nil { // Condition-only loop
			condValRaw := stmt.Condition.Accept(i)
			if _, isRet := IsReturnSignal(condValRaw); isRet {
				return condValRaw
			}
			if err, isErr := condValRaw.(RuntimeError); isErr {
				return err
			}

			condVal, ok := condValRaw.(Value)
			if !ok {
				return RuntimeError{Message: fmt.Sprintf("loop condition evaluated to non-Value type %T", condValRaw)}
			}
			conditionHolds = i.isTruthy(condVal)
		}
		// If stmt.Condition is nil here, it's an infinite loop (conditionHolds remains true)

		if !conditionHolds {
			break // Exit condition-only or infinite loop
		}

		bodyResult := stmt.Body.Accept(i)

		if IsBreakSignal(bodyResult) {
			break
		}
		if IsContinueSignal(bodyResult) {
			continue
		} // No Post to worry about here
		if retVal, isRet := IsReturnSignal(bodyResult); isRet {
			return retVal
		}
		if err, isErr := bodyResult.(RuntimeError); isErr {
			return err
		}
	}
	return None // Normal completion of this type of loop
}

func (i *Interpreter) VisitReturnStmt(node *ast.ReturnStmt) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting ReturnStmt")
	}
	var returnValue Value = None // Default to None if no expression. Use singleton.
	if node.Value != nil {
		val, err := i.evaluate(node.Value)
		if err != nil {
			return err // Propagate evaluation error
		}
		returnValue = val
	}
	return ReturnSignal{Value: returnValue} // Return ReturnSignal struct
}

func (i *Interpreter) VisitBreakStmt(node *ast.BreakStmt) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting BreakStmt")
	}
	return BreakSignalValue
}

func (i *Interpreter) VisitContinueStmt(node *ast.ContinueStmt) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting ContinueStmt")
	}
	return ContinueSignalValue
}

// VisitIncDecStmt handles increment (++) and decrement (--) statements.
func (i *Interpreter) VisitIncDecStmt(stmt *ast.IncDecStmt) interface{} {
	if i.DebugMode {
		fmt.Printf("Visiting IncDecStmt: %s %s\n", stmt.LValue.String(), stmt.Operator)
	}

	// The LValue in an IncDecStmt must be an identifier based on current grammar.
	identExpr, ok := stmt.LValue.(*ast.IdentifierExpr)
	if !ok {
		// TODO: Attempt to get line/column information from stmt if available in AST node
		return RuntimeError{Message: fmt.Sprintf("invalid LValue for increment/decrement: expected identifier, got %T", stmt.LValue)}
	}

	varName := identExpr.Name

	// Get the current value of the variable
	currentVal, err := i.environment.Get(varName)
	if err != nil {
		// Variable not found or other error from Get
		// TODO: Attempt to get line/column information from stmt if available
		return RuntimeError{Message: err.Error()} // err from environment.Get should be descriptive
	}

	// Ensure the variable is an integer
	intVal, ok := currentVal.(*IntegerValue)
	if !ok {
		// TODO: Attempt to get line/column information from stmt if available
		return RuntimeError{Message: fmt.Sprintf("cannot %s non-integer type %s for variable '%s'", stmt.Operator, getTypeName(currentVal), varName)}
	}

	// Perform the operation
	newValue := intVal.Value
	switch stmt.Operator {
	case "++":
		newValue++
	case "--":
		newValue--
	default:
		// This case should ideally not be reached if the parser and AST builder are correct.
		// TODO: Attempt to get line/column information from stmt if available
		return RuntimeError{Message: fmt.Sprintf("unknown operator for IncDecStmt: %s", stmt.Operator)}
	}

	// Update the variable in the environment
	err = i.environment.Assign(varName, NewIntegerValue(newValue))
	if err != nil {
		// Error during assignment (e.g., assigning to a constant, though not typical for inc/dec)
		// TODO: Attempt to get line/column information from stmt if available
		return RuntimeError{Message: err.Error()} // err from environment.Assign should be descriptive
	}

	return nil // IncDec is a statement, does not produce a value for expression contexts
}

// Expressions
func (i *Interpreter) VisitUnaryExpr(node *ast.UnaryExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting UnaryExpr: %+v\n", *node)
	}

	// Handle address-of operator specially, as it does not evaluate its operand.
	if node.Operator == "&" {
		ident, ok := node.Right.(*ast.IdentifierExpr)
		if !ok {
			return RuntimeError{Message: "lvalue required as unary '&' operand"}
		}
		// Create a pointer that refers to the variable by name in its environment.
		return &PointerValue{
			VariableName:  ident.Name,
			ReferencedEnv: i.environment,
		}
	}

	rightVal, err := i.evaluate(node.Right)
	if err != nil {
		return err
	}

	switch node.Operator {
	case "!":
		return NewBoolValue(!i.isTruthy(rightVal))
	case "-":
		if num, ok := rightVal.(*IntegerValue); ok {
			return NewIntegerValue(-num.Value)
		}
		if num, ok := rightVal.(*FloatValue); ok {
			return NewFloatValue(-num.Value)
		}
		return RuntimeError{Message: "Operand for '-' must be a number."}
	case "*":
		ptr, ok := rightVal.(*PointerValue)
		if !ok {
			return RuntimeError{Message: "operand for '*' must be a pointer"}
		}
		val, err := ptr.Dereference()
		if err != nil {
			return RuntimeError{Message: err.Error()}
		}
		return val
	default:
		return RuntimeError{Message: fmt.Sprintf("Unknown unary operator '%s'.", node.Operator)}
	}
}

func (i *Interpreter) VisitBinaryExpr(node *ast.BinaryExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting BinaryExpr: %s\n", node.String())
	}

	leftVal, errLeft := i.evaluate(node.Left)
	if errLeft != nil {
		return errLeft
	}
	rightVal, errRight := i.evaluate(node.Right)
	if errRight != nil {
		return errRight
	}

	switch node.Operator {
	case "+":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)
		lStr, lIsStr := leftVal.(*StringValue)
		rStr, rIsStr := rightVal.(*StringValue)

		if lIsInt && rIsInt { // int + int
			return NewIntegerValue(lInt.Value + rInt.Value)
		} else if lIsFloat && rIsFloat { // float + float
			return NewFloatValue(lFloat.Value + rFloat.Value)
		} else if lIsInt && rIsFloat { // int + float
			return NewFloatValue(float64(lInt.Value) + rFloat.Value)
		} else if lIsFloat && rIsInt { // float + int
			return NewFloatValue(lFloat.Value + float64(rInt.Value))
		} else if lIsStr && rIsStr { // string + string
			return NewString(lStr.Value + rStr.Value)
		} else if lIsStr { // string + non-string
			return NewString(lStr.Value + rightVal.String()) // Convert right to string
		} else if rIsStr { // non-string + string
			return NewString(leftVal.String() + rStr.Value) // Convert left to string
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '+' must be numbers (integers or floats) or two strings. Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}

	case "-":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)

		if lIsInt && rIsInt { // int - int
			return NewIntegerValue(lInt.Value - rInt.Value)
		} else if lIsFloat && rIsFloat { // float - float
			return NewFloatValue(lFloat.Value - rFloat.Value)
		} else if lIsInt && rIsFloat { // int - float
			return NewFloatValue(float64(lInt.Value) - rFloat.Value)
		} else if lIsFloat && rIsInt { // float - int
			return NewFloatValue(lFloat.Value - float64(rInt.Value))
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '-' must be numbers (integers or floats). Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}

	case "*":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)

		if lIsInt && rIsInt { // int * int
			return NewIntegerValue(lInt.Value * rInt.Value)
		} else if lIsFloat && rIsFloat { // float * float
			return NewFloatValue(lFloat.Value * rFloat.Value)
		} else if lIsInt && rIsFloat { // int * float
			return NewFloatValue(float64(lInt.Value) * rFloat.Value)
		} else if lIsFloat && rIsInt { // float * int
			return NewFloatValue(lFloat.Value * float64(rInt.Value))
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '*' must be numbers (integers or floats). Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}

	case "/":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)

		// Check for division by zero
		isDivisorZero := false
		if rIsInt {
			if rInt.Value == 0 {
				isDivisorZero = true
			}
		} else if rIsFloat {
			if rFloat.Value == 0.0 {
				isDivisorZero = true
			}
		} // Note: if right operand is not int or float, it's an invalid type for division anyway.

		if isDivisorZero {
			return RuntimeError{
				Message: "Division by zero.",
				Line:    node.Line,
				Column:  node.Column,
			}
		}

		if lIsInt && rIsInt { // int / int -> result is int
			return NewIntegerValue(lInt.Value / rInt.Value)
		} else if lIsFloat && rIsFloat { // float / float
			return NewFloatValue(lFloat.Value / rFloat.Value)
		} else if lIsInt && rIsFloat { // int / float
			return NewFloatValue(float64(lInt.Value) / rFloat.Value)
		} else if lIsFloat && rIsInt { // float / int
			return NewFloatValue(lFloat.Value / float64(rInt.Value))
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '/' must be numbers (integers or floats). Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}

	case "%":
		if lInt, okL := leftVal.(*IntegerValue); okL {
			if rInt, okR := rightVal.(*IntegerValue); okR {
				if rInt.Value == 0 {
					return RuntimeError{
						Message: "Modulo by zero.",
						Line:    node.Line,
						Column:  node.Column,
					}
				}
				return NewIntegerValue(lInt.Value % rInt.Value)
			}
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '%%' must be integers. Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}

	case "==":
		return NewBoolValue(i.isEqual(leftVal, rightVal))
	case "!=":
		return NewBoolValue(!i.isEqual(leftVal, rightVal))

	case "<":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)
		lChar, lIsChar := leftVal.(*CharValue)
		rChar, rIsChar := rightVal.(*CharValue)

		if lIsInt && rIsInt { // int < int
			return NewBoolValue(lInt.Value < rInt.Value)
		} else if lIsFloat && rIsFloat { // float < float
			return NewBoolValue(lFloat.Value < rFloat.Value)
		} else if lIsInt && rIsFloat { // int < float
			return NewBoolValue(float64(lInt.Value) < rFloat.Value)
		} else if lIsFloat && rIsInt { // float < int
			return NewBoolValue(lFloat.Value < float64(rInt.Value))
		} else if lIsChar && rIsChar { // char < char
			return NewBoolValue(lChar.Value < rChar.Value)
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '<' must be numbers (integers or floats) or two characters. Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}
	case "<=":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)
		lChar, lIsChar := leftVal.(*CharValue)
		rChar, rIsChar := rightVal.(*CharValue)

		if lIsInt && rIsInt { // int <= int
			return NewBoolValue(lInt.Value <= rInt.Value)
		} else if lIsFloat && rIsFloat { // float <= float
			return NewBoolValue(lFloat.Value <= rFloat.Value)
		} else if lIsInt && rIsFloat { // int <= float
			return NewBoolValue(float64(lInt.Value) <= rFloat.Value)
		} else if lIsFloat && rIsInt { // float <= int
			return NewBoolValue(lFloat.Value <= float64(rInt.Value))
		} else if lIsChar && rIsChar { // char <= char
			return NewBoolValue(lChar.Value <= rChar.Value)
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '<=' must be numbers (integers or floats) or two characters. Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}
	case ">":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)
		lChar, lIsChar := leftVal.(*CharValue)
		rChar, rIsChar := rightVal.(*CharValue)

		if lIsInt && rIsInt { // int > int
			return NewBoolValue(lInt.Value > rInt.Value)
		} else if lIsFloat && rIsFloat { // float > float
			return NewBoolValue(lFloat.Value > rFloat.Value)
		} else if lIsInt && rIsFloat { // int > float
			return NewBoolValue(float64(lInt.Value) > rFloat.Value)
		} else if lIsFloat && rIsInt { // float > int
			return NewBoolValue(lFloat.Value > float64(rInt.Value))
		} else if lIsChar && rIsChar { // char > char
			return NewBoolValue(lChar.Value > rChar.Value)
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '>' must be numbers (integers or floats) or two characters. Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}
	case ">=":
		lInt, lIsInt := leftVal.(*IntegerValue)
		rInt, rIsInt := rightVal.(*IntegerValue)
		lFloat, lIsFloat := leftVal.(*FloatValue)
		rFloat, rIsFloat := rightVal.(*FloatValue)
		lChar, lIsChar := leftVal.(*CharValue)
		rChar, rIsChar := rightVal.(*CharValue)

		if lIsInt && rIsInt { // int >= int
			return NewBoolValue(lInt.Value >= rInt.Value)
		} else if lIsFloat && rIsFloat { // float >= float
			return NewBoolValue(lFloat.Value >= rFloat.Value)
		} else if lIsInt && rIsFloat { // int >= float
			return NewBoolValue(float64(lInt.Value) >= rFloat.Value)
		} else if lIsFloat && rIsInt { // float >= int
			return NewBoolValue(lFloat.Value >= float64(rInt.Value))
		} else if lIsChar && rIsChar { // char >= char
			return NewBoolValue(lChar.Value >= rChar.Value)
		}
		return RuntimeError{
			Message: fmt.Sprintf("Operands for '>=' must be numbers (integers or floats) or two characters. Got %s and %s.", leftVal.Type(), rightVal.Type()),
			Line:    node.Line,
			Column:  node.Column,
		}

	case "&&":
		lBool, okL := leftVal.(*BoolValue)
		if !okL {
			return RuntimeError{
				Message: fmt.Sprintf("Left operand for '&&' must be a boolean, got %s.", leftVal.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		rBool, okR := rightVal.(*BoolValue)
		if !okR {
			return RuntimeError{
				Message: fmt.Sprintf("Right operand for '&&' must be a boolean, got %s.", rightVal.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		return NewBoolValue(lBool.Value && rBool.Value)

	case "||":
		lBool, okL := leftVal.(*BoolValue)
		if !okL {
			return RuntimeError{
				Message: fmt.Sprintf("Left operand for '||' must be a boolean, got %s.", leftVal.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		rBool, okR := rightVal.(*BoolValue)
		if !okR {
			return RuntimeError{
				Message: fmt.Sprintf("Right operand for '||' must be a boolean, got %s.", rightVal.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		return NewBoolValue(lBool.Value || rBool.Value)

	default:
		return RuntimeError{
			Message: fmt.Sprintf("Unknown binary operator '%s'.", node.Operator),
			Line:    node.Line,
			Column:  node.Column,
		}
	}
}

func (i *Interpreter) VisitParenExpr(node *ast.ParenExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting ParenExpr: %s\n", node.String())
	}
	val, err := i.evaluate(node.Expression)
	if err != nil {
		return err // err is already a RuntimeError or similar error interface{}
	}
	return val // val is a Value interface{}
}

func (i *Interpreter) VisitIdentifierExpr(node *ast.IdentifierExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting IdentifierExpr: %s\n", node.Name)
	}
	value, err := i.environment.Get(node.Name)
	if err != nil {
		// err from environment.Get is fmt.Errorf("undefined variable '%s'", name)
		return RuntimeError{
			Message: err.Error(),
			Line:    node.Line,
			Column:  node.Column,
			// Line/column information is hard to get here accurately without token info
		}
	}
	return value
}

func (i *Interpreter) VisitTypeConversionExpr(node *ast.TypeConversionExpr) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting TypeConversionExpr")
	}
	return nil
}

func (i *Interpreter) VisitCompositeLiteralExpr(node *ast.CompositeLiteralExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting CompositeLiteralExpr: %s\n", node.String())
	}

	var literalType ast.TypeNode = node.Type
	if literalType == nil && len(i.expectedCompositeLiteralTypeStack) > 0 {
		literalType = i.expectedCompositeLiteralTypeStack[len(i.expectedCompositeLiteralTypeStack)-1]
		if i.DebugMode {
			fmt.Printf("  Inferred composite literal type from context: %s\n", literalType.String())
		}
	}

	if literalType == nil {
		return RuntimeError{Message: fmt.Sprintf("cannot infer type for composite literal: %s", node.String())}
	}

	switch typ := literalType.(type) {
	case *ast.TypeName: // Struct instantiation
		typeName := typ.Name
		structDef, exists := i.structs[typeName]
		if !exists {
			return RuntimeError{Message: fmt.Sprintf("undefined struct type '%s'", typeName)}
		}

		instance := NewStruct(typeName)
		for _, element := range node.Elements {
			keyExpr, ok := element.Key.(*ast.IdentifierExpr)
			if !ok {
				return RuntimeError{Message: "struct literal field names must be identifiers"}
			}
			fieldName := keyExpr.Name

			if _, fieldExists := structDef.Fields[fieldName]; !fieldExists {
				return RuntimeError{Message: fmt.Sprintf("unknown field '%s' in struct literal of type '%s'", fieldName, typeName)}
			}

			fieldValue, err := i.evaluate(element.Value)
			if err != nil {
				return err
			}
			instance.Fields[fieldName] = fieldValue
		}
		if i.DebugMode {
			fmt.Printf("  Instantiated struct %s with fields: %v\n", typeName, instance.Fields)
		}
		return instance

	case *ast.SliceTypeNode: // Slice literal
		if i.DebugMode {
			fmt.Printf("  Composite literal is a slice of type %s\n", typ.ElementType.String())
		}
		i.expectedCompositeLiteralTypeStack = append(i.expectedCompositeLiteralTypeStack, typ.ElementType)
		defer func() {
			i.expectedCompositeLiteralTypeStack = i.expectedCompositeLiteralTypeStack[:len(i.expectedCompositeLiteralTypeStack)-1]
		}()

		elements := make([]Value, len(node.Elements))
		for idx, compElement := range node.Elements {
			if compElement.Key != nil {
				return RuntimeError{Message: fmt.Sprintf("slice literal cannot have keyed elements: found key '%s'", compElement.Key.String())}
			}
			val, err := i.evaluate(compElement.Value)
			if err != nil {
				return err
			}
			elements[idx] = val
		}
		return NewArray(elements)

	default:
		return RuntimeError{Message: fmt.Sprintf("unhandled composite literal type: %T", typ)}
	}
}

func (i *Interpreter) VisitCompositeElement(node *ast.CompositeElement) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting CompositeElement")
	}
	return nil
}

func (i *Interpreter) VisitIndexAccessExpr(node *ast.IndexAccessExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting IndexAccessExpr: %s\n", node.String())
	}

	receiverVal := node.Receiver.Accept(i)
	if err, ok := receiverVal.(RuntimeError); ok {
		return err
	}
	// Propagate other control signals if necessary (though less common for sub-expressions)
	if IsBreakSignal(receiverVal) || IsContinueSignal(receiverVal) {
		return RuntimeError{
			Message: "unexpected control signal in index receiver expression",
			Line:    node.Line,
			Column:  node.Column,
		}
	}
	if _, ok := IsReturnSignal(receiverVal); ok {
		return RuntimeError{
			Message: "unexpected return signal in index receiver expression",
			Line:    node.Line,
			Column:  node.Column,
		}
	}

	indexVal := node.Index.Accept(i)
	if err, ok := indexVal.(RuntimeError); ok {
		return err
	}
	if IsBreakSignal(indexVal) || IsContinueSignal(indexVal) {
		return RuntimeError{
			Message: "unexpected control signal in index expression",
			Line:    node.Line,
			Column:  node.Column,
		}
	}
	if _, ok := IsReturnSignal(indexVal); ok {
		return RuntimeError{
			Message: "unexpected return signal in index expression",
			Line:    node.Line,
			Column:  node.Column,
		}
	}

	receiverValue, ok := receiverVal.(Value)
	if !ok {
		return RuntimeError{
			Message: fmt.Sprintf("receiver expression for index access did not evaluate to a Value, got %T", receiverVal),
			Line:    node.Line,
			Column:  node.Column,
		}
	}
	indexValue, ok := indexVal.(Value)
	if !ok {
		return RuntimeError{
			Message: fmt.Sprintf("index expression for index access did not evaluate to a Value, got %T", indexVal),
			Line:    node.Line,
			Column:  node.Column,
		}
	}

	switch collection := receiverValue.(type) {
	case *ArrayValue:
		intIndex, ok := indexValue.(*IntegerValue)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("array index must be an integer, got %s", indexValue.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		idx := int(intIndex.Value)
		if idx < 0 || idx >= len(collection.Elements) {
			return RuntimeError{
				Message: fmt.Sprintf("array index out of bounds: %d with length %d", idx, len(collection.Elements)),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		if i.DebugMode {
			fmt.Printf("  IndexAccessExpr: Accessing array element at index %d. Value: %s\n", idx, collection.Elements[idx].String())
		}
		return collection.Elements[idx]
	case *StringValue:
		intIndex, ok := indexValue.(*IntegerValue)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("string index must be an integer, got %s", indexValue.Type()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		idx := int(intIndex.Value)
		// V string indexing is by byte.
		if idx < 0 || idx >= len(collection.Value) {
			return RuntimeError{
				Message: fmt.Sprintf("string index out of bounds: %d with length %d", idx, len(collection.Value)),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		// V's string indexing returns a byte (char).
		charVal := NewChar(rune(collection.Value[idx]))
		if i.DebugMode {
			fmt.Printf("  IndexAccessExpr: Accessing string char at byte index %d. Value: '%s'\n", idx, charVal.String())
		}
		return charVal
	default:
		// Position information might be less precise here
		return RuntimeError{
			Message: fmt.Sprintf("type %s does not support index access", collection.Type()),
			Line:    node.Line,
			Column:  node.Column,
		} // TODO: Use actual position
	}
}

func (i *Interpreter) VisitFieldAccessExpr(node *ast.FieldAccessExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting FieldAccessExpr: %s on %T\n", node.Field.Name, node.Receiver)
	}

	receiverVal := node.Receiver.Accept(i)
	if err, ok := receiverVal.(RuntimeError); ok {
		return err // Propagate error
	}
	if IsBreakSignal(receiverVal) || IsContinueSignal(receiverVal) {
		// Position information might be less precise here, could use node.Pos() if available
		return RuntimeError{
			Message: "unexpected control signal in receiver of field access expression",
			Line:    node.Line,
			Column:  node.Column} // TODO: Use actual position from node.Field when available
	}
	if retVal, ok := IsReturnSignal(receiverVal); ok {
		// If a return signal is encountered, it means an expression (like a call) returned early.
		// We should probably propagate this or decide if field access on a return value makes sense.
		// For now, let's treat it as an error or unwrap, depending on desired language semantics.
		// Propagating the return signal might be more consistent.
		return retVal
	}

	receiverValue, ok := receiverVal.(Value)
	if !ok {
		return RuntimeError{
			Message: fmt.Sprintf("receiver of field access must be a value, got %T", receiverVal),
			Line:    node.Line,
			Column:  node.Column,
		} // TODO: Use actual position from node.Field when available
	}

	// --- BEGIN MODIFICATION: Handle pointer dereferencing ---
	if ptrVal, ok := receiverValue.(*PointerValue); ok {
		if i.DebugMode {
			fmt.Printf("  FieldAccessExpr: Left side is a pointer: %s. Attempting to dereference.\n", ptrVal.String())
		}
		dereferencedVal, err := ptrVal.Dereference()
		if err != nil {
			// TODO: The Dereference method itself should ideally return a RuntimeError with line/col
			return RuntimeError{
				Message: fmt.Sprintf("cannot access field '%s': %s", node.Field.Name, err.Error()),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		if dereferencedVal == nil { // Should ideally be caught by Dereference() erroring
			return RuntimeError{
				Message: fmt.Sprintf("cannot access field '%s' on nil pointer (dereferenced to nil)", node.Field.Name),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		receiverValue = dereferencedVal // Use the dereferenced value
		if i.DebugMode {
			fmt.Printf("  FieldAccessExpr: Dereferenced to: %s (Type: %s)\n", receiverValue.String(), receiverValue.Type())
		}
	}
	// --- END MODIFICATION ---

	switch rval := receiverValue.(type) {
	case *StructValue:
		fieldName := node.Field.Name

		// 1. Check for a field
		if fieldValue, ok := rval.Fields[fieldName]; ok {
			if i.DebugMode {
				fmt.Printf("  FieldAccessExpr: Accessing field '%s' on struct '%s'. Value: %s\n", fieldName, rval.TypeName, fieldValue.String())
			}
			return fieldValue
		}

		// 2. Check for a method
		structDef, exists := i.structs[rval.TypeName]
		if !exists {
			return RuntimeError{
				Message: fmt.Sprintf("definition for struct type '%s' not found", rval.TypeName),
				Line:    node.Line,
				Column:  node.Column,
			} // TODO: Use actual position from node.Field when available
		}

		if method, ok := structDef.Methods[fieldName]; ok {
			boundMethod := NewVMethod(rval, method)
			if i.DebugMode {
				fmt.Printf("  FieldAccessExpr: Binding method '%s' to struct instance '%s'\n", fieldName, rval.TypeName)
			}
			return boundMethod
		}

		// 3. Neither field nor method found
		return RuntimeError{
			Message: fmt.Sprintf("undefined field or method '%s' on struct type '%s'", fieldName, rval.TypeName),
			Line:    node.Line,
			Column:  node.Column,
		} // TODO: Use actual position from node.Field when available

	default:
		return RuntimeError{
			Message: fmt.Sprintf("type %s does not support field access", receiverValue.Type()),
			Line:    node.Line,
			Column:  node.Column,
		} // TODO: Use actual position from node.Field when available
	}
}

// ... (rest of the code remains the same)
func (i *Interpreter) VisitCallExpr(node *ast.CallExpr) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting CallExpr")
	}

	// 1. Evaluate the function expression
	calleeVal := node.Function.Accept(i)
	if err, ok := calleeVal.(RuntimeError); ok {
		return err
	}

	// ADDED DEBUG PRINT
	if i.DebugMode {
		fmt.Printf("[DEBUG] VisitCallExpr: CalleeVal before switch. Go Type: %T, Value: %v\n", calleeVal, calleeVal)
		if cv, ok_cv := calleeVal.(Value); ok_cv {
			fmt.Printf("[DEBUG] VisitCallExpr: CalleeVal is Value. Runtime Type: %s\n", cv.Type().String())
		} else {
			fmt.Printf("[DEBUG] VisitCallExpr: CalleeVal is NOT an interpreter.Value. Go Type: %T\n", calleeVal)
		}
	}

	// 2. Evaluate arguments
	evaluatedArgs := make([]Value, len(node.Arguments))
	for idx, argExpr := range node.Arguments {
		argVal := argExpr.Accept(i)
		if err, ok := argVal.(RuntimeError); ok {
			return err // Propagate error from argument evaluation
		}
		val, ok := argVal.(Value)
		if !ok {
			return RuntimeError{
				Message: fmt.Sprintf("argument %d evaluation did not return a Value, got %T", idx+1, argVal),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		evaluatedArgs[idx] = val
	}

	// 3. Call the function based on its type
	switch callee := calleeVal.(type) {
	case *VFunction:
		if i.DebugMode {
			fmt.Printf("[DEBUG] VisitCallExpr: Matched *VFunction\n")
		}
		result, callErr := callee.Call(i, evaluatedArgs)
		if callErr != nil {
			if re, ok := callErr.(RuntimeError); ok {
				return re
			}
			return RuntimeError{
				Message: callErr.Error(),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		// If the result is a ReturnSignal, unwrap it to get the actual return value.
		if retVal, isRet := IsReturnSignal(result); isRet {
			return retVal
		}
		return result // Should be nil if no explicit return, or for functions not returning values (e.g. procedures)
	case *VMethod:
		if i.DebugMode {
			fmt.Printf("[DEBUG] VisitCallExpr: Matched *VMethod\n")
		}
		result, callErr := callee.Call(i, evaluatedArgs)
		if callErr != nil {
			if re, ok := callErr.(RuntimeError); ok {
				return re
			}
			// TODO: Consider if line/col info can be added here if not already in callErr
			return RuntimeError{
				Message: callErr.Error(),
				Line:    node.Line,
				Column:  node.Column,
			} // Line: node.Pos().Line, Column: node.Pos().Column (if available)
		}
		// If the result is a ReturnSignal, unwrap it to get the actual return value.
		// This is crucial for methods that explicitly return.
		if retVal, isRet := IsReturnSignal(result); isRet {
			return retVal
		}
		return result
	case *NativeFunction:
		if i.DebugMode {
			fmt.Printf("[DEBUG] VisitCallExpr: Matched *NativeFunction\n")
		}
		result, callErr := callee.Call(i, evaluatedArgs)
		if callErr != nil {
			if re, ok := callErr.(RuntimeError); ok {
				return re
			}
			return RuntimeError{
				Message: callErr.Error(),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
		return result
	default:
		if i.DebugMode {
			fmt.Printf("[DEBUG] VisitCallExpr: Default case in switch. Callee type from val.(type): %T, Value: %v\n", callee, callee)
		}
		var calleeTypeStr string
		if cVal, ok := calleeVal.(Value); ok {
			calleeTypeStr = cVal.Type().String()
		} else {
			calleeTypeStr = fmt.Sprintf("%T", calleeVal)
		}
		funcExprStr := "<unknown function expression>"
		if node.Function != nil {
			// Assuming ast.Node has a String() method for a textual representation
			funcExprStr = node.Function.String()
		}
		return RuntimeError{
			Message: fmt.Sprintf("cannot call non-function type %s (expression: %s)", calleeTypeStr, funcExprStr),
			Line:    node.Line,
			Column:  node.Column,
		}
	}
}

// VisitTypeOfExpr evaluates a TypeOf(expression) node.
func (i *Interpreter) VisitTypeOfExpr(node *ast.TypeOfExpr) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting TypeOfExpr: TypeOf(%s)\n", node.Expression.String())
	}

	// Evaluate the inner expression
	valueResult := node.Expression.Accept(i)

	// Check for runtime errors from expression evaluation
	if err, ok := valueResult.(RuntimeError); ok {
		return err // Propagate runtime error
	}
	// Check for control signals (though less common from simple expression evaluation, good practice)
	if IsBreakSignal(valueResult) || IsContinueSignal(valueResult) {
		// This shouldn't happen from just evaluating the argument of TypeOf,
		// but if it did, it would be an error in context.
		return RuntimeError{Message: "unexpected control flow signal in TypeOf argument evaluation"}
	}
	if _, isRet := IsReturnSignal(valueResult); isRet {
		// A return from within TypeOf's argument is unexpected.
		return RuntimeError{Message: "unexpected return signal in TypeOf argument evaluation"}
	}

	// Ensure we have a valid Value to work with
	val, ok := valueResult.(Value)
	if !ok {
		// This indicates a problem, e.g., the visitor for the expression returned something unexpected.
		return RuntimeError{Message: fmt.Sprintf("TypeOf argument evaluation did not return a valid Value, got %T", valueResult)}
	}

	// Use the getTypeName helper to get the string representation of the type
	typeName := getTypeName(val)

	if i.DebugMode {
		fmt.Printf("Interpreter.VisitTypeOfExpr: Evaluated expression to %v, type name: '%s'\n", val, typeName)
	}

	return NewString(typeName)
}

// Literals
func (i *Interpreter) VisitIntegerLiteral(node *ast.IntegerLiteral) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting IntegerLiteral: %s\n", node.Value)
	}
	val, err := strconv.ParseInt(node.Value, 10, 64)
	if err != nil {
		// TODO: Return a proper RuntimeError, perhaps using line/col info from the node if available
		// For now, we'll create a simple error. Line/column info would be good to add to AST nodes.
		return RuntimeError{
			Message: fmt.Sprintf("invalid integer literal '%s': %s", node.Value, err.Error()), /*, Line: node.Pos().Line, Column: node.Pos().Column*/
			Line:    node.Line,
			Column:  node.Column,
		}
	}
	return NewIntegerValue(val)
}

func (i *Interpreter) VisitFloatLiteral(node *ast.FloatLiteral) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting FloatLiteral: %s\n", node.Value)
	}
	val, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return RuntimeError{
			Message: fmt.Sprintf("invalid float literal '%s': %s", node.Value, err.Error()),
			Line:    node.Line,
			Column:  node.Column,
		}
	}
	return NewFloatValue(val)
}

func (i *Interpreter) VisitStringLiteral(node *ast.StringLiteral) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting StringLiteral for interpolation: %s\n", node.Value)
	}

	// Regex to find $variable (group 1) or ${expression} (group 2 for expression content)
	re := regexp.MustCompile(`\$(?:([a-zA-Z_][a-zA-Z0-9_]*)|\{([^}]*)\})`)

	interpolatedString := re.ReplaceAllStringFunc(node.Value, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 3 { // Should not happen: 0 is full, 1 is $var, 2 is ${expr_content}
			return match // Should not happen, return original placeholder
		}

		// Case 1: Simple variable interpolation, e.g., $foo
		if varName := submatches[1]; varName != "" {
			if i.DebugMode {
				fmt.Printf("Interpolation: Simple variable '%s' found.\n", varName)
			}
			val, err := i.environment.Get(varName)
			if err != nil {
				if i.DebugMode {
					fmt.Printf("Interpolation: Variable '%s' not found. Error: %v\n", varName, err)
				}
				return match // Variable not found, return placeholder
			}
			return val.String()
		}

		// Case 2: Expression interpolation, e.g., ${matrix[1][0]}
		if exprStr := submatches[2]; exprStr != "" {
			if i.DebugMode {
				fmt.Printf("Interpolation: Expression '%s' found.\n", exprStr)
			}
			if i.exprParser == nil {
				if i.DebugMode {
					fmt.Println("Interpolation: Expression parser is not configured. Cannot evaluate expression.")
				}
				return match // Parser not available, return placeholder
			}

			astNode, parseErr := i.exprParser(exprStr)
			if parseErr != nil {
				if i.DebugMode {
					fmt.Printf("Interpolation: Failed to parse expression '%s'. Error: %v\n", exprStr, parseErr)
				}
				return match // Parsing failed, return placeholder
			}

			evalResult := astNode.Accept(i)

			if rtErr, ok := evalResult.(RuntimeError); ok {
				if i.DebugMode {
					fmt.Printf("Interpolation: Runtime error evaluating expression '%s'. Error: %s\n", exprStr, rtErr.Error())
				}
				return match // Evaluation error, return placeholder
			}

			if val, ok := evalResult.(Value); ok {
				return val.String()
			}

			if i.DebugMode {
				fmt.Printf("Interpolation: Expression '%s' evaluated to unexpected type %T. Value: %+v\n", exprStr, evalResult, evalResult)
			}
			return match // Unexpected evaluation result type
		}

		// Should not be reached if regex is correct and one group matches
		return match
	})

	return NewString(interpolatedString)
}

func (i *Interpreter) VisitCharLiteral(node *ast.CharLiteral) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting CharLiteral: %s\n", node.Value)
	}
	// node.Value is the content of the char literal, e.g., "a" or "\n" (as a string representing the single rune).
	runes := []rune(node.Value)
	if len(runes) != 1 {
		return RuntimeError{
			Message: fmt.Sprintf("invalid char literal '%s': must be a single character", node.Value),
			Line:    node.Line,
			Column:  node.Column,
		}
	}
	return NewChar(runes[0])
}

func (i *Interpreter) VisitBoolLiteral(node *ast.BoolLiteral) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting BoolLiteral: %t\n", node.Value)
	}
	// Ensure NewBoolean returns a value whose Type() method returns BoolType
	return NewBoolValue(node.Value)
}

func (i *Interpreter) VisitNilLiteral(node *ast.NilLiteral) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting NilLiteral")
	}
	return None // Use the global None instance
}

func (i *Interpreter) VisitArrayLiteral(node *ast.ArrayLiteral) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting ArrayLiteral")
	}
	elements := make([]Value, 0, len(node.Elements))
	for _, exprNode := range node.Elements {
		evaluatedElement := exprNode.Accept(i)
		if err, ok := evaluatedElement.(RuntimeError); ok {
			return err
		}
		if val, ok := evaluatedElement.(Value); ok {
			elements = append(elements, val)
		} else {
			return RuntimeError{
				Message: fmt.Sprintf("unexpected type %T from expression in array literal", evaluatedElement),
				Line:    node.Line,
				Column:  node.Column,
			}
		}
	}
	return NewArray(elements)
}

func (i *Interpreter) VisitMapLiteral(node *ast.MapLiteral) interface{} {
	if i.DebugMode {
		fmt.Fprintf(os.Stderr, "Interpreter.Visiting MapLiteral: %s\n", node.String())
	}

	mapInstance := NewMap() // NewMap creates a MapValue with an initialized Elements map

	for _, entryNode := range node.Entries {
		entryResultRaw := entryNode.Accept(i)

		if err, ok := entryResultRaw.(RuntimeError); ok {
			return err // Propagate error from VisitMapEntry
		}

		entryPair, ok := entryResultRaw.([]Value)
		if !ok || len(entryPair) != 2 {
			return RuntimeError{
				Message: fmt.Sprintf("internal error: map entry evaluation did not return a key-value pair, got %T", entryResultRaw),
				Line:    node.Line,
				Column:  node.Column,
			}
		}

		key := entryPair[0]
		value := entryPair[1]

		stringKey, ok := key.(*StringValue)
		if !ok {
			// Attempt to convert other simple types to string if desired, or enforce string keys strictly.
			// For now, strict string keys.
			return RuntimeError{
				Message: fmt.Sprintf("map keys must be strings, got %s (%T)", key.String(), key),
				Line:    node.Line,
				Column:  node.Column,
			}
		}

		mapInstance.Elements[stringKey.Value] = value
	}

	return mapInstance
}

func (i *Interpreter) VisitMapEntry(node *ast.MapEntry) interface{} {
	if i.DebugMode {
		fmt.Fprintf(os.Stderr, "Interpreter.Visiting MapEntry: Key=%s, Value=%s\n", node.Key.String(), node.Value.String())
	}

	keyResult := node.Key.Accept(i)
	if err, ok := keyResult.(RuntimeError); ok {
		return err
	}
	// Ensure keyResult is a Value, not some other interface type from Accept
	keyValue, ok := keyResult.(Value)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("map entry key evaluation did not return a Value, got %T", keyResult)}
	}

	valueResult := node.Value.Accept(i)
	if err, ok := valueResult.(RuntimeError); ok {
		return err
	}
	// Ensure valueResult is a Value
	valueValue, ok := valueResult.(Value)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("map entry value evaluation did not return a Value, got %T", valueResult)}
	}

	return []Value{keyValue, valueValue}
}

// Helper for equality comparison
func (i *Interpreter) isEqual(a, b Value) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Handle cross-type equality for numeric types (int and float)
	if a.Type() != b.Type() {
		if valAInt, okA := a.(*IntegerValue); okA {
			if valBFloat, okB := b.(*FloatValue); okB {
				return float64(valAInt.Value) == valBFloat.Value
			}
		}
		if valAFloat, okA := a.(*FloatValue); okA {
			if valBInt, okB := b.(*IntegerValue); okB {
				return valAFloat.Value == float64(valBInt.Value)
			}
		}
		// If types are different and not a handled numeric cross-type comparison, they are not equal.
		return false
	}

	// Types are the same, proceed with value comparison
	switch valA := a.(type) {
	case *IntegerValue:
		return valA.Value == b.(*IntegerValue).Value
	case *FloatValue: // Ensure this is uncommented and correct
		valB, okB := b.(*FloatValue)
		if !okB { // Should not happen if types were already checked as same, but good for safety
			return false
		}
		return valA.Value == valB.Value
	case *BoolValue:
		return valA.Value == b.(*BoolValue).Value
	case *StringValue:
		return valA.Value == b.(*StringValue).Value
	case *CharValue:
		return valA.Value == b.(*CharValue).Value
	case *NoneValue:
		_, okB := b.(*NoneValue)
		return okB
	// TODO: Add comparisons for ArrayValue, MapValue, StructValue (reference or deep equality based on language spec)
	default:
		// For unhandled types, consider them not equal or panic if this state is unexpected.
		// V might compare by reference for complex types if not overloaded.
		return false
	}
}

// Types (Interpreter.Visiting type nodes usually means returning type information, not executing anything)
func (i *Interpreter) VisitTypeName(node *ast.TypeName) interface{} {
	if i.DebugMode {
		fmt.Fprintf(os.Stderr, "Interpreter.Visiting TypeName: %s\n", node.Name)
	}
	// TypeName node itself represents type information.
	// Return the node, or a more structured type representation if needed later.
	return node
}

func (i *Interpreter) VisitPrimitiveTypeNode(node *ast.PrimitiveTypeNode) interface{} {
	if i.DebugMode {
		fmt.Fprintf(os.Stderr, "Interpreter.Visiting PrimitiveTypeNode: %s\n", node.Kind.String())
	}
	// For type nodes, we often return the node itself or a representation of the type.
	// Here, returning the node is consistent with VisitTypeName, etc.
	return node
}

func (i *Interpreter) VisitArrayTypeNode(node *ast.ArrayTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting ArrayTypeNode")
	}
	elementTypeVal := node.ElementType.Accept(i)
	if err, ok := elementTypeVal.(RuntimeError); ok {
		return err
	}
	elementTv, ok := elementTypeVal.(*TypeValue)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for array element type, got %T", elementTypeVal)}
	}

	sizeStr := ""
	if node.Size != nil { // Size is ast.Expression
		// For simplicity, we expect a literal integer for size in this context.
		// A full implementation might evaluate constant expressions.
		if sizeLit, ok := node.Size.(*ast.IntegerLiteral); ok {
			sizeStr = sizeLit.Value
		} else {
			// If size is not an integer literal, creating a type string is complex.
			// For now, use a placeholder or return an error.
			return RuntimeError{Message: fmt.Sprintf("array size expression must be an integer literal, got %T", node.Size)}
		}
	} else {
		// This case should ideally be handled by SliceTypeNode or indicate an error if ArrayTypeNode expects a size.
		// However, grammar `arrayType: '[' expression? ']' typeName;` expression is optional.
		// If expression is nil, it might mean a slice in some V-like syntaxes, but we have SliceTypeNode.
		// Let's assume for ArrayTypeNode, size must be present.
		return RuntimeError{Message: "array type declaration is missing size"}
	}

	return NewTypeValue(fmt.Sprintf("[%s]%s", sizeStr, elementTv.Name))
}

func (i *Interpreter) VisitSliceTypeNode(node *ast.SliceTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting SliceTypeNode")
	}
	elementTypeVal := node.ElementType.Accept(i)
	if err, ok := elementTypeVal.(RuntimeError); ok {
		return err
	}
	elementTv, ok := elementTypeVal.(*TypeValue)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for slice element type, got %T", elementTypeVal)}
	}
	return NewTypeValue(fmt.Sprintf("[]%s", elementTv.Name))
}

func (i *Interpreter) VisitMapTypeNode(node *ast.MapTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting MapTypeNode")
	}
	keyTypeVal := node.KeyType.Accept(i)
	if err, ok := keyTypeVal.(RuntimeError); ok {
		return err
	}
	keyTv, ok := keyTypeVal.(*TypeValue)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for map key type, got %T", keyTypeVal)}
	}

	valueTypeVal := node.ValueType.Accept(i)
	if err, ok := valueTypeVal.(RuntimeError); ok {
		return err
	}
	valueTv, ok := valueTypeVal.(*TypeValue)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for map value type, got %T", valueTypeVal)}
	}
	return NewTypeValue(fmt.Sprintf("map[%s]%s", keyTv.Name, valueTv.Name))
}

func (i *Interpreter) VisitPointerTypeNode(node *ast.PointerTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting PointerTypeNode")
	}
	elementTypeVal := node.ElementType.Accept(i)
	if err, ok := elementTypeVal.(RuntimeError); ok {
		return err
	}
	elementTv, ok := elementTypeVal.(*TypeValue)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for pointer element type, got %T", elementTypeVal)}
	}
	return NewTypeValue(fmt.Sprintf("*%s", elementTv.Name))
}

func (i *Interpreter) VisitFunctionTypeNode(node *ast.FunctionTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting FunctionTypeNode")
	}
	paramTypeNames := make([]string, len(node.ParameterTypes))
	for idx, paramTypeNode := range node.ParameterTypes {
		paramTypeVal := paramTypeNode.Accept(i)
		if err, ok := paramTypeVal.(RuntimeError); ok {
			return err
		}
		paramTv, ok := paramTypeVal.(*TypeValue)
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for function parameter type, got %T", paramTypeVal)}
		}
		paramTypeNames[idx] = paramTv.Name
	}

	returnsStr := ""
	if node.ReturnType != nil {
		returnTypeVal := node.ReturnType.Accept(i)
		if err, ok := returnTypeVal.(RuntimeError); ok {
			return err
		}
		returnTv, ok := returnTypeVal.(*TypeValue)
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for function return type, got %T", returnTypeVal)}
		}
		// How to handle multiple return values if ReturnType is a single TypeNode (e.g. a TupleTypeNode)?
		// For now, assume single return type or it's already formatted in returnTv.Name if it's a tuple.
		// If your language supports explicit multiple return values like Go (e.g., fn() (int, string)),
		// then ast.FunctionTypeNode.ReturnType might be a *ast.TupleTypeNode or similar.
		// If it's a tuple, returnTv.Name should ideally be like "(int, string)".
		// For now, we just use its name. If it needs parentheses, the TypeValue.String() for that tuple type should handle it.
		returnsStr = returnTv.Name
	}

	paramsStr := strings.Join(paramTypeNames, ", ")

	if returnsStr == "" {
		return NewTypeValue(fmt.Sprintf("fn(%s)", paramsStr))
	} else {
		// If ReturnType is a single type and not a tuple, and you want Go-like syntax for multiple returns (fn() (int, error)),
		// this part might need adjustment based on how TypeValue.Name for tuples is formatted.
		// If returnTv.Name for a tuple type is already "(T1, T2)", then this is fine.
		// Otherwise, if it's just "T1, T2", you might need to add parentheses here.
		// Assuming TypeValue.Name for a tuple type is already parenthesized if needed.
		return NewTypeValue(fmt.Sprintf("fn(%s) %s", paramsStr, returnsStr))
	}
}

func (i *Interpreter) VisitOptionalTypeNode(node *ast.OptionalTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting OptionalTypeNode")
	}
	underlyingTypeVal := node.ElementType.Accept(i)
	if err, ok := underlyingTypeVal.(RuntimeError); ok {
		return err
	}
	underlyingTv, ok := underlyingTypeVal.(*TypeValue)
	if !ok {
		return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for optional underlying type, got %T", underlyingTypeVal)}
	}
	return NewTypeValue(fmt.Sprintf("?%s", underlyingTv.Name))
}

func (i *Interpreter) VisitAnonymousStructTypeNode(node *ast.AnonymousStructTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting AnonymousStructTypeNode")
	}
	fieldStrings := make([]string, len(node.Fields))
	for idx, field := range node.Fields {
		fieldTypeVal := field.Type.Accept(i)
		if err, ok := fieldTypeVal.(RuntimeError); ok {
			return err
		}
		fieldTv, ok := fieldTypeVal.(*TypeValue)
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for struct field type, got %T", fieldTypeVal)}
		}
		fieldStrings[idx] = fmt.Sprintf("%s %s", field.Name.Name, fieldTv.Name)
	}
	return NewTypeValue(fmt.Sprintf("struct { %s }", strings.Join(fieldStrings, "; ")))
}

func (i *Interpreter) VisitAnonymousInterfaceTypeNode(node *ast.AnonymousInterfaceTypeNode) interface{} {
	if i.DebugMode {
		fmt.Println("Interpreter.Visiting AnonymousInterfaceTypeNode")
	}
	// For now, visiting an interface type definition doesn't produce a runtime value.
	// This might change if types become first-class values.
	return nil
}

func (i *Interpreter) VisitMethodSignatureNode(node *ast.MethodSignatureNode) interface{} {
	if i.DebugMode {
		fmt.Printf("Interpreter.Visiting MethodSignatureNode: %s\n", node.Name.Name)
	}
	paramTypeNames := make([]string, len(node.Parameters))
	for idx, paramDecl := range node.Parameters { // paramDecl is *ast.ParameterDecl
		// Assuming ParameterDecl has a Type field of type ast.TypeNode
		paramTypeVal := paramDecl.Type.Accept(i)
		if err, ok := paramTypeVal.(RuntimeError); ok {
			return err
		}
		paramTv, ok := paramTypeVal.(*TypeValue)
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for method parameter type, got %T", paramTypeVal)}
		}
		paramTypeNames[idx] = paramTv.Name
	}

	returnsStr := ""
	if node.ReturnType != nil {
		returnTypeVal := node.ReturnType.Accept(i)
		if err, ok := returnTypeVal.(RuntimeError); ok {
			return err
		}
		returnTv, ok := returnTypeVal.(*TypeValue)
		if !ok {
			return RuntimeError{Message: fmt.Sprintf("expected *TypeValue for method return type, got %T", returnTypeVal)}
		}
		// Similar to FunctionTypeNode, assuming returnTv.Name is correctly formatted (e.g., parenthesized for tuples)
		returnsStr = returnTv.Name
	}

	paramsStr := strings.Join(paramTypeNames, ", ")

	var signature string
	if returnsStr == "" {
		signature = fmt.Sprintf("%s(%s)", node.Name.Name, paramsStr)
	} else {
		signature = fmt.Sprintf("%s(%s) %s", node.Name.Name, paramsStr, returnsStr)
	}
	return NewTypeValue(signature)
}

// GetGlobals retorna el entorno global del intérprete
func (i *Interpreter) GetGlobals() *Environment {
	return i.globals
}
