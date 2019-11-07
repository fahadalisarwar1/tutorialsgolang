package main

import "fmt"

// The idiomatic way to handle abnormal conditions in a program in Go is using errors. Errors are sufficient for most of the abnormal
// conditions arising in the program.
//
//But there are some situations where the program cannot simply continue executing after an abnormal situation. In this case we use
// panic to terminate the program. When a function
// encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to its caller.

// One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the
// program just cannot continue execution should a panic and recover mechanism be used

// panic looks like this
// func panic(interface{})
// The argument passed to panic will be printed when the program terminates.
func areaRectangle(length float64, width float64)(float64){
	if length == 0{
		panic("runtime error: Length is Zero")
	}
	if width == 0 {
		panic("runtime error: Width is Zero")
	}
	fmt.Println("Normal Execution")
	return length * width
}



func main(){
	var l, w float64
	//l = 33
	//w = 55
	area := areaRectangle(l, w)
	fmt.Println(area)
}