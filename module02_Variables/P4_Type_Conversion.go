package main

import "fmt"

func main() {

	fmt.Println("type conversion tutorial")

	var (
		num1 int32
		num2 int64
	)
	num1 = 34
	num2 = 45

	fmt.Println(int64(num1) + num2)
}
