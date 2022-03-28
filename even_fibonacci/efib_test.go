package main

import "testing"

func TestFibonacci(t *testing.T) {
	if fib(10) == 89 {
		t.Fail()
	}
}

func Testevenp(t *testing.T) {
	if evenp(2) != true {
		t.Fail()
	}

	if evenp(3) == true {
		t.Fail()
	}
}
