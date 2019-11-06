package main

import "fmt"

// Address data
type Address struct {
	city    string
	zipcode int64
}

// Person data
type Person struct {
	name string
	Address
}

func main() {
	fmt.Println("Promoted Fields")
	var p Person
	p.name = "fahad"
	p.city = "Grenoble" // Promoted fields as they can be accessed directly
	p.zipcode = 38100

	fmt.Println(p.name)
	fmt.Println(p.city)
	fmt.Println(p.zipcode)
}
