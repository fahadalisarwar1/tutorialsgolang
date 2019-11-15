package main

import "fmt"

func main() {
	greet := func(name string) (welcomeMsg string) {
		return "Welcome, " + name
	}
	fmt.Println(greet("Fahad"))
	fmt.Println(greet("Sarwar"))
}
