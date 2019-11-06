package main

import "fmt"

func main() {
	fmt.Println("Slices tutorials")
	// slices don't own any data they only reference to existing arrays
	arr := [5]int{100, 200, 300, 400, 500}
	slice := arr[2:4]
	//array indexes start at 0
	// [2:4] refers to 2 inclusive and 4 exclusive thats why only two values are present
	// slices contain references, this means if you modify slice original will also be modified

	fmt.Println(slice)
	slice[1] = 3

	fmt.Println(len(slice))
	fmt.Println(len(arr))

	fmt.Println(slice)
	fmt.Println(arr)
}
