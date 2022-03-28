package main

import "fmt"

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func evenp(x int) bool {
	var i bool
	
	if x % 2 == 0 {
		i = true
	} else {
		i = false
	}

	return i
}

func sumEvenFib(x int) int {
	// ifib == internal fibonacci
	var acc int
	var ifib int
	
	for i := 0; i <= x; i++ {
		ifib = fib(i)
		if ifib >= 4000000 {
			break
		}

		if evenp(ifib) {
			acc += ifib
		}
	}

	return acc
}


func main() {
	fmt.Println(sumEvenFib(50000))
}
