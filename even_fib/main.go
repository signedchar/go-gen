package main

import "fmt"

// Return the fibonacci of a number
func fib(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
}

// Predicate
// Return true if the value is major 4_000_000
// else return false
func major_4m_p(num int) bool {
	if num >= 4000000 {
		return true
	} else {
		return false
	}
}

// Calculate the max value of fibonacci that not
// exceed 4_000_000
func fib_minor_4m() int {
	var prev = 0
	var limit_fib = 0

	for i := 0; i < 50; i++ {
		prev = i - 1
		limit_fib = fib(i)

		if major_4m_p(limit_fib) {
			break
		}
	}
	return prev
}

// Predicate
// Return true if the number is even, else return false
func evenp(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

// sum the even fibonacci values minors of 4_000_000
func sum_even_fib(num int) int {
	var acc = 0
	var internal_fib = 0

	for i := 0; i != (num + 1); i++ {
		internal_fib = fib(i)
		if evenp(internal_fib) {
			acc += internal_fib
		}
	}

	return acc
}

func main() {
	fmt.Println(sum_even_fib(fib_minor_4m()))
}
