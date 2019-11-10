package main

import "fmt"

func main() {
	fmt.Println("Pointer Arthematic")
	// Go doesn't support pointer arthematic
	arr := [3]int{1, 2, 3}
	ptr := &arr
	ptr++ // cant work since it holds memory address not value
}
