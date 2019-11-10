package main

import "fmt"

// lets say we have two people who are male and female
type human interface {
	sayHello() string
}

type man struct {
	greeting string
}

type woman struct {
	greeting string
}

type child struct {
	greeting string
}

func (m man) sayHello()(string){
	return m.greeting
}

func (w woman) sayHello()(string){
	return w.greeting
}

func (c child) sayHello()(string){
	return c.greeting
}

func GreetUser(h human){
	fmt.Println(h.sayHello())
}

func main(){
	fahad := man{greeting:"HEllo man whats up"}
	GreetUser(fahad)

	sarah := woman{greeting:"Hi"}
	GreetUser(sarah)

	kid := child{greeting:"I'm hungry"}
	GreetUser(kid)
	// different functions will be called based on what type of object is passed
	// also all methods are called automatically by this funciton

}