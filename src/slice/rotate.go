package main

import (
	"errors"
	"fmt"
)

func rotate(s []int) ([]int, error) {
	if len(s) <= 0 {
		err := errors.New("Slice empty!\n")
		return s, err
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.
	return s, nil
}

func main() {
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)
	slice, err := rotate(slice)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(slice)
}
