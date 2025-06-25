fn main() {
    mut uninitialized_int int
    mut uninitialized_float float64
    mut uninitialized_string string
    mut uninitialized_bool bool
    println("Default int value:", uninitialized_int)       // Expected: 0
    println("Default float64 value:", uninitialized_float) // Expected: 0.0
    println("Default string value: '$uninitialized_string'") // Expected: "" (empty string)
    println("Default bool value:", uninitialized_bool)     // Expected: false

}