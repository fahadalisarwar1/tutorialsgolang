package main

// channels are pipes through which go routines communicate
// with each other
// data can be sent from one go routine to another

// Declaring Channels

// each channel has a type associated with it.
// channel can only transport predefined data types
// water pipe can't transport oil

import "fmt"

func main(){

	chanVar := make(chan int)
	fmt.Printf("%T", chanVar)
	fmt.Println()
}