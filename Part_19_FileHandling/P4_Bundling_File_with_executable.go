package main

//  The most convinient option is to bundle the file with the executbale

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("../filehandling")
	//data = box.String("sample.txt")
	data, _ := box.FindString("sample.txt")
	fmt.Println("Contents of file:", data)
}