package main

import "fmt"

func main() {
	fmt.Println("Passing Slices")
	arr := [3]int{20, 30, 40} // declaring an array
	fmt.Println("Before")
	fmt.Println(arr) // printing an initial array

	passingSlices(arr[:]) // passing slice to an array, since they are referenced by types not values

	fmt.Println("After")
	fmt.Println(arr)
}

func passingSlices(slc []int) {
	slc[1] = 99
}
