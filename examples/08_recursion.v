// 09_recursion.vch
// ==================================
// Tests the interpreter's ability to handle recursive functions and methods.
// This is a critical test for the function call stack, parameter passing,
// and scope management during recursion.

// --------------------------------------------------
// Section A: Direct Recursive Function - Factorial
// --------------------------------------------------

// A classic recursive function to calculate factorial.
// Base Case: factorial(0) is 1.
// Recursive Step: n * factorial(n - 1).
fn factorial(n int) int {
    if n == 0 {
        return 1 // Base case
    }
    // Recursive call
    return n * factorial(n - 1)
}


// --------------------------------------------------
// Section B: Mutually Recursive Functions - Even/Odd
// --------------------------------------------------

// Two functions that call each other to determine if a number is even or odd.
// This tests the handling of a more complex call stack.

fn is_even(n int) bool {
    if n == 0 {
        return true // Base case
    }
    // Call the other function
    return is_odd(n - 1)
}

fn is_odd(n int) bool {
    if n == 0 {
        return false // Base case
    }
    // Call the other function
    return is_even(n - 1)
}


// --------------------------------------------------
// Section C: Recursive Method on a Struct - Linked List
// --------------------------------------------------

// A struct to represent a node in a simple linked list.
// The `next` field can point to another Node, creating the recursive structure.
struct Node {
    int value
    *Node next // Pointer to the next node in the list
}

// A recursive method to sum all values in the linked list starting from the current node.
fn (n *Node) sum() int {
    if n.next == nil {
        return n.value // Base case: last node in the list
    }
    // Recursive call on the next node in the chain
    return n.value + n.next.sum()
}


// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------
fn main() {
    println("--- Test 9: Recursion ---")

    // --- Testing Direct Recursion ---
    println("\n--- A. Direct Recursion (Factorial) ---")
    
    mut result_fact_5 := factorial(5) // 5 * 4 * 3 * 2 * 1 = 120
    println("Factorial of 5 is:", result_fact_5) // Expected: 120
    
    mut result_fact_0 := factorial(0) // Base case
    println("Factorial of 0 is:", result_fact_0) // Expected: 1
    
    mut result_fact_1 := factorial(1)
    println("Factorial of 1 is:", result_fact_1) // Expected: 1
    
    println("-----------------------------------")

    // --- Testing Mutual Recursion ---
    println("\n--- B. Mutual Recursion (Even/Odd) ---")
    
    mut even_check_10 := is_even(10)
    println("Is 10 even?", even_check_10) // Expected: true
    
    mut odd_check_7 := is_odd(7)
    println("Is 7 odd?", odd_check_7) // Expected: true

    mut even_check_7 := is_even(7)
    println("Is 7 even?", even_check_7) // Expected: false
    
    println("-----------------------------------")
    
    // --- Testing Recursive Method ---
    println("\n--- C. Recursive Method (Linked List Sum) ---")

    // Manually create a linked list: 10 -> 20 -> 30
    mut node3 := Node{value: 30, next: nil}
    mut node2 := Node{value: 20, next: &node3}
    mut node1 := Node{value: 10, next: &node2}

    // Call the recursive method on the head of the list
    mut total_sum := node1.sum() // 10 + (20 + 30) = 60
    println("Sum of linked list (10->20->30) is:", total_sum) // Expected: 60

    // Test the base case with a single-node list
    mut single_node := Node{value: 100, next: nil}
    println("Sum of single node list is:", single_node.sum()) // Expected: 100

    println("\n--- Test 9 Finished ---")
}// 09_recursion.vch
// ==================================
// Tests the interpreter's ability to handle recursive functions and methods.
// This is a critical test for the function call stack, parameter passing,
// and scope management during recursion.

// --------------------------------------------------
// Section A: Direct Recursive Function - Factorial
// --------------------------------------------------

// A classic recursive function to calculate factorial.
// Base Case: factorial(0) is 1.
// Recursive Step: n * factorial(n - 1).
fn factorial(n int) int {
    if n == 0 {
        return 1 // Base case
    }
    // Recursive call
    return n * factorial(n - 1)
}


// --------------------------------------------------
// Section B: Mutually Recursive Functions - Even/Odd
// --------------------------------------------------

// Two functions that call each other to determine if a number is even or odd.
// This tests the handling of a more complex call stack.

fn is_even(n int) bool {
    if n == 0 {
        return true // Base case
    }
    // Call the other function
    return is_odd(n - 1)
}

fn is_odd(n int) bool {
    if n == 0 {
        return false // Base case
    }
    // Call the other function
    return is_even(n - 1)
}


// --------------------------------------------------
// Section C: Recursive Method on a Struct - Linked List
// --------------------------------------------------

// A struct to represent a node in a simple linked list.
// The `next` field can point to another Node, creating the recursive structure.
struct Node {
    int value
    *Node next // Pointer to the next node in the list
}

// A recursive method to sum all values in the linked list starting from the current node.
fn (n *Node) sum() int {
    if n.next == nil {
        return n.value // Base case: last node in the list
    }
    // Recursive call on the next node in the chain
    return n.value + n.next.sum()
}


// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------
fn main() {
    println("--- Test 9: Recursion ---")

    // --- Testing Direct Recursion ---
    println("\n--- A. Direct Recursion (Factorial) ---")
    
    mut result_fact_5 := factorial(5) // 5 * 4 * 3 * 2 * 1 = 120
    println("Factorial of 5 is:", result_fact_5) // Expected: 120
    
    mut result_fact_0 := factorial(0) // Base case
    println("Factorial of 0 is:", result_fact_0) // Expected: 1
    
    mut result_fact_1 := factorial(1)
    println("Factorial of 1 is:", result_fact_1) // Expected: 1
    
    println("-----------------------------------")

    // --- Testing Mutual Recursion ---
    println("\n--- B. Mutual Recursion (Even/Odd) ---")
    
    mut even_check_10 := is_even(10)
    println("Is 10 even?", even_check_10) // Expected: true
    
    mut odd_check_7 := is_odd(7)
    println("Is 7 odd?", odd_check_7) // Expected: true

    mut even_check_7 := is_even(7)
    println("Is 7 even?", even_check_7) // Expected: false
    
    println("-----------------------------------")
    
    // --- Testing Recursive Method ---
    println("\n--- C. Recursive Method (Linked List Sum) ---")

    // Manually create a linked list: 10 -> 20 -> 30
    mut node3 := Node{value: 30, next: nil}
    mut node2 := Node{value: 20, next: &node3}
    mut node1 := Node{value: 10, next: &node2}

    // Call the recursive method on the head of the list
    mut total_sum := node1.sum() // 10 + (20 + 30) = 60
    println("Sum of linked list (10->20->30) is:", total_sum) // Expected: 60

    // Test the base case with a single-node list
    mut single_node := Node{value: 100, next: nil}
    println("Sum of single node list is:", single_node.sum()) // Expected: 100

    println("\n--- Test 9 Finished ---")
}