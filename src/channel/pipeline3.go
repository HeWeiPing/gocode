package main

import (
	"fmt"
	"os"
	"strconv"
)

func counter(n int, out chan<- int) {
	for i := 0; i < n; i++ {
		out <- i
	}
	close(out)
}
func squarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	var n int
	naturals := make(chan int)
	squares := make(chan int)

	if len(os.Args[1:]) < 1 {
		n = 5
	} else {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("n=%d\n", n)
	}

	go counter(n, naturals)
	go squarer(naturals, squares)
	printer(squares)
}
