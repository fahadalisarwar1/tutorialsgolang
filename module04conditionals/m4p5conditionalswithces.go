package main

import "fmt"

func main() {
	fmt.Println("conditional swithces")
	num1 := 50
	num2 := 100
	switch {
	case num1 == 50 && num2 < 100:
		fmt.Println("Number is 50")
	case num1 < 50:
		fmt.Println("number is less than 50")
	case num1 == 50 && num2 >= 50:
		fmt.Println("number is 50 and number2 is more than 50")
	}
}
