package main

import "fmt"

// Employee data
type Employee struct {
	name string
	age  int64
}

func main() {
	fmt.Println("Pointers to a struct")

	emp := &Employee{"john", 25}
	// emp is pointer to the Employee struct
	fmt.Println("name: ", (*emp).name, " Age: ", (*emp).age)
	// it can also be accessed directly
	fmt.Println("name: ", emp.name, " Age: ", emp.age)

}
