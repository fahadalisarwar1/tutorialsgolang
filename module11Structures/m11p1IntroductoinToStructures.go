package main

// A structure is a user defined type which represents a
//collection of fields. It can be used in places where it
//makes sense to group the data into a single unit rather
// than maintaining each of them as separate types

import "fmt"

// Student struct contains data about student
type Student struct {
	firstName string
	lastName  string
	marks     float64
	degree    string
	age       float64
}

func main() {
	fmt.Println("Structures tutorial")
	s1 := Student{
		firstName: "Fahad",
		lastName:  "Ali Sarwar",
		age:       25,
		degree:    "Bachelors in Electrical Engineering",
		marks:     80,
	}
	fmt.Println(s1)
	s2 := Student{
		firstName: "Joe",
		lastName:  "Tribbiani",
		age:       35,
		degree:    "degree in acting",
		marks:     25,
	}
	fmt.Println(s2)

	// you can also create a struct as follows

	var s3 Student
	s3.firstName = "Chandler"
	s3.lastName = "bing"
	s3.degree = "Accounting"
	s3.age = 35
	s3.marks = 20

}
