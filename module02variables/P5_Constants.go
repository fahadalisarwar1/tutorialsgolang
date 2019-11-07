package main

import "fmt"

const name = "fahad"

func main() {

	name = "sommethign else" // ERROR cant re-assign value to a constant

	fmt.Println("Constants tutorials")
	fmt.Println(name)

}
