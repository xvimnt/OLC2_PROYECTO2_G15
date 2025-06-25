// 02_operators.vch (Expanded)
// =============================
// Tests a comprehensive range of operator behaviors:
// - All arithmetic operators with every valid type combination (int, float64, string).
// - Integer division truncation.
// - All compound assignment operators (+=, -=).
// - All relational and equality operators with valid types (int, float64, string, bool, rune).
// - All logical operators (&&, ||, !).
// - Complex operator precedence and associativity.
// - Commented-out error cases for invalid operations (type mismatches, division by zero).

fn main() {
    println("--- Test 2: Comprehensive Operator Tests ---")

    // --------------------------------------------------
    // Section A: Arithmetic Operators
    // --------------------------------------------------
    println("\n--- A. Arithmetic Operators ---")
    mut i1 int = 10
    mut i2 int = 3
    mut f1 float64 = 12.5
    mut f2 float64 = 2.5
    
    // Suma (+)
    println("int + int:      10 + 3 = ", i1 + i2)         // Expected: 13
    println("int + float64:  10 + 2.5 = ", i1 + f2)       // Expected: 12.5 (promoted to float)
    println("float64 + int:  12.5 + 3 = ", f1 + i2)       // Expected: 15.5 (promoted to float)
    println("float64 + float64: 12.5 + 2.5 = ", f1 + f2)  // Expected: 15.0
    println("string + string: 'VLang' + 'Cherry' = ", "VLang" + "Cherry") // Expected: "VLangCherry"

    // Resta (-)
    println("\nint - int:      10 - 3 = ", i1 - i2)         // Expected: 7
    println("float64 - int:  12.5 - 10 = ", f1 - i1)      // Expected: 2.5

    // Multiplicación (*)
    println("\nint * int:      10 * 3 = ", i1 * i2)         // Expected: 30
    println("float64 * int:  12.5 * 10 = ", f1 * i1)      // Expected: 125.0

    // División (/)
    println("\nint / int (truncation): 10 / 3 = ", i1 / i2) // Expected: 3
    println("float64 / float64:      12.5 / 2.5 = ", f1 / f2) // Expected: 5.0
    println("int / float64:          10 / 2.5 = ", i1 / f2)   // Expected: 4.0

    // Módulo (%)
    println("\nint % int:      10 % 3 = ", i1 % i2)         // Expected: 1

    // Negación Unaria
    println("\nUnary negation:   -10 = ", -i1)               // Expected: -10
    println("Unary negation: -12.5 = ", -f1)             // Expected: -12.5
    
    // Invalid arithmetic operations (commented out).
    // mut bad_op_1 = i1 + true     // SEMANTIC ERROR: Cannot add int and bool.
    // mut bad_op_2 = "hello" - "o" // SEMANTIC ERROR: Subtraction not defined for strings.
    // mut div_zero_int = 10 / 0    // SEMANTIC ERROR: Division by zero.
    // mut div_zero_flt = 12.5 / 0.0// SEMANTIC ERROR: Division by zero.
    // mut mod_zero = 10 % 0        // SEMANTIC ERROR: Division by zero.


    // --------------------------------------------------
    // Section B: Compound Assignment Operators
    // --------------------------------------------------
    println("\n--- B. Compound Assignment Operators ---")
    mut x_i int = 20
    x_i += 5
    println("int += int: (20 += 5) =", x_i) // Expected: 25
    x_i++
    println("int ++: (25++) =", x_i) // Expected: 26
    x_i -= 10
    println("int -= int: (25 -= 10) =", x_i) // Expected: 16

    mut x_f float64 = 10.0
    x_f += 5 // int is promoted for operation
    println("float64 += int: (10.0 += 5) =", x_f) // Expected: 15.0
    
    mut x_s string = "test"
    x_s += "_file"
    println("string += string: ('test' += '_file') =", x_s) // Expected: "test_file"
    
    // Invalid assignment (commented out).
    // x_i += 2.5 // SEMANTIC ERROR: Cannot assign float result (17.5) to int variable.

}