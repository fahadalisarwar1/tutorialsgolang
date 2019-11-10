package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "newFile.txt"
	file, err :=  os.Create(fileName)
	if err != nil{
		fmt.Println("couldn't create a new file named:", fileName)
	}
	defer file.Close()
	file.WriteString("This is a newly created file")
}
