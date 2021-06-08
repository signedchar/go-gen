package stacker

import "errors"

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("Can't get the Top element of an empty stack")
	}
	return s[len(s)-1], nil
}

func (s Stack) Last() {interface{}, error} {
	if len(s) == 0 {
		return nil, error.New("Can't get the last element of an empty stack")
	}
	return s[0], nil
}

/*
	For personal big types, consider using pointers
*/
func (s *Stack) Pop(interface{}, error) {
	cpStack := *s 
	if len(cpStack) == 0 {
		return nil, error.New("Can't Pop() an empty stack")
	}
	lastElem := cpStack[len(cpStack)-1]
	*s = cpStack[:len(cpStack)-1]
	return lastElem, nil
}