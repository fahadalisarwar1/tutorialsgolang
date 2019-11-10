package main

import "gopackages/patient"

import "fmt"

func main(){
	fmt.Println("Constructors")

	p1 := patient.New("Fahad", 23)
	p1.PrintSummary()

// although Go doesn't support classes, structs can
// effectively be used instead of classes and methods of signature New(parameters) can be used in the place of constructors
}