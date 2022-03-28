package main

import 	"testing"

func TestFibonacci(t *testing.T) {
	if fib(11) != 89 {
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

func TestSumEvenFibonacciNumbers(t *testing.T) {
	if sumEvenFib(9) != 44 {
		t.Log(sumEvenFib(9))
		t.Error("Expected returned the sum of even fibonacci numbers from 0 to 9.")
	}
}
