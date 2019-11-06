package main

import "fmt"

func main() {
	fmt.Println("Introduction to Pointers")
	// if you are coming from C++, you will be familiar with the concept of pointers in
	// GoLang also supports pointers.

	// here add a slice about pointers

	a := 64
	b := &a
	fmt.Println(a) // value of variable

	fmt.Println(b) // memory address of the variable
	// points to address in memory where the variable is stored

	// to find the value stored on that memory address

	fmt.Println(*b)
	// The Zero value of pointer is nil

	x := 25
	var ptr *int
	if ptr == nil {
		fmt.Println("Ptr is not pointing anywhere")
		ptr = &x
	}
	fmt.Println("Pointer is pointing to: ", ptr)
}
