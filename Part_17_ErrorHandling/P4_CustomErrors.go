package main

import (
	"errors"
	"fmt"
)

// The simplest way to create a custom error is to use New function offered by error package.
// show documentation of errors here

// https://golang.org/src/errors/errors.go?s=293:320#L1

func divide(num1 float64, num2 float64)(float64, error){
	if num2 == 0{
		return 0, errors.New("Denominator can't be zero, Currently it is Zero")
	}
	return num1/num2, nil
}

func main(){
	var marks float64
	var totalMarks float64
	marks = 50
	totalMarks = 0
	res, err := divide(marks, totalMarks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result:  %0.2f\n", res)
}