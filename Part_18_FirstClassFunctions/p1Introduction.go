package main

import "fmt"

// A language which supports first class functions allows functions to be assigned to variables, passed as
// arguments to other functions and returned from other functions. Go has support for first class functions.


func main(){
	funcVar := func(){
		fmt.Println("First class function")
	}
	// The only way to call this function is using the variable a. We have done this in the next line. a() calls the
	// function and this prints hello world first class function. In line no. 12 we print variable a's type . This will print func().
	funcVar()
	fmt.Printf("%T\n", funcVar)
}