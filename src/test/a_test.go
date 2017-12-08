package main

import "fmt"
import "testing"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

func Test_test(t *testing.T) {
	fmt.Println("testing...")

	ch := make(chan int)
	fmt.Println(len(ch))
}
