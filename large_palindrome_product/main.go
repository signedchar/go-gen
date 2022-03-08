package main

import (
	"fmt"
	"strconv"
)

type palindrom struct {
	m   int
	n   int
	pal string
}

type lpalindrom struct {
	lstpal []palindrom
}

func reverse(s string) string {
	r := ""

	for _, v := range s {
		r = string(v) + r
	}

	return r
}

func (p palindrom) even_word_p() bool {
	if len(p.pal)%2 == 0 {
		return true
	} else {
		return false
	}
}

func (p palindrom) h_tail_pal() [2]string {
	var h_pal, tail_pal string
	var r [2]string
	lng := len(p.pal)

	if p.even_word_p() {
		hlf := lng / 2
		h_pal = p.pal[0:hlf]
		tail_pal = reverse(p.pal[hlf:lng])
		r = [2]string{h_pal, tail_pal}
	} else {
		hlf := (lng - 1) / 2
		h_pal = p.pal[0:hlf]
		tail_pal = reverse(p.pal[hlf+1 : lng])
		r = [2]string{h_pal, tail_pal}
	}

	return r
}

func (p palindrom) palindrom_p() bool {
	pal := p.h_tail_pal()
	if pal[0] == pal[1] {
		return true
	} else {
		return false
	}
}

func (lp lpalindrom) palindromes(start, stop int) lpalindrom {
	// spal == single palindrome
	var spal palindrom

	for i := start; i >= stop; i-- {
		for j := start; j >= stop; j-- {
			r := i * j
			spal.pal = strconv.Itoa(r)
			if spal.palindrom_p() {
				spal.m = i
				spal.n = j
				lp.lstpal = append(lp.lstpal, spal)
				break
			}
		}
	}

	return lp
}

func (lp lpalindrom) max() palindrom {
	var pal_tmp palindrom
	var r int

	r, _ = strconv.Atoi(pal_tmp.pal)

	for i := 1; i < len(lp.lstpal); i++ {
		lr, _ := strconv.Atoi(lp.lstpal[i].pal)
		if r < lr {
			r = lr
			pal_tmp = lp.lstpal[i]
		} else {
			continue
		}
	}

	return pal_tmp
}

func main() {
	var lp lpalindrom
	palindroms := lp.palindromes(999, 100)
	pal_max := palindroms.max()
	fmt.Printf("The largest palindrome made from the product of 3 digits numbers is: %d x %d = %s", pal_max.m, pal_max.n, pal_max.pal)
}
