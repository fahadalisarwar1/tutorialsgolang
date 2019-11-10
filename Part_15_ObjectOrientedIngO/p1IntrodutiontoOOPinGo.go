package main

import (
	"fmt"
	"gopackages/student"
)

func main() {
	fmt.Println("Main Program")
	s1 := student.Student{
		Name:       "Fahad ",
		LastName:   "Ali",
		Marks:      340,
		TotalMarks: 500,
		Age:        33,
	}
	s1.PercentageMarks()

	var s2 student.Student
	s2.PercentageMarks() // this doesnt has valid results in print function
	// the problem can be solved using constructors

}
