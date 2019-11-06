package main

import (
	"fmt"
	"time"
)

// Every go routine should handle the sending and receiving of data properly,
// suppose if a go routine is sending data but the other go routine is not receiving the data, then a deadlock is
// occuring, similarly if a go routine is waiting to receive data and the other go routine doesn't send data then we
// say that a deadlock has occurred.

func sayHi(){
	fmt.Println("Hi in go routine")
	time.Sleep(2*time.Second)
}

func main(){
	chanVar := make(chan int)
	go sayHi()
	<- chanVar // we assume that go routine will send data but go routine is not sending data so an error will
	// be caused by the dead lock

	fmt.Println("Exiting ")
}