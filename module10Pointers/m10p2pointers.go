package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	// creating pointers using a new function

	ptr := new(int)
	fmt.Printf("value: %d\n", *ptr)
	fmt.Printf("Type: %T\n", ptr)

	fmt.Printf("address: %d\n", ptr)
	*ptr = 100
	fmt.Println(*ptr)

	// chainging the value of the poitner

	*ptr = 200
	fmt.Println(*ptr)
	fmt.Printf("%d ", ptr)
}
