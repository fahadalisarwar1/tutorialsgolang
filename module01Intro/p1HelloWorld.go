package main

// Topics to be covered
// -----------------------

// -----------------------

// Package main, This should be the fist line of any go code, this refers to creating a package
// main is the package which makes the files executable,
// if the name of the package is something else, then the package is not executable and can't be run !
// There are two types of packages in GoLang.
// 1: Reusable packages without executable
// 2: Executable packages.

import "fmt"

// this shows how to import packages into the file, the scope of the import is limited to this file only
// fmt is the part of standard package library of GoLang and is a short hand for "format"
// This package helps in formatting of go code and also the standard output of the GoLang

// keyword "func" defines the function similar to other programming languages
// main is the entry point of the executable for this project

func main() {
	// This block contains code run by the main program....

	fmt.Println("Hello World, Welcome to GoLang Programming")
	// Println function prints the text to the standard output, in this case the terminal,.....

}
