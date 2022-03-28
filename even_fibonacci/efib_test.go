package main

import "testing"

func TestFibonacci(t *testing.T) {
	if fib(10) == 89 {
		t.Fail()
	}
}
