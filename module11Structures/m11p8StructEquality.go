package main

// tructs are value types and are comparable if each of
//their fields are comparable. Two
// struct variables are considered equal if their
// corresponding fields are equal.
import (
	"computer"
	"fmt"
)

func main() {
	fmt.Println("Structs Equality")
	s1 := computer.ProcessingUnit{
		Maker: "intel",
		MHz:   1200,
		Price: 500,
	}

	s2 := computer.ProcessingUnit{
		Maker: "amd",
		MHz:   1200,
		Price: 500,
	}

	if s1 == s2 {
		fmt.Println("Structs are equal in all fiels")
	} else {
		fmt.Println("Structs are not equal")
	}
}
