package main

import (
	"crypto/sha256"
	"fmt"
)

var count = 0

func main() {
	s1 := sha256.Sum256([]byte("a"))
	s2 := sha256.Sum256([]byte("A"))
	same := make([]byte, len(s1))

	fmt.Printf("s1:")
	for _, v := range s1 {
		fmt.Printf("%3x", v)
	}
	fmt.Printf("\ns2:")
	for _, v := range s2 {
		fmt.Printf("%3x", v)
	}
	fmt.Println()

	for i, _ := range s1 {
		if s1[i] == s2[i] {
			same[i] = '^'
			continue
		}
		same[i] = '_'
		count++
	}

	fmt.Printf("\n%3c", ' ')

	for _, v := range same {
		fmt.Printf("%3c", v)
	}
	fmt.Printf("\n")

	fmt.Printf("Different bits: %d\n", count)
}
