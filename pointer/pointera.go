package main

import "fmt"

func main() {
	a, b := 20, 10

	p := &a
	fmt.Println(*p)
	*p = 30 //set a
	fmt.Println(a)

	p = &b
	*p = *p * 100
	fmt.Println(b)
}
