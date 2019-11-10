package main

// don't pass pointer to arrays to a function

import "fmt"

func main() {

	fmt.Println("pointer to an array")
	arr1 := [3]int{20, 90, 40}
	fmt.Println(arr1)
	changeArray(&arr1)
	fmt.Println(arr1)

	// passing sli
}

// although this works perfectly,, this is not an ideal way of doing thigns in Go

func changeArray(arr *[3]int) {
	(*arr)[1] = 30 // (*arr[0]) is same as arr[0]

}
