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
    println("Inferred string via ':='")
    println("Inferred float via ':='")
    
    // Type inference with `mut` keyword
    mut mutable_inferred_bool := true
    println("Inferred mutable bool via 'mut ... :='")
    mutable_inferred_bool = false // Should be allowed
    println("Modified mutable bool:", mutable_inferred_bool) // Expected: false

    // Reassignment is allowed for mutable variables
    explicit_int = 100
    println("Reassigned int:", explicit_int) // Expected: 100

    // Type mismatch on reassignment (commented out to prevent semantic error).
    // explicit_int = "not an int" // SEMANTIC ERROR: Cannot assign string to int.

    
    // --------------------------------------------------
    // Section D: Scope Rules and Shadowing
    // --------------------------------------------------
    println("\n--- D. Testing Scope Rules and Shadowing ---")
    mut outer_scope_var int = 1
    println("1. outer_scope_var in main scope:", outer_scope_var) // Expected: 1
    println("2. Accessing global from main scope:", global_integer)  // Expected: 100

    {
        // First level of nesting
        mut inner_scope_var int = 2
        println("3. outer_scope_var inside block 1:", outer_scope_var) // Expected: 1
        println("4. inner_scope_var inside block 1:", inner_scope_var) // Expected: 2

        // Shadowing a variable from an outer scope
        mut outer_scope_var int = 99 // This 'outer_scope_var' only exists here.
        println("5. Shadowed outer_scope_var in block 1:", outer_scope_var) // Expected: 99
        
        {
            // Second level of nesting
            mut deepest_var bool = true
            println("6. Accessing inner_scope_var from block 2:", inner_scope_var) // Expected: 2
            println("7. Accessing shadowed var from block 2:", outer_scope_var) // Expected: 99 (accesses the one from block 1)
        }
        
        // This would be a semantic error as `deepest_var` is out of scope.
        // println(deepest_var) 
    }

    println("8. outer_scope_var back in main scope:", outer_scope_var) // Expected: 1 (shadowed variable is gone)
    
    // This would be a semantic error as `inner_scope_var` is out of scope.
    // println(inner_scope_var) 

    println("\n--- Test 1 Finished ---")
}