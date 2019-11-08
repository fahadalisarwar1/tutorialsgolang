package main

import "fmt"

func main() {

	fmt.Println("Switch tutorial")
	option := 1
	switch option {
	case 1:
		fmt.Println("go running")
	case 2:
		fmt.Println("go jogging")
	case 3:
		fmt.Println("do squats")
	case 4:
		fmt.Println("do push-ups")
	default:
		fmt.Println("invalid exercise entnered")
	}

	// multiple switch cases merged together
	alphabet := 'z'
	switch alphabet {
	// case 'a', 'e", "i", "o", "u":
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("A vowel is entered")
	default:
		fmt.Println("Not a vowel")
	}
}
