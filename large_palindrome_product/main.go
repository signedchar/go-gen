package main

import (
	"fmt"
	"strconv"
)

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

// Return a list of palindromes. The numbers have the length 6 and 5.
func pals() []int {
	var lpals []int
	
	for i:= 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			str := strconv.Itoa(i * j)

			if pal6(str) {
				stoi, _ := strconv.Atoi(str)
				lpals = append(lpals, stoi)
				break
			}

			if pal5(str) {
				stoi, _ := strconv.Atoi(str)
				lpals = append(lpals, stoi)
				break
			}
		}
	}

	return lpals
}

func main() {
	fmt.Println(pals())
}
