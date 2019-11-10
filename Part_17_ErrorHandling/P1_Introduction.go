package main

import (
	"fmt"
	"os"
)

// Errors: Errors indicate an abnormal condition in the program. Let's say we are trying to open a file and the
// file does not exist in the file system. This is an abnormal condition and it's represented as an error.

// Simple Example


// Lets open a file which is not present in the OS

func main(){
	file, err := os.Open("/abc.txt")
	if err != nil{
		fmt.Println(err) // an error is caused since there is no such file or directory
		return
	}

	// The idiomatic way of handling error in Go is to compare the returned error to nil. A nil value
	// indicates that no error has occurred and a non nil value indicates the presence of an error
	fmt.Println(file.Name(), "Opened successfully")
}