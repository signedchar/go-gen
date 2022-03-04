package main

import "fmt"

// Predicate
// Verify if a number is divisible between 2 and (number - 1)
// if return true is a prime number
// if return false is not a prime number
func prime_p(n int) bool {
	var lbool bool

	if n < 2 {
		return false
	} else if n == 2 {
		return true
	}

	for i := 2; i != n; i++ {
		if n%i == 0 {
			lbool = false
			break
		} else {
			lbool = true
		}
	}
	return lbool
}

// Return a segment of primes from "start" to "stop"
// it's not necessary calculate all the prime numbers.
func seg_primes(start, stop int) []int {
	var lprimes []int

	for i := start; i < stop; i++ {
		if prime_p(i) {
			lprimes = append(lprimes, i)
		}
	}
	return lprimes
}

// Calculate the summatory of a list of numbers.
func sum_reduce(lst []int) int {
	var acc int

	for i := 0; i < len(lst); i++ {
		acc += lst[i]
	}
	return acc
}

// Calculate the prime factors of a number
func primes_factors(num int) []int {
	lprimes := seg_primes(2, 10000)
	orign := num
	var fcts []int

	for i := 0; i < len(lprimes); i++ {
		if num%(lprimes[i]) == 0 {
			fcts = append(fcts, lprimes[i])
			num = num / lprimes[i]
			i -= 1
			continue
		}

		if num%(lprimes[i]) != 0 {
			if sum_reduce(fcts) == orign {
				break
			} else {
				continue
			}
		}
	}
	return fcts
}

func main() {
	/*
		fmt.Println(prime_p((600851475143)))
		fmt.Println(primes_factors(13195))
		fmt.Println(primes_factors(12))
		fmt.Println(primes_factors(18))
		fmt.Println(primes_factors(24))
		fmt.Println(primes_factors(20))
		fmt.Println(primes_factors(36))
	*/
	fmt.Println(primes_factors(600851475143))
}
