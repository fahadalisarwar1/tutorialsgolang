package main

import "fmt"

// Go doesn't support inheritance but it does support composition

// A Car is COMPOSED up of tyres etc. so composition inherently provides the same capability as inheritance

// In layman terms, a composition is a struct composed up of other structs

// the main struct has access to all the fields and methods of the sub-struct

// Phone
type Phone struct {
	Make string
	Model string
	Camera
	Screen
}

type Camera struct {
	Megapixels int  // 16
	Zoom bool
}

type Screen struct {
	ScreenSize int // 5 inches
}

func (p Phone)PrintSummary(){
	fmt.Println("Make:\t",p.Make)
	fmt.Println("Model:\t",p.Model)
	fmt.Println("MP:\t",p.Camera.Megapixels)
	fmt.Println("Zoom:\t",p.Camera.Zoom)
	fmt.Println("Screen:\t",p.Screen.ScreenSize)
}

func main()  {
	cam := Camera{Megapixels:16, Zoom:true}
	scr := Screen{ScreenSize:5}
	myPhone := Phone{Make:"Oneplus", Model:"3T", Camera: cam, Screen:scr}
	myPhone.PrintSummary()

}