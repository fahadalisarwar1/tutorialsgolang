package main

import (
	"fmt"
	"os"
)

func main(){
	dir, err := os.Open(".")
	if err != nil{
		fmt.Println("Unable to read directory")
		return
	}
	defer dir.Close()
	fileInfo, err := dir.Readdir(-1) // n limits the number of entries returned, -1 represents all entries
	if err != nil{
		fmt.Println("ERror")
		return
	}
	for _, file := range fileInfo{
		fmt.Println(file.Name())
	}
}