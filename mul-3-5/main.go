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

func main() {
	var count = 0

	for i := 0; i < 1000; i++ {
		if mulThree(i) || mulFive(i) {
			count += i
		}
	}
	
	fmt.Printf("Sum of first 1000 multiples 3 and 5: %d\n", count)
}
