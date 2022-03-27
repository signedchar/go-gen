package mult

func mult3(x int) bool {
	var i bool
	
	if x % 3 == 0 {
		i = true
	} else {
		i = false
	}

	return i
}

func mult5(x int) bool {
	var i bool

	if x % 5 == 0 {
		i = true
	} else {
		i = false
	}

	return i
}

func mult(div, x int) bool {
	var i bool

	if div % x == 0 {
		i = true
	} else {
		i = false
	}

	return i
}
