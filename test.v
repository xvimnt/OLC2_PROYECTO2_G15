fn main() {

   mut i1 int = 10
    mut i2 int = 3
    mut f1 float64 = 12.5
    mut f2 float64 = 2.5
    
    // División (/)
    println("\nint / int (truncation): 10 / 3 = ", i1 / i2) // Expected: 3
    println("float64 / float64:      12.5 / 2.5 = ", f1 / f2) // Expected: 5.0
    println("int / float64:          10 / 2.5 = ", i1 / f2)   // Expected: 4.0
}