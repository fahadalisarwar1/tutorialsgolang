package main

import "fmt"

func main() {
	fmt.Println("Introduction to Loops in GoLang")
	// Loop Syntax
	// for initialisation; condition; post {
	// }
	for i := 1; i <= 10; i++ {

		// i will be initialized at the start of the loop
		// it will be 1 at start, then the control checks the condition,
		// if the condition is true i.e i=1<10 the control goes into loop body
		// it prints the variable valuee then once it has completed the execution it goes to post condition and
		// updates the value of the variable i and checks the condition and does the whole process again

		// break statement
		// it is used to terminate the loop abruptly without completing the loop
		//if i > 5 {
		//	break
		//}

		if i%2 != 0 {
			continue
		}
		// continue is used to skip the current loop
		fmt.Println(i)

	}
}
