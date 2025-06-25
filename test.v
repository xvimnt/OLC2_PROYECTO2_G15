fn main() {
   mut x_f float64 = 10.0
    x_f-- // int is promoted for operation
    println("float64 --: (10.0 --) =", x_f) // Expected: 9.0
}   