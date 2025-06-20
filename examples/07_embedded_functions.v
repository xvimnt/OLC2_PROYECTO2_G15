// 07_embedded_functions.vch (Expanded and Corrected)
// ====================================================
// Tests all built-in (embedded) functions as per the language specification.
// - All function declarations correctly use the `fn` keyword.
// - All `println` variations and permitted types are tested for correct output format.
// - `Atoi` and `parseFloat` are tested with valid and invalid inputs.
// - `typeOf` is tested against all primitive and composite types.
// - Includes other embedded functions like `len` for completeness.

// A struct needed for testing `println` and `typeOf`.
struct Vehicle {
    string make
    int    year
}

fn main() {
    println("--- Test 7: Embedded Functions ---")
    
    // --------------------------------------------------
    // Section A: Testing println (7.2.1)
    // --------------------------------------------------
    println("\n--- A. Testing println ---")
    
    // With zero arguments (should just print a newline)
    println("Line 1")
    println()
    println("Line 3 (should be a blank line above this)")

    // With multiple arguments of different primitive types
    println("\nMultiple primitives:", 10, -5.5, true, 'X', "hello")
    // Expected: Multiple primitives: 10 -5.5 true X hello

    // With permitted composite types (7.2.1.1)
    mut my_slice := []int{1, 2, 3}
    mut my_vehicle := Vehicle{make: "Toyota", year: 2024}
    println("Slice output:", my_slice)    // Expected: Slice output: {1, 2, 3}
    println("Struct output:", my_vehicle) // Expected: Struct output: Vehicle{make: "Toyota", year: 2024}


    // --------------------------------------------------
    // Section B: Testing Type Conversion Functions
    // --------------------------------------------------
    println("\n--- B. Testing Type Conversion ---")
    
    // 7.2.2: Atoi (string to int)
    mut int_val := Atoi("12345")
    println("Atoi('12345') results in:", int_val)
    // This should cause an error as per the spec (no rounding/truncation)
    // mut bad_int := Atoi("123.45") // SEMANTIC ERROR

    // 7.2.3: parseFloat (string to float64)
    mut float_val_1 := parseFloat("123.45")
    mut float_val_2 := parseFloat("50") // Integer strings should also work
    println("parseFloat('123.45') results in:", float_val_1)
    println("parseFloat('50') results in:", float_val_2) // Expected: 50.0
    // This should cause an error
    // mut bad_float := parseFloat("not a number") // SEMANTIC ERROR


    // --------------------------------------------------
    // Section C: Testing Type Reflection
    // --------------------------------------------------
    println("\n--- C. Testing typeOf ---")

    // 7.2.4: typeOf with all specified types
    println("Type of 42:", typeOf(42))                  // Expected: int
    println("Type of 3.14:", typeOf(3.14))              // Expected: float64
    println("Type of 'hello':", typeOf("hello"))        // Expected: string
    println("Type of true:", typeOf(true))              // Expected: bool
    println("Type of 'R':", typeOf('R'))                // Expected: rune
    println("Type of my_slice:", typeOf(my_slice))      // Expected: []int
    println("Type of my_vehicle:", typeOf(my_vehicle))  // Expected: Vehicle
    println("Type of a multi-dim slice:", typeOf([][]int{{1}})) // Expected: [][]int


    // --------------------------------------------------
    // Section D: Other Embedded Functions
    // --------------------------------------------------
    println("\n--- D. Testing Other Embedded Functions ---")

    // len is an embedded function primarily used with slices.
    mut test_len_slice := []string{"a", "b", "c", "d"}
    println("len of a 4-element slice:", len(test_len_slice)) // Expected: 4
    
    // indexOf, append, and join are tested more thoroughly in 04_slices.vch

    println("\n--- Test 7 Finished ---")
}