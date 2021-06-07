package main

import "fmt"

type Square struct {
	Width  float64
	Height float64
}

// normal function
func addseven(num float64) float64 {
	return num + 7
}

// pointer function
func addeight(s *Square, num float64) {
	s.Height = s.Height*num + 1
	s.Width = s.Width*2 + num
}

// pointer method
func (s *Square) heightNew(num float64) {
	s.Height = s.Height * num
	s.Width = s.Width * num
}

func main() {
	square := Square{4, 4}

	fmt.Printf("Square addseven, %.2f, square: %v\n", addseven(square.Height), square) //using normal fn
	square.heightNew(2.0)                                                              //pointer method
	fmt.Printf("The new square: %v\n", square)                                         //printing the change

	anotherSquare := Square{7, 7}
	fmt.Printf("Printing anotherSquare: %v\n", anotherSquare)
	addeight(&anotherSquare, 5) //applying pointer fn

	fmt.Printf("The new anotherSquare: %v\n", anotherSquare)
}
