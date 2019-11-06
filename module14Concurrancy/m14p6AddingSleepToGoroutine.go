package main

// Now we will introduce some sleep timers to this method to have a better understanding

import (
	"fmt"
	"time"
)

// Hello
func Hello(chanVar chan bool){
	fmt.Println("Go routine before sleep")
	time.Sleep(4*time.Second) // during this time main go routine will be blocked and wont be used. s
	fmt.Println("Go routine after sleep and writing to the chanVar")
	chanVar <- true

}

func main(){
	fmt.Println("Adding Sleep in the go routine")
	data := make(chan bool)
	fmt.Println("calling the go routine")
	go  Hello(data)

	<-data
	fmt.Println("Main GO routine executing")



}
