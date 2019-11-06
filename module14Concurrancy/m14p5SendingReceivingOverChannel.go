package main

// how to send and receive data over channel
// data := <- chanVar read from channel

// chanVar <- data write to channel variable

// <- shows the direction of the channel

// send and receives are blocking by default which means until the data is
// read in the other routine the exeuction will pause.
import "fmt"

func sayHello(data chan bool){
	fmt.Println("sayHello Go routine")
	data <- true
}
func main(){

	done := make(chan bool)
	go sayHello(done)
	x := <- done
	fmt.Println(x)

	fmt.Println("Main Go routine")
}