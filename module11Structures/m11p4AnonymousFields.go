package main

import "fmt"

// Players data
type Players struct {
	string
	int // instead of defininf a variable name we just define
	// the type
}

func main() {
	fmt.Println("Anonymous Fields")
	plyr1 := Players{
		"Fahad",
		50,
	}
	fmt.Println(plyr1)

	// one thing to note here is that
	// even thoough we didnt define varibale names
	// the variable names are the type of the fieds
	fmt.Println(plyr1.string)
	fmt.Println(plyr1.int)

}
