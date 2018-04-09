package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	p1 := f()
	p2 := f()
	f()
	fmt.Printf("p1=%v\np2=%v\n", p1, p2)
	fmt.Println(p1 == p2)
}
