package main

import (
	"fmt"
	"os"
)

// Error is an interface which a method called Error()
// Open documentation here
// in the example above we just printed the error, we can also see more details about the error

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error()string{
	return e.Op + " " + e.Path + " "+ e.Err.Error()
}

func main(){
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}