package main

import (
	"fmt"
	"time"
	"math/rand"
)

var abc = "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ"

// Return a random number from zero to len of some string.
func RandomIndex(s string) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(s))
}

// Return a random character from a string.
func RandomChar(s string) string {
	word := ""
	indx := RandomIndex(abc)
	char := string(abc[indx])
	word += char
	return word
}

// Return a random word using a string as base information.
func RandomWord(num int, s string) string {
	word := ""

	for i := 0; i < num; i++ {
		chr := RandomChar(abc)
		word += chr
	}

	return word
}

func main() {
	word := RandomWord(4, abc)
	word1 := RandomWord(10, abc)
	word2 := RandomWord(25, abc)
	fmt.Println("Four Chars:", word)
	fmt.Println("Ten Chars:", word1)
	fmt.Println("Twenty five Chars:", word2)
}
