package main

import "fmt"

// All the previous go routines were bidirectional this means that data can flow in both directions
// this is not the case always, some go routine can only transfer data in one direction only

// Use of Unidirectional Channel: The unidirectional channel is used to provide the type-safety of the program so,
// that the program gives less error.
// Or you can also use a unidirectional channel when you want to create a channel that can only send or receive data.

func sayGreetings(chanVar chan <- int ){  // The arrow indicates the unidirectional go routine
	fmt.Println("SayGreetings Go Routine")
	chanVar <- 100
}


func main(){
	fmt.Println("Main go routine")

	chanVar := make(chan<- int)
	go sayGreetings(chanVar)
	fmt.Println(<-chanVar)// This is not valid

}