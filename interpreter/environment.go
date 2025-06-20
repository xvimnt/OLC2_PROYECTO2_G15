package interpreter

import (
	"fmt"
	"strings"
)

// Environment represents a lexical scope for variables in the interpreter
type Environment struct {
	values    map[string]Value
	enclosing *Environment
	globals   *Environment // Reference to the top-level environment
}

// NewEnvironment creates a new environment with optional enclosing environment
func NewEnvironment(enclosing *Environment) *Environment {
	env := &Environment{
		values:    make(map[string]Value),
		enclosing: enclosing,
	}

	if enclosing == nil {
		// This is a global environment
		env.globals = env
	} else if enclosing.globals != nil {
		// Inherit globals reference from parent
		env.globals = enclosing.globals
	}

	return env
}

// Define defines a new variable in the current environment
func (e *Environment) Define(name string, value Value) {
	e.values[name] = value
}

// Get retrieves a variable's value from the environment
// It searches in the current environment and then in any enclosing environments
func (e *Environment) Get(name string) (Value, error) {
	if value, ok := e.values[name]; ok {
		return value, nil
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	return nil, fmt.Errorf("undefined variable '%s'", name)
}

// GetAt gets a value from an environment at a specific depth in the chain
func (e *Environment) GetAt(distance int, name string) (Value, error) {
	ancestor := e.ancestor(distance)
	if value, ok := ancestor.values[name]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("undefined variable '%s' at distance %d", name, distance)
}

// Assign updates an existing variable's value
// It searches in the current environment and then in any enclosing environments
func (e *Environment) Assign(name string, value Value) error {
	if _, ok := e.values[name]; ok {
		e.values[name] = value
		return nil
	}

	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}

	return fmt.Errorf("undefined variable '%s'", name)
}

// AssignAt assigns a value to a variable at a specific depth in the environment chain
func (e *Environment) AssignAt(distance int, name string, value Value) error {
	ancestor := e.ancestor(distance)
	if _, ok := ancestor.values[name]; ok {
		ancestor.values[name] = value
		return nil
	}
	return fmt.Errorf("cannot assign to undefined variable '%s' at distance %d", name, distance)
}

// ancestor returns the environment that is 'distance' steps from the current one
func (e *Environment) ancestor(distance int) *Environment {
	environment := e
	for i := 0; i < distance; i++ {
		environment = environment.enclosing
	}
	return environment
}

// GetGlobal gets a value from the global environment
func (e *Environment) GetGlobal(name string) (Value, error) {
	if e.globals == nil {
		return nil, fmt.Errorf("global environment not accessible")
	}

	if value, ok := e.globals.values[name]; ok {
		return value, nil
	}

	return nil, fmt.Errorf("undefined global variable '%s'", name)
}

// DefineGlobal defines a variable in the global environment
func (e *Environment) DefineGlobal(name string, value Value) error {
	if e.globals == nil {
		return fmt.Errorf("global environment not accessible")
	}

	e.globals.values[name] = value
	return nil
}

// Clone creates a copy of the environment with the same values (but not the same references)
func (e *Environment) Clone() *Environment {
	cloned := NewEnvironment(e.enclosing)
	for k, v := range e.values {
		cloned.values[k] = v // Shallow copy of values
	}
	cloned.globals = e.globals
	return cloned
}

// SymbolTable manages variable resolution and scope tracking
type SymbolTable struct {
	scopes       []map[string]bool
	currentScope int
}

// NewSymbolTable creates a new symbol table for tracking variables
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		scopes:       []map[string]bool{make(map[string]bool)}, // Start with global scope
		currentScope: 0,
	}
}

// BeginScope starts a new scope for variable declarations
func (s *SymbolTable) BeginScope() {
	s.scopes = append(s.scopes, make(map[string]bool))
	s.currentScope++
}

// EndScope ends the current scope
func (s *SymbolTable) EndScope() {
	if s.currentScope > 0 {
		s.scopes = s.scopes[:len(s.scopes)-1]
		s.currentScope--
	}
}

// Declare adds a variable to the current scope
func (s *SymbolTable) Declare(name string) error {
	if s.currentScope < 0 || s.currentScope >= len(s.scopes) {
		return fmt.Errorf("invalid scope")
	}

	scope := s.scopes[s.currentScope]
	if _, exists := scope[name]; exists {
		return fmt.Errorf("variable '%s' already declared in this scope", name)
	}

	scope[name] = true
	return nil
}

// Resolve finds the scope depth of a variable
func (s *SymbolTable) Resolve(name string) (int, bool) {
	for i := s.currentScope; i >= 0; i-- {
		if _, exists := s.scopes[i][name]; exists {
			return s.currentScope - i, true
		}
	}

	return -1, false
}

// TypeEnvironment tracks type information for variables
type TypeEnvironment struct {
	types     map[string]string // Maps variable names to their type names
	enclosing *TypeEnvironment
}

// NewTypeEnvironment creates a new type environment
func NewTypeEnvironment(enclosing *TypeEnvironment) *TypeEnvironment {
	return &TypeEnvironment{
		types:     make(map[string]string),
		enclosing: enclosing,
	}
}

// Define adds a variable type to the current scope
func (te *TypeEnvironment) Define(name string, typeName string) {
	te.types[name] = typeName
}

// Get retrieves a variable's type from the environment
func (te *TypeEnvironment) Get(name string) (string, bool) {
	if typeName, ok := te.types[name]; ok {
		return typeName, true
	}

	if te.enclosing != nil {
		return te.enclosing.Get(name)
	}

	return "", false
}

// Assign updates a variable's type
func (te *TypeEnvironment) Assign(name string, typeName string) bool {
	if _, ok := te.types[name]; ok {
		te.types[name] = typeName
		return true
	}

	if te.enclosing != nil {
		return te.enclosing.Assign(name, typeName)
	}

	return false
}

// DumpSymbols devuelve una representación tabular de las variables del entorno actual
func (e *Environment) DumpSymbols() string {
	var sb strings.Builder
	sb.WriteString("Nombre\tTipo\tValor\n")
	for name, val := range e.values {
		sb.WriteString(name)
		sb.WriteString("\t")
		sb.WriteString(val.Type().String())
		sb.WriteString("\t")
		sb.WriteString(val.String())
		sb.WriteString("\n")
	}
	return sb.String()
}
