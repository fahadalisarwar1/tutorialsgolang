package main

import "fmt"

// function is a block of code that performs a certain function
// func  keyword is used to define a function

func main() {

	fmt.Println("Functions Tutorials")
	name := "Fahad"

	greetings := greetUser(name, "ali sarwar")
	fmt.Println(greetings)
	a := 22.33
	b := 44.55
	mul, div := multiplyAndDivide(a, b)
	fmt.Println(mul, div)
}

// paramters to be passed to a function are writted inside the brackets
// the return type is written after the paranthesis
// go requires you to specify the return type of the function

// if the parameters are of the same type we can only write once which type these variables are as shown below

func greetUser(name, secondName string) string {
	return "Welcome to the Go Lang Tutorials " + name + " " + secondName
}

func multiplyAndDivide(num1, num2 float64) (float64, float64) {
	return num1 * num2, num1 / num2
}
