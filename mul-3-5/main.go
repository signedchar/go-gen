package main

import (
	"fmt"
)

func mulThree(x int) bool {
	if x%3 == 0 {
		return true
	} else {
		return false
	}
}

func mulFive(x int) bool {
	if x%5 == 0 {
		return true
	} else {
		return false
	}
}

// sum of multiples 3 and 5 of x
func m35(x int) int {
	var count = 0

	for i := 0; i < x; i++ {
		if mulThree(i) || mulFive(i) {
			count += i
		}
	}

	return count
}

func main() {
	var sum35_1000 = m35(1000)
	var sum35_10 = m35(10)

	fmt.Printf("Sum of first 1000 multiples 3 and 5: %d\n", sum35_1000)
	fmt.Printf("Sum of first 10 multiples 3 and 5:   %d\n", sum35_10)
}
