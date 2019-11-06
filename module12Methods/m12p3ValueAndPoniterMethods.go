package main

// There are two types of methods,
// we have already studied the
// value type methods

import (
	"computer"
	"fmt"
)

func main() {
	fmt.Println("Pointer type methods")
	c1 := computer.ProcessingUnit{
		Maker: "Intel",
		MHz:   1200,
		Price: 500,
	}

	// intial price
	fmt.Println("Original summary")
	c1.PrintSummary()

	// after discount
	c1.GiveDiscount(400)
	fmt.Println("After discount")
	c1.PrintSummary()

}
