package main

// interfaces allows us to do polymorphism in go
// different function is called based on type of object passed.
import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64

}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area()(float64){
	return r.width * r.height
}

func (r rectangle) perimeter()(float64){
	return 2 *(r.width + r.height)
}
func (c circle) area()(float64){
	return math.Pi*c.radius*c.radius
}

func (c circle) perimeter()(float64){
	return 2*math.Pi*c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {
	fmt.Println("Introduction to Interfaces.")


	c := circle{radius:5}
	r := rectangle{
		width:3,
		height:4,
	}

	measure(c)
	measure(r)

	// interface defines the behaviour of an object
	// Interfaces are set of methods signatures
	// interface determines what methods a type should have and
	//  type determines how to implement these methods.

}
