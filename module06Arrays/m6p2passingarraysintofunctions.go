package main

import "fmt"

func main() {
	fmt.Println("Passing arrays into functions")
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Old array: ", arr)

	arr = modifyArray(arr)

	fmt.Println("new array: ", arr)

}

func modifyArray(array [5]int) (arr [5]int) {
	arr = array
	arr[2] = 33
	return arr
}
