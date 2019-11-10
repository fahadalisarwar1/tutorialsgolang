package main

import (
	"errors"
	"fmt"
)

func main(){
	// we can create our custom error by using the built in error type

	err := errors.New("Custom Error: An error has occured")
	fmt.Println(err)
}
