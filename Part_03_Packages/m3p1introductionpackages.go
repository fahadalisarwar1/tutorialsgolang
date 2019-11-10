package main

// so far we have been using only one file to write go programs, this approach is good for learning but not very helpful
// in code reusablility and managing large amount of code.

// packages offer compartmentalization of code hence better readability and maintenanace.

// to put custom packages in /go/src

import (
	"fmt"
	"geometery"
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println(geometery.Area(rectLen, rectWidth))
}
