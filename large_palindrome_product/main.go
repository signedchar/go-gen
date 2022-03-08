package main

import (
	"fmt"
	"strconv"
)

type palindrome struct {
	prd [2]int
	r   int
}

// Verify if "len(str)" is equal to "lng"
func pal_str_p(str string, lng int) bool {
	var lbool bool

	if len(str) == lng {
		lbool = true
	} else {
		lbool = false
	}

	return lbool
}

// Verify if the string of length == 6 is a palindrome.
func pal6(str string) bool {
	var lbool bool
	
	if pal_str_p(str, 6) {
		if str[0:3] == (str[5:6] + str[4:5] + str[3:4]) {
			lbool = true
		} else {
			lbool = false
		}
	}

	return lbool
}

// Verify if the string of length == 5 is a palindrome.
func pal5(str string) bool {
	var lbool bool

	if pal_str_p(str, 5) {
		if str[0:2] == (str[4:5] + str[3:4]) {
			lbool = true
		} else {
			lbool = false
		}
	}

	return lbool
}

func pals() []palindrome {
	var lpals []palindrome
	var ij [2]int
	var pal_tmp palindrome

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			str := strconv.Itoa(i * j)

			if pal6(str) {
				//string to integer == stoi
				stoi, _ := strconv.Atoi(str)
				ij = [2]int{i, j}
				pal_tmp.prd = ij
				pal_tmp.r = stoi
				lpals = append(lpals, pal_tmp)
				break
			}

			if pal5(str) {
				stoi, _ := strconv.Atoi(str)
				ij = [2]int{i, j}
				pal_tmp.prd = ij
				pal_tmp.r = stoi
				lpals = append(lpals, pal_tmp)
				break
			}
		}
	}

	return lpals
}

func maxl(lst []palindrome) palindrome {
	var pal_tmp palindrome
	max_tmp := lst[0].r

	for i := 1; i < len(lst); i++ {
		if max_tmp < lst[i].r {
			max_tmp = lst[i].r
			pal_tmp = lst[i]
		} else {
			continue
		}
		
	}

	return pal_tmp
}
   
func main() {
	palindromes := pals()
	pal := maxl(palindromes)
	x := pal.prd[0]
	y := pal.prd[1]
	fmt.Printf("The largest palindrome made from the product of 3 digits numbers is: %d = %d x %d\n", pal.r, x, y)
}
