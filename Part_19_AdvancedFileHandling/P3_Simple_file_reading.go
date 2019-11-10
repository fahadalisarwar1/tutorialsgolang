package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	data, err := ioutil.ReadFile("sample.txt")
	if err != nil{
		return
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
