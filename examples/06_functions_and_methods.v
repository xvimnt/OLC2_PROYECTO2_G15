// 06_functions_and_methods.vch (Expanded and Corrected)
// ====================================================
// Tests all features of functions and methods as per the language specification.
// - All declarations use the correct `fn` keyword.
// - Global-only declaration rule for functions and structs.
// - Parameter passing: by value for primitives, by reference for slices/structs.
// - Various function signatures (with/without params, with/without return).
// - Methods attached to structs (value and pointer receivers).
// - Verification of return types.
// - Testing for disallowed features like function overloading.

// --------------------------------------------------
// Section A: Global Struct and Function Declarations
// --------------------------------------------------
struct Data {
    int value
    string tag
}

// -- Methods (functions associated with a struct) --

// Method with a value receiver (cannot modify the original struct).
fn (d Data) check_value() bool {
    // d.value = 999 // This would not affect the original `d`
    return d.value > 100
}

// Method with a pointer receiver (modifies the original struct).
// This tests the "operan por referencia" concept.
fn (d *Data) update_tag(new_tag string) {
    d.tag = new_tag
}

// -- Regular Functions --

// Function with no parameters, no return value.
fn print_separator() {
    println("------------------------------")
}

// Function with parameters, no return value (pass by value).
fn test_pass_by_value(num int, text string) {
    num = 999  // This change is local and won't affect the caller.
    text = "modified" // This change is also local.
    println("Inside test_pass_by_value -> num:", num, ", text:", text)
}

// Function with a slice parameter (pass by reference).
fn add_to_slice(s []int) {
    s[0] = 111 // This modification WILL affect the caller's slice.
}

// Function with parameters and a return value.
fn multiply(a int, b int) int {
    return a * b
}

// Function returning a composite type (slice).
fn generate_sequence(start int, count int) []int {
    mut sequence := []int{}
    for mut i := 0; i < count; i = i + 1 {
        sequence = append(sequence, start + i)
    }
    return sequence
}

// The spec disallows function overloading.
// Uncommenting this would cause a semantic error.
// fn multiply(a float64, b float64) float64 { 
//    return a * b
// }


// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------
fn main() {
    println("--- Test 6: Functions and Methods ---")
    
    // Function declaration is not allowed inside another function.
    // fn local_fn() {} // SYNTAX ERROR
    
    // --------------------------------------------------
    // B. Calling Regular Functions
    // --------------------------------------------------
    println("\n--- B. Calling Regular Functions ---")
    print_separator()
    
    // Test pass by value for primitives
    mut original_num int = 10
    mut original_text string = "original"
    println("Before pass_by_value -> num:", original_num, ", text:", original_text)
    test_pass_by_value(original_num, original_text)
    println("After pass_by_value -> num:", original_num, ", text:", original_text) // Values should be unchanged.
    print_separator()

    // Test pass by reference for slices
    mut my_slice := []int{1, 2, 3}
    println("Before add_to_slice:", my_slice)
    add_to_slice(my_slice)
    println("After add_to_slice:", my_slice) // First element should be 111.
    print_separator()
    
    // Test function with return value
    mut product := multiply(7, 6)
    println("Result of multiply(7, 6):", product) // Expected: 42
    
    // Test function returning a slice
    mut seq := generate_sequence(5, 4)
    println("Generated sequence:", seq) // Expected: {5, 6, 7, 8}
    print_separator()
    
    // --------------------------------------------------
    // C. Calling Methods on Structs
    // --------------------------------------------------
    println("\n--- C. Calling Methods on Structs ---")
    mut my_data := Data{value: 200, tag: "initial"}
    println("Initial Data:", my_data)
    
    // Call value-receiver method
    mut is_large := my_data.check_value()
    println("Is my_data.value > 100?", is_large) // Expected: true
    
    // Call pointer-receiver method to modify the struct
    my_data.update_tag("updated")
    println("Data after update_tag:", my_data) // Expected: {value: 200, tag: "updated"}

    println("\n--- Test 6 Finished ---")
}