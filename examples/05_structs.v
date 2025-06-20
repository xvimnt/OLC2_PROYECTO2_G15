// 05_structs.vch (Expanded and Corrected)
// ======================================
// Tests all features of structs as defined in the language specification.
// - All function declarations correctly use the `fn` keyword.
// - Structs are declared only in the global scope.
// - Struct instantiation using keyed elements: `Struct{FieldName: value}`.
// - Nested structs (a struct as a field of another struct).
// - Attribute access and modification.
// - Methods operating on structs (both read-only and modification by reference).
// - Structs used as function parameters and return types.
// - Printing structs to verify output format.

// --------------------------------------------------
// Section A: Global Struct Definitions
// --------------------------------------------------

// As per section 6.1, structs can only be declared in the global scope.
struct Point {
    int x
    int y
}

// A more complex struct with various primitive types.
struct Person {
    string name
    int    age
    bool   is_employed
}

// A nested struct, containing a Person struct.
struct Employee {
    Person personal_info
    string employee_id
    float64 salary
}

// An empty struct is not allowed (commented out to prevent syntax error).
// struct Empty {} // SYNTAX ERROR

// --------------------------------------------------
// Section B: Methods (Functions Associated with Structs)
// --------------------------------------------------

// Method with a value receiver (read-only).
fn (p Person) get_info() string {
    mut info := "Name: " + p.name + ", Age: " + p.age
    return info
}

// Method with a pointer receiver (modifies the struct).
// This demonstrates the "operan por referencia" concept from section 6.3.
fn (p *Person) have_birthday() {
    p.age = p.age + 1
}

// --------------------------------------------------
// Section C: Functions Handling Structs
// --------------------------------------------------

// Function that returns a struct instance.
fn create_employee(name string, age int, id string, salary float64) Employee {
    mut person := Person{
        name: name,
        age: age,
        is_employed: true
    }
    mut emp := Employee{
        personal_info: person,
        employee_id: id,
        salary: salary
    }
    return emp
}

// Function that modifies a struct passed by reference (pointer).
fn promote_employee(emp *Employee, new_salary float64) {
    emp.salary = new_salary
}

// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------
fn main() {
    println("--- Test 5: Structs, Methods, and Functions ---")
    
    // Struct definition inside a function is not allowed.
    // struct LocalStruct {} // SYNTAX ERROR
    
    // --------------------------------------------------
    // D. Instantiation and Attribute Access
    // --------------------------------------------------
    println("\n--- D. Instantiation and Attribute Access ---")
    
    // 6.2: Instantiating a simple struct
    mut p1 := Point{x: 10, y: 20}
    println("Initial Point:", p1) // Tests struct printing format
    
    // Accessing attributes
    println("p1.x =", p1.x) // Expected: 10
    
    // Modifying attributes
    p1.y = 25
    println("Modified p1.y =", p1.y) // Expected: 25

    // --------------------------------------------------
    // E. Using Methods
    // --------------------------------------------------
    println("\n--- E. Using Methods ---")
    mut person1 := Person{name: "Alice", age: 30, is_employed: true}
    println("Initial person info:", person1.get_info()) // Expected: "Name: Alice, Age: 30"
    
    person1.have_birthday() // Call method to modify struct
    println("Info after birthday:", person1.get_info()) // Expected: "Name: Alice, Age: 31"
    
    // --------------------------------------------------
    // F. Nested Structs and Functions with Structs
    // --------------------------------------------------
    println("\n--- F. Nested Structs and Functions with Structs ---")
    
    // Using a function that returns a complex, nested struct
    mut employee1 := create_employee("Bob", 42, "E1234", 75000.50)
    println("Created Employee:", employee1)
    
    // Accessing attributes of a nested struct
    println("Employee's name:", employee1.personal_info.name) // Expected: "Bob"
    println("Initial salary:", employee1.salary) // Expected: 75000.50

    // Using a function that modifies a struct by reference
    promote_employee(&employee1, 80000.0)
    println("Salary after promotion:", employee1.salary) // Expected: 80000.0

    println("\n--- Test 5 Finished ---")
}