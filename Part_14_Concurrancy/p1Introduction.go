package main

import (
	"fmt"
	"time"
)

//  Lets start coding


func sayHello(){
	fmt.Println("hello user, I'm running a GO routine!!")
}

func main(){
	// main go routine is already running by default when we start the program, we dont have to handle it.

	// When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not
	// wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the
	// Goroutine call and any return values from the Goroutine are ignored.
	//The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the
	// program will be terminated and no other Goroutine will run.
	go sayHello()
	// after exeuting this line the control runs back to the main program instead of showing the output for the sayHello
	// function
	// Now add timer
	time.Sleep(1*time.Second) // Now the go routine sayHello has enough time to execute and show its output
	// This sleep hack is used to undertand the working of a go routine.
	// In practical cases channels are used to block executions to let go routines finish properly

	fmt.Println("Concurrency in GO")




}
