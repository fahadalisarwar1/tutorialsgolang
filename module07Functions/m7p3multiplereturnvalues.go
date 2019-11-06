package main

import "fmt"

func main() {
	fmt.Println("Multiple Return values function tutorial")
	a := 33.44
	b := 64.33
	sum, diff := Calculator(a, b)
	fmt.Println("sum is ", sum)
	fmt.Println("diff is: ", diff)

}

// Calculator calculates the sum and difference of the two variables
func Calculator(num1 float64, num2 float64) (sum float64, diff float64) {
	sum = num1 + num2
	diff = num1 - num2
	return sum, diff
}
