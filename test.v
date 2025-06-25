fn main() {
   println("\n--- B. Compound Assignment Operators ---")
    mut x_i int = 20
    x_i += 5
    println("int += int: (20 += 5) =", x_i) // Expected: 25
    x_i++
    println("int ++: (25++) =", x_i) // Expected: 26
    x_i -= 10
    println("int -= int: (25 -= 10) =", x_i) // Expected: 15
}   