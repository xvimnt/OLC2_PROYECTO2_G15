fn get_grade(score int) string {
    if score < 0 || score > 100 {
        return "Invalid score" // Early return for invalid input.
    }
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    }
    // The final return acts as the default case.
    return "C or below"
}

fn main() {
    println("Grade for score 95:", get_grade(95))
} 