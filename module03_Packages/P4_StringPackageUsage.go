package main

import (
	"fmt"
	"strings"
)

func main(){
	// check if a string contains a sub string
	fmt.Println(strings.Contains("my name is fahad", "fahad"))

	// Count the number of times a substring occurs in a bigger string
	fmt.Println(strings.Count("yes yes yes no no", "yes"))
	fmt.Println(strings.Count("yes yes yes no no", "no"))

	// check if a string has a certain prefix
	fmt.Println(strings.HasPrefix("hello world", "hell"))
	fmt.Println(strings.HasPrefix("hello world", "ell"))

	// check if a string has certain suffix
	fmt.Println(strings.HasSuffix("fahad ali sarwar", "war"))
	fmt.Println(strings.HasSuffix("fahad ali sarwar", "ss"))

	// strings is an array of characters
	// to find the index of a substring
	fmt.Println(strings.Index("fahad ali sarwar", "ali"))
	fmt.Println(strings.Index("fahad ali sarwar", "usman"))  // returns -1 if no match is found

	// joining a list of strings to string
	fmt.Println(strings.Join([]string{"fahad", "ali", "sarwar"}, " "))

	// repeat a string
	fmt.Println(strings.Repeat("fahad", 5))

	// to replace a sub string in a larger string
	fmt.Println(strings.Replace("fahad ali sarwar", "a", "A", -1))
	fmt.Println(strings.Replace("fahad ali sarwar", "a", "A", 2))


	// To split a string into a list of strings by a separating string (e.g., a comma), use the
	// Split function
	newString := strings.Split("fahad ali sarwar", " ")
	fmt.Println(newString)
	fmt.Println(len(newString))


	// convert to uppercase
	fmt.Println(strings.ToUpper("fahad ali sarwar"))

	fmt.Println(strings.ToLower("FAHAD ALI SAARWAR"))



}