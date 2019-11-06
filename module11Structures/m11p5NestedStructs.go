package main

import "fmt"

// It is necessary to note here that the structs doesn't necessarily contain only the buil-in types
// structs can contain fields which are composed of other structs.

// Address data
type Address struct {
	address string
	ZIPCODE int
}

// Match data
type Match struct {
	teamA string
	teamB string
	venue Address
}

func main() {
	var m Match
	fmt.Println("Example of Struct Within a struct.")
	m.teamA = "Gladiators"
	m.teamB = "Knights"
	m.venue = Address{
		address: "Grenoble",
		ZIPCODE: 38100,
	}

	// printing individual fields
	fmt.Println("Team A : ", m.teamA)
	fmt.Println("Team B: ", m.teamB)
	fmt.Println("Venue: ", m.venue.address)
	fmt.Println("ZIPCODE: ", m.venue.ZIPCODE)

}
