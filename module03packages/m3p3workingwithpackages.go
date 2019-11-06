package main

import (
	"fmt"
	"gopackages/geometery"
)

func main() {
	fmt.Println("main package executed")
	a, b := 44.4, 3.4
	area := geometery.Area(a, b)
	fmt.Println(area)

}
