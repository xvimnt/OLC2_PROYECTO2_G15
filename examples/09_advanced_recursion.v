// 10_advanced_recursion.vch
// ==================================
// Tests advanced and computationally intensive recursive algorithms.
// This file is a stress test for the interpreter's call stack efficiency,
// parameter handling, and overall performance.

// --------------------------------------------------
// Section A: Fibonacci Sequence
// --------------------------------------------------

// Calculates the nth number in the Fibonacci sequence.
// This function has two recursive calls, leading to exponential growth.
fn fibonacci(n int) int {
    if n <= 1 {
        return n // Base cases: fib(0) = 0, fib(1) = 1
    }
    // Two recursive calls per step
    return fibonacci(n - 1) + fibonacci(n - 2)
}


// --------------------------------------------------
// Section B: Towers of Hanoi
// --------------------------------------------------

// Solves the Towers of Hanoi puzzle by printing the moves.
// Parameters:
//   n: number of disks
//   from_rod: the starting rod
//   to_rod: the destination rod
//   aux_rod: the auxiliary rod
fn towers_of_hanoi(n int, from_rod string, to_rod string, aux_rod string) {
    if n == 1 {
        // Base case: move the smallest disk
        println("Move disk 1 from rod", from_rod, "to rod", to_rod)
        return
    }
    // Move n-1 disks from source to auxiliary, using destination as auxiliary
    towers_of_hanoi(n - 1, from_rod, aux_rod, to_rod)
    
    // Move the nth disk from source to destination
    println("Move disk", n, "from rod", from_rod, "to rod", to_rod)
    
    // Move the n-1 disks from auxiliary to destination, using source as auxiliary
    towers_of_hanoi(n - 1, aux_rod, to_rod, from_rod)
}


// --------------------------------------------------
// Section C: Ackermann Function (Stress Test)
// --------------------------------------------------

// The Ackermann function is a classic example of a total computable function
// that is not primitive recursive. It grows extremely fast.
// WARNING: Do not use large inputs! A(3, 4) is a common benchmark but may be
// too slow or cause a stack overflow in a simple interpreter.
fn ackermann(m int, n int) int {
    if m == 0 {
        return n + 1
    }
    if n == 0 {
        return ackermann(m - 1, 1) // Recursive call
    }
    // Double recursive call
    return ackermann(m - 1, ackermann(m, n - 1))
}


// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------
fn main() {
    println("--- Test 10: Advanced Recursive Algorithms ---")

    // --- Testing Fibonacci ---
    println("\n--- A. Fibonacci Sequence ---")
    
    mut fib_10 := fibonacci(10)
    println("Fibonacci(10) is:", fib_10) // Expected: 55
    
    mut fib_7 := fibonacci(7)
    println("Fibonacci(7) is:", fib_7)   // Expected: 13
    
    println("-----------------------------------")

    // --- Testing Towers of Hanoi ---
    println("\n--- B. Towers of Hanoi (n=3) ---")
    // Expected output for n=3:
    // Move disk 1 from rod A to rod C
    // Move disk 2 from rod A to rod B
    // Move disk 1 from rod C to rod B
    // Move disk 3 from rod A to rod C
    // Move disk 1 from rod B to rod A
    // Move disk 2 from rod B to rod C
    // Move disk 1 from rod A to rod C
    towers_of_hanoi(3, "A", "C", "B")
    
    println("-----------------------------------")
    
    // --- Testing Ackermann Function ---
    // These are safe, small values that should execute quickly.
    println("\n--- C. Ackermann Function (Small Values) ---")

    mut ack_1_2 := ackermann(1, 2)
    println("Ackermann(1, 2) is:", ack_1_2) // Expected: 4

    mut ack_2_1 := ackermann(2, 1)
    println("Ackermann(2, 1) is:", ack_2_1) // Expected: 5

    mut ack_2_3 := ackermann(2, 3)
    println("Ackermann(2, 3) is:", ack_2_3) // Expected: 9

    // --- STRESS TEST ---
    // This value is larger and will test the stack more intensely.
    // A correct but unoptimized interpreter might take a few seconds here.
    // If it crashes, it indicates a stack overflow problem.
    println("\n--- Ackermann Stress Test ---")
    mut ack_3_2 := ackermann(3, 2)
    println("Ackermann(3, 2) is:", ack_3_2) // Expected: 29
    
    // --- DANGEROUS VALUES (DO NOT RUN UNLESS OPTIMIZED) ---
    // ackermann(3, 3) -> 61
    // ackermann(3, 4) -> 125
    // ackermann(4, 1) -> 65533
    // These will almost certainly cause a stack overflow in a basic interpreter.
    
    println("\n--- Test 10 Finished ---")
}