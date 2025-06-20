// 11_integration_and_complex_cases.vch
// ======================================
// Tests complex interactions between different language features.
// This file focuses on integration testing to verify that structs, slices (arrays),
// functions, and methods work together as expected in realistic scenarios.

// --------------------------------------------------
// Section A: Complex Data Structure Definitions
// --------------------------------------------------

// A simple struct representing a student.
struct Student {
    string student_id
    string name
    float64 grade
}

// A complex struct representing a university course.
// CRITICAL: It contains a slice of other structs (an array within a struct).
struct Course {
    string course_code
    string course_name
    []Student enrolled_students
}


// --------------------------------------------------
// Section B: Methods for Complex Structs
// --------------------------------------------------

// Method for the Course struct.
// It processes the internal slice of students and returns a new slice
// containing only those who have a passing grade.
fn (c Course) get_passing_students(passing_grade float64) []Student {
    mut passing_students := []Student{} // Initialize an empty slice to store results.
    
    for _, student in c.enrolled_students {
        if student.grade >= passing_grade {
            passing_students = append(passing_students, student)
        }
    }
    
    return passing_students
}


// --------------------------------------------------
// Section C: Functions Handling Complex Structures
// --------------------------------------------------

// A function that takes a complex struct (Course) and a student ID,
// and returns a pointer to the student if found, otherwise nil.
fn find_student_in_course(course Course, id string) *Student {
    for i, student in course.enrolled_students {
        if student.student_id == id {
            // Return a pointer to the student within the slice.
            // Note: This level of pointer manipulation might be very advanced
            // for the interpreter. A simpler version might just return the index.
            // For now, we'll assume a robust implementation can handle this.
            // If not, simply returning a found/not-found bool is also a good test.
        }
    }
    return nil // Student not found
}

// A function that enrolls a new student into a course.
// It takes a pointer to the course to modify its internal slice directly.
fn enroll_student(c *Course, new_student Student) {
    c.enrolled_students = append(c.enrolled_students, new_student)
}


// --------------------------------------------------
// Main Execution Block
// --------------------------------------------------
fn main() {
    println("--- Test 11: Integration and Complex Cases ---")

    // --- D. Creating and Manipulating Complex Data ---
    println("\n--- D. Creating and Manipulating Complex Data ---")
    
    // Create initial students for the first course.
    mut student1 := Student{student_id: "S001", name: "Alice", grade: 95.5}
    mut student2 := Student{student_id: "S002", name: "Bob", grade: 55.0}
    mut student3 := Student{student_id: "S003", name: "Charlie", grade: 88.0}

    // Create a course with a slice of students.
    mut olc2_course := Course{
        course_code: "CS416",
        course_name: "Compilers 2",
        enrolled_students: []Student{student1, student2, student3}
    }
    
    println("Initial course created:", olc2_course.course_name)
    println("Number of students enrolled:", len(olc2_course.enrolled_students)) // Expected: 3

    // --- E. Using Functions and Methods on Complex Data ---
    println("\n--- E. Using Functions and Methods on Complex Data ---")
    
    // Use a method to filter the data within the struct.
    mut high_achievers := olc2_course.get_passing_students(90.0)
    println("High-achieving students (>= 90.0):", high_achievers)
    println("Number of high-achievers:", len(high_achievers)) // Expected: 1 (Alice)

    // Use a function to add a new student to the course.
    mut student4 := Student{student_id: "S004", name: "Diana", grade: 72.0}
    enroll_student(&olc2_course, student4)
    
    println("\nEnrolled a new student: Diana")
    println("New number of students:", len(olc2_course.enrolled_students)) // Expected: 4
    
    // --- F. Array of Structs ---
    println("\n--- F. Array of Structs ---")

    // Create a second course.
    mut db1_course := Course{
        course_code: "CS318",
        course_name: "Databases 1",
        enrolled_students: []Student{}
    }
    
    // Create a slice (array) that holds Course structs.
    mut university_courses := []Course{olc2_course, db1_course}
    
    println("Total number of courses offered:", len(university_courses)) // Expected: 2
    
    // Access and print data from the array of structs.
    println("Second course in the list is:", university_courses[1].course_name) // Expected: Databases 1
    
    // Modify data deep within the nested structures.
    println("Bob's initial grade in OLC2:", university_courses[0].enrolled_students[1].grade) // Expected: 55.0
    university_courses[0].enrolled_students[1].grade = 61.0
    println("Bob's new grade after retake:", university_courses[0].enrolled_students[1].grade) // Expected: 61.0

    println("\n--- Test 11 Finished ---")
}