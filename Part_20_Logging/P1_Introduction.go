package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file, err := os.OpenFile("error.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Application started")
	num1 := 44
	num2 := 3
	if num2 == 0{
		log.Print("Denominator is Zero") // whenever log.print is called, log file will be apended

	} else {
		res := num1 / num2
		fmt.Println(res)
	}
}
