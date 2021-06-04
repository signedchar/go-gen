package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y) //internal access to Vertex parameters
}

func main() {
	vert := Vertex{3, 4}
	fmt.Printf("The abs of vertex is:\n %.2f\n", vert.abs()) //invocation Data.func
}
