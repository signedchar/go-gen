package main

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
