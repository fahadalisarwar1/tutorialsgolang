package main

import "fmt"

func main() {
	fmt.Println("Variadic function tutorials")
	// A variadic function is a function that accepts a variable number of arguments.
	// If the last parameter of a function definition
	// is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.

	fmt.Println(summation(1, 2, 3, 4, 5))
	fmt.Println(summation(144, 55))

}

func summation(b ...float64) (sum float64) {
	sum = 0
	for _, j := range b {
		sum += j
	}
	return sum
}
