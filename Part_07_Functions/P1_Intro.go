package main

import "fmt"

func main() {
	fmt.Println("Functions tutorial")
	av := average(3.3, 64.5)
	fmt.Println(av)
}

// example of non export function
func average(num1 float64, num2 float64) (average float64) {
	average = (num1 + num2) / 2
	return
}
