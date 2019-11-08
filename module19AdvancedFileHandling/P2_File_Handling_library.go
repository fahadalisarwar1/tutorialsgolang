package main

import (
	"fmt"
	"os"
)

// to open a file in go we use the Open function

func main(){
	file, err := os.Open("sample.txt")
	if err != nil{
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil{
		fmt.Println("Error getting file stats")
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil{
		fmt.Println("ERROR reading file")
		return
	}
	str := string(bs)
	fmt.Println(str)


}