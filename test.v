fn main() {
    mut logged_in bool = true
    mut user_role string = "admin"
    if logged_in && (user_role == "admin" || user_role == "moderator") {
        println("Access level: Privileged") // Expected
        if user_role == "admin" {
            println("Admin panel access granted.") // Expected
        }
    }
} 