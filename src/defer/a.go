package main

import "fmt"

func foo(a int, b int, c int){
	fmt.Println(a, b, c)
}

func main() {
	a := 1
	b := 2
	c := 3
	defer foo(a, b, c)
	a = 4
	b = 5
	c = 6
	defer foo(a, b, c)
}
