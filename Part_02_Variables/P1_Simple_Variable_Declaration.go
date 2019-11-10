package main

import "fmt"

// Topics to be covered
// -----------------------

// -----------------------

// int {8, 16, 32, 64}
// float {32, 64}
// bool
// strings



var gpa int // project scoped and initially assigned Zero
var studentName string
var isGraduated bool // No assignment should be made here.

func main() {
	gpa = 3 // int
	studentName = "Fahad Ali"
	isGraduated = true

	fmt.Println("Student Name: ", studentName, "\nMarks: ", gpa, "\nis Graduated: ", isGraduated)

}
