package mult

func mult(num, div int) bool {
	var i bool

	if num % div == 0 {
		i = true
	} else {
		i = false
	}

	return i
}

func smult(limit int) int {
	var count int
	
	for i := 1; i < limit; i++ {
		if mult(i,3) || mult(i,5) {
			count += i
		}
	}

	return count
}
