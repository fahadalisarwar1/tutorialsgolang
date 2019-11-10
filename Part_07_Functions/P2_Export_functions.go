package main

import "fmt"

func main() {
	fmt.Println("Functions tutorial")
	sum := Sum(3.3, 64.5)
	fmt.Println(sum)
}

// Sum returns the sum of two numbers
func Sum(num1 float64, num2 float64) (sum float64) {
	sum = (num1 + num2)
	return
}

// example of non export function
