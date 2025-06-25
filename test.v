fn main() {
    mut p bool = true
    mut q bool = false
    println("true && false:", p && q) // Expected: false
    println("true || false:", p || q) // Expected: true
    println("!true:", !p)             // Expected: false
    println("!false:", !q)            // Expected: true
}