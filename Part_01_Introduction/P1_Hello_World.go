package main
// This is known as a package declaration, and every Go program must start with it.
//Packages are Go’s way of organizing and reusing code. There are two types of Go pro‐
//grams: executables and libraries. Executable applications are the kinds of programs
//that we can run directly from the terminal (on Windows, they end with .exe). Libra‐
//ries are collections of code that we package together so that we can use them in other
//programs. We will explore libraries in more detail later; for now, just make sure to
//include this line in any program you writ


import "fmt"
// import keyword is used for importing libraries
// "" means string literals

// This is single line comment

/*
This is multiline comment
*/

func main() {
	// functions are building blocks of go

	// This block contains code run by the main program....

	fmt.Println("Hello World, Welcome to GoLang Programming")
	// Println function prints the text to the standard output, in this case the terminal,.....

}
