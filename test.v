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

fn main() {
    mut i_am_valid int = 1
    println("value of 'i_am_valid': $i_am_valid")
    println("value of 'global_integer': $global_integer")
    println("value of 'globalVarWithCaps': $globalVarWithCaps")
    println("value of '_global_string': $_global_string")
    A_Function()
    a_function()
}