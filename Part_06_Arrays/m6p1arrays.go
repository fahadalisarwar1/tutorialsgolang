package main

import "fmt"

func main() {
	fmt.Println("Arrays and slices")
	// An array is a collection of elements that belong to the same type.
	// For example the collection of integers 5, 8, 9, 79, 76 form an array. Mixing values of different types,
	// for example an array that contains both strings and integers is not allowed in Go.

	// array declaration

	var arr [5]int

	arr[0] = 33
	arr[1] = 44
	arr[2] = arr[0]
	arr[3] = arr[2]
	arr[4] = arr[1]
	// arr[5] = 3
	// fmt.Println(arr)
	arr1 := [...]int{2, 3, 4} // ... means compiler finds the length oof the array from the data given
	// fmt.Println(arr1)

	// NOTE: The size of the array is a part of the type
	// hence [5]int and [10]int are totally different types

	// TODO
	// arrays are valued types this means copying an array to another will not modify the contents of the orginal array
	b := arr1
	b[1] = 333
	fmt.Println("original array ", arr1)
	fmt.Println("copied array ", b)
	students := [...]string{"fahad", "ali", "sarwar"}
	newStudents := students
	newStudents[1] = "usman"
	fmt.Println("Original array: ", students)
	fmt.Println("New array: ", newStudents)

	var emptyArr []string
	if emptyArr == nil {
		fmt.Println("string array is nil")
	}

}
