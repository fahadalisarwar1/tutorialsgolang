package main

// Iin this tutorial we will work on passing pointer to a
// function

import "fmt"

func main() {

	fmt.Println("Passing pointers to a function")
	num1 := 100
	fmt.Println("The Value before passing pointer: ", num1)
	passingPointer(&num1) // passing the address of the varibale
	fmt.Println("The Value after passing pointer: ", num1)

}

func passingPointer(ptr *int) {
	*ptr = 34 // changing the value of the address pointed by the ptr
}
