fn main() {
    mut prec_result := 5 * 2 + 3 > 12 && !false // (10 + 3) > 12 && true -> 13 > 12 && true -> true && true -> true
    println("5 * 2 + 3 > 12 && !false is", prec_result) // Expected: true
}