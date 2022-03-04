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

// Return the primes from 2 to `num`
func primes(num int) []int {
	var lprimes []int

	for i := 2; i < num; i++ {
		if prime_p(i) {
			lprimes = append(lprimes, i)
		}
	}
	return lprimes
}

func main() {
	fmt.Println(prime_p((600851475143)))
	fmt.Println(primes(100))
	/*
		for i := 0; i < 550; i++ {
			fmt.Printf("%d: %t\n", i, prime_p(i))
		}
	*/

}
