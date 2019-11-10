package main

import "fmt"

func main() {
	fmt.Println("Range tutorial")

	var arr = [5]int{11, 22, 33, 44, 55}

	for i, value := range arr { // range returns the index and value of the array

		fmt.Println(i, " ", value)
	}
	var arr2 = [5]string{"my", "name", "is", "fahad", "ali"}
	for i, value := range arr2 { // range returns the index and value of the array

		fmt.Println(i, " ", value)
	}

}
