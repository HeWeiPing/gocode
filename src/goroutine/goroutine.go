package main

import "fmt"

func add(a, b int) {
	z := a + b
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go add(1, 2)
	}
}
