package main

import "fmt"

func main() {
	fmt.Println("Anonymous Struct")
	// so far we have created a struct with explicit name called Student
	// it is not always necessary to create a struct with a name

	employee := struct {
		name string
		age  float64
	}{
		name: "Fahad",
		age:  25,
	}
	fmt.Println(employee)
}
