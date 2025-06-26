fn main() {
    mut i int = 0
    for i < 5 {
        i = i + 1
        if i == 3 {
            continue // Skip printing 3
        }
        print(i) // Expected: 1 2 4 5
    }
} 