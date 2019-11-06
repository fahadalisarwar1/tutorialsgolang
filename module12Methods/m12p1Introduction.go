package main

import (
	"computer"
	"fmt"
)

// A method is just a function with a special receiver type between the fun
// keyword and the method name. The receiver can either be a struct type or non-
// struct type

func main() {
	fmt.Println("Methods Introduction")
	s1 := computer.ProcessingUnit{
		Maker: "intel",
		MHz:   1200,
		Price: 500,
	}
	s1.PrintSummary()
}
