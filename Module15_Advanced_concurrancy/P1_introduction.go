package main

import (
	"fmt"
)

func printNumbers(num int){
	for i := 0; i <= num; i++{
		fmt.Println(i)
	}
}
func main(){
	go printNumbers(10)

	// Next Step add wait for user input
	var input string
	fmt.Scanln(&input)
}
