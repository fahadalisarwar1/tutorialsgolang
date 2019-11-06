package main

import "fmt"

// Strings deserve a special mention in Go as they are different
// in implementation when compared to other languages.
// A string in Go is a slice of bytes. Strings can be created by enclosing their contents inside " ".

// strings in go are Unicode compliant and UTF-8 encoded
func printBytes(s string) {
	// strings are slices of bytes
	// This function prints the bytes HEX values of the given string
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // %x is the format specifier for hexadecimal values
		fmt.Printf("%x ", s[i])
		fmt.Println()
	}
}

func main() {
	fmt.Println("Strings introduction")
	str := "Fahad Ali Sarwar"
	printBytes(str) // This works fine for ASCII characters but dosn't Work for Unicode characters

	// fmt.Println()
	// fmt.Println()
	printBytes("seÑor")
	// This causes a problem because unicode encoding of N occupies two
	// bytes, we are assuming that each character will be only one byte which
	// is true of ASCII code but not for UNICODE
	// for this we need "rune" data type which will be explained next

	fmt.Println("Len of string is ", len(str))

	// fmt.Println()
}
