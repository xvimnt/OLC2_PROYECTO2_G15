fn main() {
    mut fruit string = "jurassic"
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
} 