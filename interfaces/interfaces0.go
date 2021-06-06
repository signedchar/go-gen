package main

import (
	"fmt"
	"math"
)

// using two new data types
type Circle struct {
	Radius float64
}

type Square struct {
	Width  float64
	Height float64
}

// methods for calculate the area

// Circle area
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Square area
func (s Square) Area() float64 {
	return s.Width * s.Height
}

// unify area abstraction in a new type
type Sizer interface {
	Area() float64
}

// use the interface for build a more powerfull function
func LessArea(polg0, polg1 Sizer) Sizer {
	if polg0.Area() < polg1.Area() {
		return polg0
	}
	return polg1
}

func main() {
	cir := Circle{Radius: 10}
	sqare := Square{Width: 20, Height: 20}

	areas := LessArea(cir, sqare)
	fmt.Printf("%+v is the smaller area.\n", areas)
}
