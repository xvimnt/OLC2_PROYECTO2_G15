// 01_basics_and_scope.vch (Expanded)
// ===================================
// Tests core language features:
// - Valid and invalid identifiers.
// - Case sensitivity for variables, functions, and keywords.
// - Comment handling (single-line, block, nested).
// - Variable declarations (with and without initial values, type inference).
// - Scope rules (global, function, block-level, and shadowing).
// - Default values for uninitialized variables.

// --------------------------------------------------
// 1. Global Scope Declarations
// --------------------------------------------------

// 3.5. Comments at the top level
/*
  This is a global-level block comment.
  It should be ignored by the parser.
*/

// Global variables with various valid identifier formats.
mut global_integer int = 100
mut _global_string string = "I am a global variable."
mut globalVarWithCaps int = 200

// This function will be used to test case sensitivity in function calls.
fn A_Function() {
    println("Called A_Function (uppercase)")
}

fn a_function() {
    println("Called a_function (lowercase)")
}

// --------------------------------------------------
// 2. Main Execution Block
// --------------------------------------------------

fn main() {
    println("--- Test 1: Basics, Scope, and Declarations ---")

    // --------------------------------------------------
    // Section A: Comments and Identifiers
    // --------------------------------------------------
    println("\n--- A. Testing Comments and Identifiers ---")
    // A single-line comment inside a function.
    println("This line should execute.") // This comment should be ignored.
    /* Block comment inside a function */ println("This one too.")

    // 3.3. Valid Identifiers
    mut i_am_valid int = 1
    mut _can_start_with_underscore int = 2
    mut var123isFine int = 3
    println("Valid identifiers declared successfully.")

    // Invalid identifiers (commented out to prevent syntax errors).
    // mut 1_is_not_valid int = 1    // SYNTAX ERROR: Starts with a number.
    // mut not-valid-either int = 2  // LEXICAL ERROR: Contains invalid character '-'.
    // mut for int = 3               // SYNTAX ERROR: 'for' is a reserved keyword.


    // --------------------------------------------------
    // Section B: Case Sensitivity
    // --------------------------------------------------
    println("\n--- B. Testing Case Sensitivity ---")
    mut myvar int = 10
    mut MyVar int = 20 // Different variable due to case sensitivity.
    
    println("value of 'myvar': $myvar") // Expected: 10
    println("value of 'MyVar': $MyVar") // Expected: 20

    // Case sensitivity in function calls.
    A_Function() // Expected: "Called A_Function (uppercase)"
    a_function() // Expected: "Called a_function (lowercase)"

    // Case sensitivity in keywords (commented out to prevent syntax errors).
    // MUT x int = 5    // SYNTAX ERROR: 'MUT' is not 'mut'.
    // IF true {}       // SYNTAX ERROR: 'IF' is not 'if'.


    // --------------------------------------------------
    // Section C: Variable Declarations and Default Values
    // --------------------------------------------------
    println("\n--- C. Testing Variable Declarations ---")
    
    // Declaration with explicit type, no initial value.
    mut uninitialized_int int
    mut uninitialized_float float64
    mut uninitialized_string string
    mut uninitialized_bool bool
    println("Default int value:", uninitialized_int)       // Expected: 0
    println("Default float64 value:", uninitialized_float) // Expected: 0.0
    println("Default string value: '$uninitialized_string'") // Expected: "" (empty string)
    println("Default bool value:", uninitialized_bool)     // Expected: false

    // Declaration with explicit type and initial value.
    mut explicit_int int = -50
    println("Explicitly initialized int:", explicit_int) // Expected: -50

    // Type inference using `:=`
    inferred_string := "This is a string"
    inferred_float := 3.14
    println("Inferred string via ':=' $inferred_string")
    println("Inferred float via ':=' $inferred_float")
    
   
}