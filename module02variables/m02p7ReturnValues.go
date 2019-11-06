package main

import "fmt"

/*It is possible to return named values from a function. If a return value is named,
it can be considered as being declared as a variable in the first line of the function.*/

func main() {
	fmt.Println("Named Return Values")

	lenght := 4
	width := 4
	aa := getArea(lenght, width)
	fmt.Println(aa)
	// lets say we receive aa from the function but we dont want to print it, we can use a blank identifier to
	// ignore the aa to get rid of "not used" error

}

func getArea(length, width int) (area int) {
	area = length * width
	return // here we didnt return anything but we already specified which variable we want to return, golang automatically
	// returns that variable
}
