package interpreter

import (
	"fmt"
	"strings"

	"github.com/xvimnt/OLC2_PROYECTO2_G15/ast"
)

// ValueType represents the type of a value in the V language
type ValueType int

const (
	IntegerType ValueType = iota
	FloatType
	StringType
	BoolType // Added missing BoolType
	CharType
	ArrayType
	MapType
	FunctionType
	StructType
	NoneType
	OptionalType
	ResultType
	EnumType
	TypeValueType // Represents a type itself as a value
	MethodType    // Represents a bound method
	PointerType   // Represents a pointer to a variable
)

func (vt ValueType) String() string {
	switch vt {
	case IntegerType:
		return "integer"
	case FloatType:
		return "float"
	case StringType:
		return "string"
	case BoolType:
		return "bool"
	case CharType:
		return "char"
	case ArrayType:
		return "array"
	case MapType:
		return "map"
	case FunctionType:
		return "function"
	case StructType:
		return "struct"
	case NoneType:
		return "none"
	case OptionalType:
		return "optional"
	case ResultType:
		return "result"
	case EnumType:
		return "enum"
	case TypeValueType:
		return "type_value"
	case MethodType:
		return "method"
	case PointerType:
		return "pointer"
	default:
		return "unknown_type"
	}
}

// Value is the interface for all runtime values in the V language interpreter
type Value interface {
	Type() ValueType
	String() string
	Equals(other Value) bool
}

// IntegerValue represents an integer value in V
type IntegerValue struct {
	Value int64
}

func NewIntegerValue(value int64) *IntegerValue {
	return &IntegerValue{Value: value}
}

func (v *IntegerValue) Type() ValueType {
	return IntegerType
}

func (v *IntegerValue) String() string {
	return fmt.Sprintf("%d", v.Value)
}

func (v *IntegerValue) Equals(other Value) bool {
	if otherInt, ok := other.(*IntegerValue); ok {
		return v.Value == otherInt.Value
	}
	return false
}

// FloatValue represents a floating-point value in V
type FloatValue struct {
	Value float64
}

func NewFloatValue(value float64) *FloatValue {
	return &FloatValue{Value: value}
}

func (v *FloatValue) Type() ValueType {
	return FloatType
}

func (v *FloatValue) String() string {
	return fmt.Sprintf("%g", v.Value)
}

func (v *FloatValue) Equals(other Value) bool {
	if otherFloat, ok := other.(*FloatValue); ok {
		return v.Value == otherFloat.Value
	}
	return false
}

// BoolValue represents a boolean value in V
type BoolValue struct {
	Value bool
}

func NewBoolValue(value bool) *BoolValue {
	return &BoolValue{Value: value}
}

func (v *BoolValue) Type() ValueType {
	return BoolType
}

func (v *BoolValue) String() string {
	return fmt.Sprintf("%t", v.Value)
}

func (v *BoolValue) Equals(other Value) bool {
	if otherBool, ok := other.(*BoolValue); ok {
		return v.Value == otherBool.Value
	}
	return false
}

// StringValue represents a string value in V
type StringValue struct {
	Value string
}

func NewString(value string) *StringValue {
	return &StringValue{Value: value}
}

func (v *StringValue) Type() ValueType {
	return StringType
}

func (v *StringValue) String() string {
	return v.Value
}

func (v *StringValue) Equals(other Value) bool {
	if otherStr, ok := other.(*StringValue); ok {
		return v.Value == otherStr.Value
	}
	return false
}

// CharValue represents a character value in V
type CharValue struct {
	Value rune
}

func NewChar(value rune) *CharValue {
	return &CharValue{Value: value}
}

func (v *CharValue) Type() ValueType {
	return CharType
}

func (v *CharValue) String() string {
	return string(v.Value)
}

func (v *CharValue) Equals(other Value) bool {
	if otherChar, ok := other.(*CharValue); ok {
		return v.Value == otherChar.Value
	}
	return false
}

// ArrayValue represents an array value in V
type ArrayValue struct {
	Elements []Value
}

func NewArray(elements []Value) *ArrayValue {
	return &ArrayValue{Elements: elements}
}

func (v *ArrayValue) Type() ValueType {
	return ArrayType
}

func (v *ArrayValue) String() string {
	var elems []string
	for _, e := range v.Elements {
		elems = append(elems, e.String())
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

func (v *ArrayValue) Equals(other Value) bool {
	otherArray, ok := other.(*ArrayValue)
	if !ok || len(v.Elements) != len(otherArray.Elements) {
		return false
	}

	for i, elem := range v.Elements {
		if !elem.Equals(otherArray.Elements[i]) {
			return false
		}
	}
	return true
}

// MapValue represents a map value in V
type MapValue struct {
	Elements map[string]Value
}

func NewMap() *MapValue {
	return &MapValue{Elements: make(map[string]Value)}
}

func (v *MapValue) Type() ValueType {
	return MapType
}

func (v *MapValue) String() string {
	var pairs []string
	for k, val := range v.Elements {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k, val.String()))
	}
	return "{" + strings.Join(pairs, ", ") + "}"
}

func (v *MapValue) Equals(other Value) bool {
	otherMap, ok := other.(*MapValue)
	if !ok || len(v.Elements) != len(otherMap.Elements) {
		return false
	}

	for k, val := range v.Elements {
		otherVal, exists := otherMap.Elements[k]
		if !exists || !val.Equals(otherVal) {
			return false
		}
	}
	return true
}

// Function represents a callable function in V
type Function interface {
	Value
	Call(interpreter *Interpreter, arguments []Value) (Value, error)
}

// NativeFunction represents a Go function that can be called from V code
type NativeFunction struct {
	Name       string
	ParamNames []string
	Function   func(interpreter *Interpreter, arguments []Value) (Value, error)
}

func NewNativeFunction(name string, paramNames []string, function func(interpreter *Interpreter, arguments []Value) (Value, error)) *NativeFunction {
	return &NativeFunction{
		Name:       name,
		ParamNames: paramNames,
		Function:   function,
	}
}

func (f *NativeFunction) Type() ValueType {
	return FunctionType
}

func (f *NativeFunction) String() string {
	return fmt.Sprintf("<native fn %s>", f.Name)
}

func (f *NativeFunction) Equals(other Value) bool {
	if otherNat, ok := other.(*NativeFunction); ok {
		return f == otherNat
	}
	return false
}

func (f *NativeFunction) Call(interpreter *Interpreter, arguments []Value) (Value, error) {
	return f.Function(interpreter, arguments)
}

// VFunction represents a function defined in V code
type VFunction struct {
	Declaration *ast.FunctionDecl
	Closure     *Environment
	Name        string
	ParamNames  []string
}

func NewVFunction(declaration *ast.FunctionDecl, closure *Environment, name string, paramNames []string) *VFunction {
	return &VFunction{
		Declaration: declaration,
		Closure:     closure,
		Name:        name,
		ParamNames:  paramNames,
	}
}

func (f *VFunction) Type() ValueType {
	return FunctionType
}

func (f *VFunction) String() string {
	return fmt.Sprintf("<fn %s>", f.Name)
}

func (f *VFunction) Equals(other Value) bool {
	if otherVFunc, ok := other.(*VFunction); ok {
		return f == otherVFunc
	}
	return false
}

func (f *VFunction) Call(interpreter *Interpreter, arguments []Value) (Value, error) {
	// Create a new environment for the function call, with the function's closure as the enclosing environment.
	env := NewEnvironment(f.Closure)
	if len(arguments) != len(f.ParamNames) { // Use f.ParamNames for arity check
		return nil, RuntimeError{
			Message: fmt.Sprintf("function %s expects %d arguments, but got %d",
				f.Name, len(f.ParamNames), len(arguments)), // Use f.ParamNames for error message
			// Line/Col info might be hard to get here accurately for the call site.
			// Could potentially pass the call expression node for better error reporting.
		}
	}

	// Bind arguments to parameter names in the new environment.
	for i, paramName := range f.ParamNames { // Use f.ParamNames for binding
		env.Define(paramName, arguments[i])
	}

	if f.Declaration.Body == nil { // Ensure body is not nil before accepting visitor
		return nil, fmt.Errorf("cannot call function '%s': body is nil", f.Name)
	}

	// Execute the function body in the new environment.
	// executeBlock will handle switching the interpreter's environment.
	result := interpreter.executeBlock(f.Declaration.Body.Statements, env)

	// Check if the result from the function body is a ReturnSignal
	if actualValue, isRetSignal := IsReturnSignal(result); isRetSignal {
		return actualValue, nil // Propagate the unwrapped return value
	}
	if err, ok := result.(RuntimeError); ok {
		return nil, err
	}
	if err, ok := result.(error); ok { // Catch other generic errors
		return nil, err
	}

	return None, nil
}

// StructValue represents an instance of a struct in V
type StructValue struct {
	TypeName string
	Fields   map[string]Value
}

func NewStruct(typeName string) *StructValue {
	return &StructValue{
		TypeName: typeName,
		Fields:   make(map[string]Value),
	}
}

func (v *StructValue) Type() ValueType {
	return StructType
}

func (v *StructValue) String() string {
	var fields []string
	for name, value := range v.Fields {
		fields = append(fields, fmt.Sprintf("%s: %s", name, value.String()))
	}
	return v.TypeName + "{" + strings.Join(fields, ", ") + "}"
}

func (v *StructValue) Equals(other Value) bool {
	otherStruct, ok := other.(*StructValue)
	if !ok || v.TypeName != otherStruct.TypeName || len(v.Fields) != len(otherStruct.Fields) {
		return false
	}

	for name, value := range v.Fields {
		otherValue, exists := otherStruct.Fields[name]
		if !exists || !value.Equals(otherValue) {
			return false
		}
	}
	return true
}

// NoneValue represents the V's none value (similar to null)
type NoneValue struct{}

var None = &NoneValue{} // Singleton instance

func (v *NoneValue) Type() ValueType {
	return NoneType
}

func (v *NoneValue) String() string {
	return "none"
}

func (v *NoneValue) Equals(other Value) bool {
	_, ok := other.(*NoneValue)
	return ok
}

// OptionalValue represents V's optional type (containing a value or none)
type OptionalValue struct {
	Value    Value // Can be None if the optional is empty
	HasValue bool
}

func NewOptional(value Value) *OptionalValue {
	if value == nil || value == None {
		return &OptionalValue{Value: None, HasValue: false}
	}
	return &OptionalValue{Value: value, HasValue: true}
}

func (v *OptionalValue) Type() ValueType {
	return OptionalType
}

func (v *OptionalValue) String() string {
	if !v.HasValue {
		return "none"
	}
	return v.Value.String()
}

func (v *OptionalValue) Equals(other Value) bool {
	otherOpt, ok := other.(*OptionalValue)
	if !ok {
		return false
	}
	if v.HasValue != otherOpt.HasValue {
		return false
	}
	if !v.HasValue { // Both are empty optionals
		return true
	}
	return v.Value.Equals(otherOpt.Value)
}

// ResultValue represents V's result type (success value or error)
type ResultValue struct {
	Value   Value // Value if success
	Error   error // Error if failure
	IsError bool
}

func NewResult(value Value) *ResultValue {
	return &ResultValue{Value: value, Error: nil, IsError: false}
}

func NewResultError(err error) *ResultValue {
	return &ResultValue{Value: nil, Error: err, IsError: true}
}

func (v *ResultValue) Type() ValueType {
	return ResultType
}

func (v *ResultValue) String() string {
	if v.IsError {
		return fmt.Sprintf("error(%s)", v.Error.Error())
	}
	return v.Value.String()
}

func (v *ResultValue) Equals(other Value) bool {
	otherRes, ok := other.(*ResultValue)
	if !ok {
		return false
	}
	if v.IsError != otherRes.IsError {
		return false
	}
	if v.IsError {
		return v.Error.Error() == otherRes.Error.Error()
	}
	return v.Value.Equals(otherRes.Value)
}

// ReturnValue is a wrapper for return values.

type ReturnValue struct {
	Value Value
}

func NewReturnValue(value Value) *ReturnValue {
	return &ReturnValue{Value: value}
}

func (rv *ReturnValue) Type() ValueType {
	if rv.Value != nil {
		return rv.Value.Type()
	}
	return NoneType
}

func (rv *ReturnValue) String() string {
	return fmt.Sprintf("ReturnValue(%s)", rv.Value.String())
}

func (rv *ReturnValue) Error() string {
	return fmt.Sprintf("internal: return signal with value %s", rv.Value.String())
}

func (rv *ReturnValue) Equals(other Value) bool {
	if otherRV, ok := other.(*ReturnValue); ok {
		if rv.Value == nil && otherRV.Value == nil {
			return true
		}
		if rv.Value != nil && otherRV.Value != nil {
			return rv.Value.Equals(otherRV.Value)
		}
	}
	return false
}

// EnumValue represents an enum value in V
type EnumValue struct {
	EnumType string
	Value    string
	IntValue int
}

func NewEnum(enumType string, value string, intValue int) *EnumValue {
	return &EnumValue{
		EnumType: enumType,
		Value:    value,
		IntValue: intValue,
	}
}

func (v *EnumValue) Type() ValueType {
	return EnumType
}

func (v *EnumValue) String() string {
	return fmt.Sprintf("%s.%s", v.EnumType, v.Value)
}

func (v *EnumValue) Equals(other Value) bool {
	otherEnum, ok := other.(*EnumValue)
	if !ok {
		return false
	}
	return v.EnumType == otherEnum.EnumType && v.Value == otherEnum.Value
}

// TypeValue represents a type itself as a runtime value.
// This allows types to be potentially first-class or used in type assertions/reflection.
type TypeValue struct {
	// A string representation of the type, e.g., "int", "string", "MyStruct", "[]int", "map[string]bool"
	Name string
	// OriginalAstNode ast.TypeNode // Can be added if more detailed introspection from AST is needed
}

// NewTypeValue creates a new runtime representation of a type.
func NewTypeValue(name string) *TypeValue {
	return &TypeValue{Name: name}
}

func (tv *TypeValue) Type() ValueType {
	return TypeValueType
}

func (tv *TypeValue) String() string {
	return fmt.Sprintf("<type %s>", tv.Name)
}

// Equals checks if two TypeValues represent the same type.
// For now, it's a simple string comparison of their names.
func (tv *TypeValue) Equals(other Value) bool {
	if otherTv, ok := other.(*TypeValue); ok {
		return tv.Name == otherTv.Name
	}
	return false
}

// VMethod represents a method bound to a struct instance.
type VMethod struct {
	Instance *StructValue // The 'this' or 'self' for the method call
	Method   *VFunction   // The actual function definition
}

// NewVMethod creates a new bound method.
func NewVMethod(instance *StructValue, method *VFunction) *VMethod {
	return &VMethod{
		Instance: instance,
		Method:   method,
	}
}

// Type returns the type of this value (FunctionType).
func (vm *VMethod) Type() ValueType {
	return FunctionType // Methods are a kind of function
}

// String returns a string representation of the bound method.
func (vm *VMethod) String() string {
	return fmt.Sprintf("method %s.%s()", vm.Instance.TypeName, vm.Method.Name)
}

// Equals checks if this bound method is equal to another value.
func (vm *VMethod) Equals(other Value) bool {
	if otherMethod, ok := other.(*VMethod); ok {
		return vm.Instance.Equals(otherMethod.Instance) && vm.Method.Equals(otherMethod.Method)
	}
	return false
}

// Call executes the bound method.
func (vm *VMethod) Call(interpreter *Interpreter, arguments []Value) (Value, error) {
	// Create a new environment for the method call, enclosed by the method's original closure.
	methodEnv := NewEnvironment(vm.Method.Closure)

	// Define the receiver in the method's environment.
	// The receiver's name (e.g., 'p' in 'fn (p Person) get_info()')
	// is in vm.Method.Declaration.Receiver.Name.Name
	if vm.Method.Declaration.Receiver == nil || vm.Method.Declaration.Receiver.Name == nil {
		// This should ideally be caught during AST building or an earlier semantic check
		return nil, fmt.Errorf("internal error: method %s receiver not properly defined in AST (node %s)", vm.Method.Name, vm.Method.Declaration.Name.Name)
	}
	receiverName := vm.Method.Declaration.Receiver.Name.Name
	methodEnv.Define(receiverName, vm.Instance) // Bind the instance

	// Bind arguments to parameters
	if len(arguments) != len(vm.Method.ParamNames) {
		return nil, fmt.Errorf("arity error: method %s expected %d arguments, got %d (node %s)",
			vm.Method.Name, len(vm.Method.ParamNames), len(arguments), vm.Method.Declaration.Name.Name)
	}
	for i, paramName := range vm.Method.ParamNames {
		methodEnv.Define(paramName, arguments[i])
	}

	// Store current environment and set new one for method execution
	previousEnv := interpreter.environment
	interpreter.environment = methodEnv
	defer func() { interpreter.environment = previousEnv }() // Restore previous environment

	if interpreter.DebugMode {
		fmt.Printf("Calling method: %s.%s on %s with closure: %p and methodEnv: %p (enclosing: %p)\n",
			vm.Instance.TypeName, vm.Method.Name, vm.Instance.String(), vm.Method.Closure, methodEnv, methodEnv.enclosing)
		fmt.Printf("  Receiver: %s = %s\n", receiverName, vm.Instance.String())
		for i, param := range vm.Method.Declaration.Parameters {
			fmt.Printf("  Arg %d: %s = %s\n", i, param.Name.Name, arguments[i].String())
		}
	}

	// Execute the method body in the new environment.
	// VisitBlockStmt is expected to handle ReturnValue and errors.
	blockResult := interpreter.VisitBlockStmt(vm.Method.Declaration.Body)

	if actualValue, isRetSignal := IsReturnSignal(blockResult); isRetSignal {
		return actualValue, nil // Propagate the unwrapped return value
	}
	if errVal, ok := blockResult.(RuntimeError); ok {
		return nil, errVal // Propagate runtime error
	}
	if errVal, ok := blockResult.(error); ok { // Catch other generic errors
		// Ensure we don't wrap a RuntimeError again, though IsReturnSignal should catch it if it's wrapped in ReturnSignal
		return nil, fmt.Errorf("unexpected error during method '%s' execution on %s: %w", vm.Method.Name, vm.Method.Declaration.Name.Name, errVal)
	}

	// If no explicit return, V functions (and thus methods) return 'None' by default.
	return None, nil
}

// Ensure VMethod satisfies the Function interface (if defined with Call method).
// var _ Function = (*VMethod)(nil) // Assuming Function interface has: Call(*Interpreter, []Value) (Value, error)

// PointerValue represents a pointer to a variable in an environment.
type PointerValue struct {
	VariableName  string
	ReferencedEnv *Environment
}

func NewPointerValue(variableName string, env *Environment) *PointerValue {
	return &PointerValue{VariableName: variableName, ReferencedEnv: env}
}

func (v *PointerValue) Type() ValueType {
	return PointerType
}

func (v *PointerValue) String() string {
	// Avoid causing an infinite loop if a pointer refers to itself through a variable.
	// Just show the variable name it points to.
	return fmt.Sprintf("&%s", v.VariableName)
}

func (v *PointerValue) Equals(other Value) bool {
	if otherPtr, ok := other.(*PointerValue); ok {
		// Two pointers are equal if they point to the same variable in the same environment.
		return v.VariableName == otherPtr.VariableName && v.ReferencedEnv == otherPtr.ReferencedEnv
	}
	return false
}

// Dereference retrieves the value that the pointer points to.
// Dereference retrieves the value that the pointer points to.
func (v *PointerValue) Dereference() (Value, error) {
	val, err := v.ReferencedEnv.Get(v.VariableName)
	if err != nil {
		// Wrap the original error to provide more context.
		return nil, fmt.Errorf("dangling pointer: %w", err)
	}
	return val, nil
}
