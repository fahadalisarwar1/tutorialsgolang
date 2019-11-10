package main

import (
	"fmt"
	"io/ioutil"
)

// abs path:  /home/fahad/go/src/github.com/fahadalisarwar1/tutorialsgolang
// Relative Path: github.com/fahadalisarwar1/tutorialsgolang

func main(){
	data, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}