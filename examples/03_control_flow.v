// 03_control_flow.vch (Corrected)
// =================================
// Tests core control flow statements without relying on complex data structures.
// - All function declarations now correctly use the `fn` keyword.
// - If/Else-If/Else statements with simple and complex boolean conditions.
// - Nested If statements.
// - Switch statements on variables and expressions.
// - The two fundamental forms of For loops: condition-only (while) and C-style.
// - Nested For loops.
// - Transfer statements (break, continue, return) in their proper contexts.

// --------------------------------------------------
// Helper functions for testing
// --------------------------------------------------

// Helper to test various `return` scenarios.
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

// Helper to test `if` and complex conditions.
fn test_if_else() {
    println("\n--- A. Testing If/Else-If/Else ---")

    // Simple if-else
    mut temperature int = 25
    if temperature > 30 {
        println("It's hot.")
    } else {
        println("It's not hot.") // Expected
    }

    // if-else if-else chain
    mut time int = 14
    if time < 12 {
        println("Good morning.")
    } else if time < 18 {
        println("Good afternoon.") // Expected
    } else {
        println("Good evening.")
    }

    // Nested if statement with complex condition
    mut logged_in bool = true
    mut user_role string = "admin"
    if logged_in && (user_role == "admin" || user_role == "moderator") {
        println("Access level: Privileged") // Expected
        if user_role == "admin" {
            println("Admin panel access granted.") // Expected
        }
    }
}

// Helper to test `switch` variations.
fn test_switch() {
    println("\n--- B. Testing Switch ---")

    // Switch on a variable
    mut fruit string = "apple"
    switch fruit {
        case "banana":
            println("It's yellow.")
        case "apple":
            println("It's red or green.") // Expected. Implicit break prevents fall-through.
        case "orange":
            println("It's orange.")
        default:
            println("It's some other fruit.")
    }
    
    // Switch on an expression
    mut value int = 5
    switch value * 2 {
        case 10:
            println("Result is 10.") // Expected
        default:
            println("Result is not 10.")
    }
}

// Helper to test `for` loop variations (excluding range-based).
fn test_for_loops() {
    println("\n--- C. Testing For Loops ---")
    
    // Form 1: Condition-only loop (while-style)
    println("While-style loop with continue:")
    mut i int = 0
    for i < 5 {
        i = i + 1
        if i == 3 {
            continue // Skip printing 3
        }
        print(i) // Expected: 1 2 4 5
    }
    println("")

    // Form 2: C-style loop with a break
    println("C-style loop with break:")
    for mut j int = 0; j < 10; j = j + 1 {
        if j > 3 {
            break // Exit loop when j is 4
        }
        print(j) // Expected: 0 1 2 3
    }
    println("")
    
    // Nested loops: break on inner loop
    println("Nested loops:")
    for mut x int = 0; x < 3; x = x + 1 {
        print("Outer($x): ")
        for mut y int = 0; y < 5; y = y + 1 {
            if y == 3 {
                break // Breaks the *inner* loop only
            }
            print(y)
        }
        println("")
    }
}

// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------

fn main() {
    println("--- Test 3: Control Flow ---")
    
    test_if_else()
    test_switch()
    test_for_loops()

    println("\n--- D. Testing Return Statements ---")
    println("Score 95 is grade:", get_grade(95)) // Expected: A
    println("Score 82 is grade:", get_grade(82)) // Expected: B
    println("Score 50 is grade:", get_grade(50)) // Expected: C or below
    println("Score -10 is grade:", get_grade(-10)) // Expected: Invalid score (early return)

    println("\n--- E. Testing Invalid Control Flow (Commented Out) ---")
    // break      // SEMANTIC ERROR: break is not inside a loop or switch.
    // continue   // SEMANTIC ERROR: continue is not inside a loop.

    println("\n--- Test 3 Finished ---")
}