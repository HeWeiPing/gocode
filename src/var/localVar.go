package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		x := i + 1
		y := i + 3
		fmt.Printf("&i=%v i=%d\n", &i, i)
		fmt.Printf("&x=%v x=%d\n", &x, x)
		fmt.Printf("&y=%v y=%d\n", &y, y)
	}
}
