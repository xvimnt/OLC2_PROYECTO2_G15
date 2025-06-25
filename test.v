fn main() {
    println("string == string: ('a' == 'a') is", "a" == "a")       // Expected: true
    println("string != string: ('a' != 'A') is", "a" != "A")       // Expected: true
    println("bool == bool: (true == true) is", true == true)       // Expected: true
    println("rune == rune: ('z' == 'z') is", 'z' == 'z')           // Expected: true
    println("float == int: (10.0 == 10) is", 10.0 == 10)         // Expected: true
    println("int != float: (10 != 10.1) is", 10 != 10.1)         // Expected: true
}