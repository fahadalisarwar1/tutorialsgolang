package main

// when we run a go program we have already seen that we pass them certain arguments

// go run -v file.go

// we can also pass flags to the go commmand

//
import (
	"fmt"
	"os"
)
func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

}