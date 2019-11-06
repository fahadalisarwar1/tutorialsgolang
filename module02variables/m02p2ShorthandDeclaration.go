package main

import "fmt"

func main() {
	fmt.Println("Short hand declaration")

	var name = "Fahad Ali" // type inference, when you initialize a variable, go automatically infers thhe data type
	secondName := "Sarwar"
	fmt.Println(name, " ", secondName)

	// multiple variable declaration

	dob, age := 1994, 26
	fmt.Println("Date of birth ", dob)
	fmt.Println("Age: ", age)
	isMarried := true
	fmt.Println(isMarried)

	// shorthand declaration with new and old variables

	isStupid, dob := true, 1900
	fmt.Println(isStupid, " ", dob)

	// isRich, salary = true, 100 // Error Since both of them were not declared before.

}
