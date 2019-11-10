package main

// Another example of interface

import "fmt"

// Interface
type VowelFinder interface {
	FindVowels() []rune
}

// struct
type Mystring string

func (str Mystring) FindVowels()[]rune{
	var vowels []rune
	for _, rune := range str{
		if rune == 'a'|| rune ==  'e'|| rune =='i'|| rune =='o'|| rune =='u'{
			vowels = append(vowels, rune)

		}
	}
	return vowels
}

func main(){
	fmt.Println("Example of Interfaces")
	name := Mystring("Fahad")
	var v VowelFinder
	v = name
	fmt.Printf("%c", v.FindVowels())
}
