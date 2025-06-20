# V Language Interpreter + ARM64 Translator

A Go-based interpreter for the V programming language using ANTLR4 with comprehensive test-driven development (TDD).

## Project Overview

This interpreter implements the V programming language, providing an execution environment that parses, analyzes, and runs V language code. The implementation follows the visitor pattern and uses ANTLR4 for parsing.

### Features

- ANTLR4-based grammar for the V language
- Lexer and parser for V syntax
- AST (Abstract Syntax Tree) construction
- Interpreter with visitor pattern implementation
- Support for core V language features:
  - Variables and constants
  - Data types (int, f64, string, bool, arrays, maps)
  - Control flow (if, match, for, while)
  - Functions (including methods, anonymous functions)
  - Structs, interfaces, and modules
  - Error handling (optionals and results)
  - String operations

## Project Structure

```
OLC2_PROYECTO2_G15/
├── grammar/            # ANTLR4 grammar definitions
│   └── MyLang.g4       # Grammar file for V language
├── parser/             # Generated ANTLR parser (generated code)
├── lexer/              # Lexer implementation
│   └── lexer.go        # Lexer wrapper for ANTLR lexer
├── ast/                # Abstract Syntax Tree definitions
│   ├── nodes.go        # AST node types
│   └── visitor.go      # Visitor interface and base visitor
├── interpreter/        # Interpreter implementation
│   ├── interpreter.go  # Core interpreter logic
│   ├── environment.go  # Symbol table and scopes
│   └── values.go       # Runtime values
├── translator/         # Translator implementation
│   ├── translator.go   # Core translator logic
├── cmd/                # Command-line interface
│   └── main.go         # CLI entry point
├── examples/           # Example V language programs
│   ├── basic.mylang
│   ├── conditionals.mylang
│   ├── loops.mylang
│   ├── functions.mylang
│   └── structs.mylang
├── tests/              # Test files
│   └── interpreter_test.go
├── Makefile            # Build automation
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- Java Runtime Environment (for ANTLR4 code generation)
- ANTLR4 Tool (downloaded automatically via the Makefile)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/xvimnt/OLC2_PROYECTO2_G15.git
cd OLC2_PROYECTO2_G15
```

2. Generate the parser and build the interpreter:

```bash
make all
```

This will:
- Download ANTLR4 if needed
- Generate Go code from the grammar file
- Build the interpreter binary

### Running the Interpreter

#### Run a V language file:

```bash
./OLC2_PROYECTO2_G15 run examples/basic.mylang
```

#### Start the REPL (interactive mode):

```bash
./OLC2_PROYECTO2_G15 repl
```

#### Parse a file and display the AST:

```bash
./OLC2_PROYECTO2_G15 parse examples/basic.mylang
```

## Development

### Testing

Run the test suite:

```bash
make test
```

Generate test coverage report:

```bash
make coverage
```

### Adding Features

1. Update the grammar in `grammar/MyLang.g4`
2. Define corresponding AST nodes in `ast/nodes.go`
3. Add visitor methods in `ast/visitor.go`
4. Implement the interpretation logic in `interpreter/interpreter.go`
5. Write test cases in `tests/`

## V Language Syntax

The V language implemented in this interpreter supports the following syntax:

### Module Declaration

```v
module main
```

### Import Statements

```v
import os
import math
import mymodule.submodule
```

### Variable Declarations

```v
x := 5                  // Inferred type
mut y := 10             // Mutable variable
const pi = 3.14159      // Constant
```

### Data Types

```v
// Basic types
i := 42                 // int
f := 3.14               // f64
b := true               // bool
s := 'hello'            // string
c := `A`                // char

// Arrays
arr := [1, 2, 3]        // Array of int
mut nums := []int{len: 5, cap: 10}  // Empty array with capacity

// Maps
m := map[string]int{}   // Empty map
m['one'] = 1            // Assignment
```

### Control Flow

```v
// If statement
if x > 0 {
    println('positive')
} else if x == 0 {
    println('zero')
} else {
    println('negative')
}

// Match statement (switch)
match x {
    1 { println('one') }
    2 { println('two') }
    else { println('other') }
}

// For loop
for i in 0..5 {
    println(i)
}

// While-style for loop
mut i := 0
for i < 5 {
    i++
}

// Infinite loop
for {
    if i > 10 { break }
    i++
}
```

### Functions

```v
// Function declaration
fn add(x int, y int) int {
    return x + y
}

// Multiple return values using struct
struct DivResult {
    quotient int
    remainder int
}

fn divide(x int, y int) DivResult {
    return DivResult{
        quotient: x / y
        remainder: x % y
    }
}

// Anonymous function
add_fn := fn(x int, y int) int {
    return x + y
}
```

### Structs and Methods

```v
struct Point {
    x int
    y int
}

// Method on struct
fn (p Point) distance_from_origin() f64 {
    return math.sqrt(p.x * p.x + p.y * p.y)
}

// Mutable method
fn (mut p Point) move(dx int, dy int) {
    p.x += dx
    p.y += dy
}
```

See the `examples/` directory for more complete code samples.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- V language design: https://vlang.io/
- ANTLR4: https://www.antlr.org/
