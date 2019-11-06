package main

import (
	"fmt"
	"time"
)

func printNumbers(){
	fmt.Println("Go Routine Print Numbers Started")
	for i := 1; i <= 5 ; i++{
		// 1   2   3   4   5
		// 1   2   3   4   5
		time.Sleep(1*time.Second)
		fmt.Printf("%d ", i)
	}
	fmt.Println("numbers printed after 5 seconds")
}

func printAlphabets(){
	fmt.Println("GO Routine print alphabets started")
	// a   b   c   d   e
	// 2   4   6   8   10
	for i := 'a' ; i <= 'e' ; i ++{
	time.Sleep(2*time.Second)
	fmt.Printf("%c ", i)
	}
	fmt.Println("numbers printed after 10 seconds")
}

func main(){
	go printNumbers()
	go printAlphabets()


	time.Sleep(12*time.Second)
	fmt.Println("\nMain GO routine terminated after 12 seconds in total")
}
