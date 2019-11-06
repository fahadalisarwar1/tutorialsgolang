package main

import "fmt"

func main() {
	fmt.Println("Constructing String From Bytes")

	bytes := []byte{0x48, 0x49}
	// bytes are represented by their hex values

	stringSlice := string(bytes)
	fmt.Println(stringSlice)
	// types conversion help us to convert bytes to strings in this case
	// type conversion also works for converting other types
	// for example int32 to int64 -> int64(int32)
	decBytes := []byte{67, 97, 102, 195, 169}

	// Note that the Word Cafe contains a unicode so it is represented by 5 bytes not four.

	str := string(decBytes)
	fmt.Println(str)
	fmt.Println("lenght of string  Senor is", len(str))

	// any valid unicode character in single quotes in a rune

	// strings are immutable
	// once a string is created, its not possible to change its value
	var myName string
	myName = "Fahad"
	// myName[0] = "A" // This will throw an error
	fmt.Println(myName)

}
