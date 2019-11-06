package main

import (
	"fmt"
	"gopackages/computer"
)

func main() {
	fmt.Println("Exported Structs")
	var myCPU computer.ProcessingUnit
	myCPU.Maker = "Intel"
	myCPU.MHz = 1200
	myCPU.Price = 500
	fmt.Println(myCPU)

}
