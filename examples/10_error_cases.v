// 08_error_cases.vch (Expanded and Corrected)
// ============================================
// This file contains a comprehensive suite of invalid code.
// It is designed to test the interpreter's error detection and reporting.
// A correct interpreter should identify these errors and NOT execute the code.

// --------------------------------------------------
// Section A: Lexical Errors (Invalid Tokens)
// --------------------------------------------------

// An invalid character in an identifier.
// mut invalid@char int = 1 // LEXICAL ERROR

// An unterminated string literal.
// mut unterminated_str string = "hello world  // LEXICAL ERROR

// An invalid number format.
// mut bad_float float64 = 1.2.3 // LEXICAL ERROR


// --------------------------------------------------
// Section B: Syntax Errors (Incorrect Grammar)
// --------------------------------------------------

// Missing closing parenthesis in function call.
// println("hello" // SYNTAX ERROR

// Struct definition inside a function.
// fn invalid_struct_placement() {
//    struct Local {} // SYNTAX ERROR
// }

// Missing semicolon in a C-style for loop.
// for mut i := 0 i < 10; i = i + 1 {} // SYNTAX ERROR

// Mismatched braces.
// fn main() {
//    if true {
//        println("Missing closing brace")
// } // SYNTAX ERROR

// Keyword misuse.
// let x := 5 // SYNTAX ERROR: 'let' is not a keyword.


// --------------------------------------------------
// Section C: Semantic Errors (Declaration and Scope)
// --------------------------------------------------

// Redeclaring a variable in the same scope.
// fn scope_errors() {
//    mut x int = 10
//    mut x int = 20 // SEMANTIC ERROR: 'x' already declared in this scope.
// }

// Using a variable before it is declared.
// fn undeclared_var_error() {
//    x = 10 // SEMANTIC ERROR: 'x' is not declared.
//    mut x int
// }

// Accessing a variable outside its scope.
// fn out_of_scope_error() {
//    if true {
//        mut inner_var int = 5
//    }
//    println(inner_var) // SEMANTIC ERROR: 'inner_var' not defined in this scope.
// }


// --------------------------------------------------
// Section D: Semantic Errors (Type Mismatches)
// --------------------------------------------------

// Assigning a string to an int variable.
// fn type_mismatch_assignment() {
//    mut num int = 10
//    num = "not an int" // SEMANTIC ERROR: Type mismatch.
// }

// Using a non-boolean expression in an if statement.
// fn type_mismatch_if() {
//    if 10 { // SEMANTIC ERROR: If condition must be a boolean.
//        println("This will fail.")
//    }
// }

// Performing an invalid operation (e.g., subtracting strings).
// fn type_mismatch_operator() {
//    mut result := "hello" - "world" // SEMANTIC ERROR: Operator '-' not defined for strings.
// }

// Passing the wrong type of argument to a function.
// fn takes_int(n int) {}
// fn type_mismatch_call() {
//    takes_int(true) // SEMANTIC ERROR: Expected int, got bool.
// }


// --------------------------------------------------
// Section E: Semantic Errors (Functions, Structs, Slices)
// --------------------------------------------------

// Calling a non-function variable.
// fn non_function_call() {
//    mut x := 10
//    x() // SEMANTIC ERROR: 'x' is not a function.
// }

// Wrong number of arguments in a function call.
// fn requires_two(a int, b int) {}
// fn wrong_arg_count() {
//    requires_two(1) // SEMANTIC ERROR: Expected 2 arguments, got 1.
// }

// Accessing a non-existent struct field.
// struct Point { int x }
// fn bad_field_access() {
//    mut p := Point{x: 1}
//    println(p.y) // SEMANTIC ERROR: Struct 'Point' has no field 'y'.
// }

// Indexing a non-slice/array type.
// fn bad_index() {
//    mut not_a_slice int = 5
//    println(not_a_slice[0]) // SEMANTIC ERROR: Type 'int' does not support indexing.
// }


// --------------------------------------------------
// Section F: Semantic Errors (Runtime-like)
// --------------------------------------------------

// Division by zero.
// fn div_by_zero() {
//    mut result := 100 / 0 // SEMANTIC ERROR: Division by zero.
// }

// Index out of bounds (might be a runtime error depending on implementation).
// fn index_out_of_bounds() {
//    mut s := []int{1, 2}
//    println(s[2]) // SEMANTIC/RUNTIME ERROR: Index out of bounds.
// }