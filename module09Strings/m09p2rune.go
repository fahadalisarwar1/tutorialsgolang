package main

import "fmt"

func main() {
	fmt.Println("Rune Introduction")
	// it ddoes not matter how many bytes a code point occupies,
	// it can be represented by a rune

	printRune("Fahad Ali ")
	printRune("señor")
	loopRune("señor")

}

func printRune(name string) {
	r := []rune(name)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c %x", r[i], r[i])
		fmt.Println()
	}
}

func loopRune(str string) {
	runes := []rune(str)
	for _, value := range runes {
		fmt.Printf("%c %x", value, value)
		fmt.Println()
	}

}
