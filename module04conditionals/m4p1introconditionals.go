package main

import "fmt"

func main() {
	fmt.Println("conditionals introduction")
	num1 := 25
	num2 := 50
	if num1 > num2 {
		fmt.Printf("%d is more than %d\n", num1, num2)
	} else if num1 < 125 {
		fmt.Println(num1, " is less than 125")

	}

	// equality operator
	num1 = 20
	num2 = 20

	if num1 == num2 {
		fmt.Println("Both numbers are equal")

	}

	num1 = 12
	if num1 != num2 {
		fmt.Println("Both numbers are not equal")
	}
}
