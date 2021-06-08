package main

import (
	"fmt"
	"stacker/stack"
)

func main() {
	var mystack stacker.Stack

	mystack.Push("abc")
	mystack.Push(123)
	mystack.Push([]string{"z", "y", "x"})
	mystack.Push(3.14)
	mystack.Push(-3)

	for {
		val, err = mystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(val)
	}
}
