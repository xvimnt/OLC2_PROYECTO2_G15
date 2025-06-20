// 04_slices.vch (Expanded and Corrected)
// ======================================
// Tests all features of single and multi-dimensional slices (arrays).
// - Uses the correct syntax: `:= []type{...}` for declaration and initialization.
// - Tests slices of ALL primitive types.
// - Thoroughly tests built-in functions: len, append, indexOf, join.
// - Includes the range-based `for...in` loop.
// - Tests multi-dimensional slices, including JAGGED slices and appending rows.
// - Includes tests for edge cases like empty slices and out-of-bounds access.

fn main() {
    println("--- Test 4: Slices (Arrays) ---")

    // --------------------------------------------------
    // Section A: Basic Slice Operations
    // --------------------------------------------------
    println("\n--- A. Basic Slice Operations ---")
    mut numbers := []int{10, 20, 30}
    println("Initial int slice:", numbers)
    
    // 5.1.4. Función len
    println("Length of slice:", len(numbers)) // Expected: 3
    
    // 5.1.6. Acceso de elemento (Read)
    println("Element at index 1:", numbers[1]) // Expected: 20
    
    // 5.1.6. Acceso de elemento (Write/Modify)
    numbers[0] = 5
    println("Modified slice (numbers[0]=5):", numbers) // Expected: {5, 20, 30}

    // Testing slices of other primitive types
    mut floats := []float64{1.1, 2.2, 3.3}
    println("\nFloat64 slice:", floats)
    mut bools := []bool{true, false, true}
    println("Bool slice:", bools)
    mut runes := []rune{'a', 'b', 'c'}
    println("Rune slice:", runes)


    // --------------------------------------------------
    // Section B: Built-in Slice Functions
    // --------------------------------------------------
    println("\n--- B. Built-in Slice Functions ---")

    // 5.1.5. Función append
    mut str_slice := []string{"VLang"}
    println("Appending to string slice:", str_slice)
    str_slice = append(str_slice, "is")
    str_slice = append(str_slice, "awesome")
    println("After appending twice:", str_slice) // Expected: {"VLang", "is", "awesome"}
    
    // 5.1.3. Funcion Join (only for []string)
    println("Joined string:", join(str_slice, "-")) // Expected: "VLang-is-awesome"

    // 5.1.2. Función indexOf
    mut search_slice := []int{5, 15, 25, 15, 35}
    println("\nSearching in:", search_slice)
    println("indexOf 25 (found):", indexOf(search_slice, 25))        // Expected: 2
    println("indexOf 15 (finds first):", indexOf(search_slice, 15))   // Expected: 1
    println("indexOf 99 (not found):", indexOf(search_slice, 99))    // Expected: -1


    // --------------------------------------------------
    // Section C: Range-Based For Loop
    // --------------------------------------------------
    println("\n--- C. Range-Based For Loop (for...in) ---")
    mut items := []string{"alpha", "beta", "gamma"}
    println("Iterating over:", items)
    for index, value in items {
        print("($index:$value) ") // Expected: (0:alpha) (1:beta) (2:gamma) 
    }
    println("")


    // --------------------------------------------------
    // Section D: Multi-dimensional Slices
    // --------------------------------------------------
    println("\n--- D. Multi-dimensional Slices ---")
    
    // Standard matrix
    mut matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6}
    }
    println("Matrix[1][2]:", matrix[1][2]) // Expected: 6
    matrix[0][0] = 99
    println("Modified Matrix[0][0]:", matrix[0][0]) // Expected: 99

    // Jagged slice (as specified in PDF)
    mut jagged := [][]int{
        {1},
        {2, 3},
        {4, 5, 6}
    }
    println("\nJagged slice element [2][1]:", jagged[2][1]) // Expected: 5

    // Appending a new row (as specified in PDF)
    mut new_row := []int{7, 8, 9}
    jagged = append(jagged, new_row)
    println("Jagged slice after append:", jagged) // Expected: {{1}, {2, 3}, {4, 5, 6}, {7, 8, 9}}


    // --------------------------------------------------
    // Section E: Edge Cases and Errors
    // --------------------------------------------------
    println("\n--- E. Edge Cases and Errors ---")
    
    // Empty slice
    mut empty_slice := []int{}
    println("Length of empty slice:", len(empty_slice)) // Expected: 0
    
    // Appending to an empty slice
    empty_slice = append(empty_slice, 100)
    println("Appended to empty slice:", empty_slice) // Expected: {100}
    
    // Out-of-bounds access (commented out to allow script to run).
    // This should produce a runtime error in a correct interpreter.
    // println(numbers[3]) // SEMANTIC/RUNTIME ERROR: index out of range

    // --------------------------------------------------
    // Section F: Multi-dimensional Slices
    // --------------------------------------------------
    println("\n--- F. Multi-dimensional Slices ---")
    mut matrix := [][]int{
        {10, 20, 30},
        {40, 50, 60},
        {70, 80, 90}
    }
    println("Element matrix[1][0]: ${matrix[1][0]}") // Expected: Element matrix[1][0]: 40
    println("Sub-slice matrix[0]: ${matrix[0]}")   // Expected: Sub-slice matrix[0]: [10, 20, 30]

    // Test accessing an element of a sub-slice
    mut first_row := matrix[0]
    println("Element first_row[2]: ${first_row[2]}") // Expected: Element first_row[2]: 30
    matrix[1][1] = 55
    println("Matrix after modification: $matrix") // Expected: Matrix: [[10, 20, 30], [40, 55, 60], [70, 80, 90]]
    println("Element matrix[1][1] after modification: ${matrix[1][1]}") // Expected: 55

    println("\n--- Test 4 Finished ---")
}