package main

import "fmt"

func main() {

	fmt.Println("Variable Types")
	// bool true or false
	isGraduated := true
	isRich := false

	fmt.Println(isGraduated && isRich) // Logical Operators

	fmt.Println(isGraduated || isRich)

	// integers
	// 8 16 32 64 bits
	// int8 -127 to 127
	// uint 0 to 255
	num1 := 22 // by default int{could be 32 or 64 bit depending on the system architecture}

	marks := 33.455
	fmt.Println(num1, marks)

	// strings

	name := "fahad ali sarwar"
	fmt.Println("My name is ", name)

	// complex data with imaginary parts
	// complex64 --> Real 32 and imaginary 32 bits

	// similarly complex128 real 64 and imaginary 64 bits
	// two ways to define them

	a := 1 + 3i
	b := complex(1, 4)
	fmt.Println(a + b)

}
