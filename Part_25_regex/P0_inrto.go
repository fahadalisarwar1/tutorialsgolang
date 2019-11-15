package main

import (
	"fmt"
	"regexp"
)

func main(){
	// regex is used to describe a search pattern for matching the text

	// regex is a sequence of some characters that define a pattern

	// regex is used for parsing filtering and validating meaningful information from large text

	// simple example with extracting text between a []
	str := "My name is Fahad and Fasfggfd"

	r, _ := regexp.Compile("F([a-z]+)d")

	matches := r.FindAllString(str, -1)
	for i, match := range matches{
		fmt.Println(i, " ", match)
	}
	fmt.Println()

}
