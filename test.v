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
}