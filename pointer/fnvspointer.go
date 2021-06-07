package main

import "fmt"

/*
	Demostration pointers vs no pointers in functions and methods.
	The normal functions or methods not change the input value
	The pointer functions or pointer methods, change the input value
*/

type Square struct {
	Width  float64
	Height float64
}

// normal function
func addseven(num float64) float64 {
	return num + 7
}

// pointer function
func ptrfn(s *Square, num float64) {
	s.Height = s.Height*num + 1
	s.Width = s.Width*2 + num
}

// normal method
func (s Square) mtUpdateHW(num float64) Square {
	s.Height = s.Height * num
	s.Width = s.Width * num
	return s
}

// pointer method
func (s *Square) mtpUpdateHW(num float64) {
	s.Height = s.Height * num
	s.Width = s.Width * num
}

func main() {
	square := Square{4, 4}

	fmt.Println()
	fmt.Println("Pointer demostration for normal functions/methods and pointer functions/methods")
	fmt.Println()

	fmt.Printf("Printing square: %v\n", square)
	fmt.Printf("Using normal function for square, %.2f, square: %v\n", addseven(square.Height), square) // using normal fn
	ptrfn(&square, 5)
	fmt.Printf("Printing square after apply pointer function: %v\n", square)

	fmt.Println("")

	anotherSquare := Square{7, 7}
	fmt.Printf("Printing anotherSquare: %v\n", anotherSquare)
	fmt.Printf("Using normal method for: %v, anotherSquare: %v\n", anotherSquare.mtUpdateHW(2), anotherSquare) // normal method

	anotherSquare.mtpUpdateHW(2)                                                  // pointer method
	fmt.Printf("Printing square after apply pointer method: %v\n", anotherSquare) //printing the change
}
