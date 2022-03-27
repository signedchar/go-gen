package mult

func mult(div, x int) bool {
	var i bool

	if div % x == 0 {
		i = true
	} else {
		i = false
	}

	return i
}
