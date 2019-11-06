package main

import "fmt"

func main() {
	fmt.Println("Conditional variant")
	if num := 9; num%2 == 0 {
		// one thing to note here, the variable "num is only available inside this conditional block"
		fmt.Println("number is even")
	} else { //else should start on the same line after finishing else otherwise it would raise an error.
		fmt.Println("number is odd")
	}

}
