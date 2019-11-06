package main

import "fmt"

var marks int // project scoped and initially assigned Zero

func main() {
	fmt.Println("Introduction to variables")
	var totalMarks = 100
	marks = 96
	gradeMarks := 59
	fmt.Printf("Marks are : %d marks out of total %d\n", marks, totalMarks)

	fmt.Println("Marks: ", marks, " Total Marks: ", totalMarks)
	fmt.Println(gradeMarks)

}

// totalMarks cant be accessed here since it is out of scope here
// marks can be accessed, since it is defined in the project scope..
