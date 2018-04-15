package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	naturals := make(chan int)
	squares := make(chan int)
	//fmt.Printf("global n:%p\n", &n)

	if len(os.Args[1:]) < 1 {
		n = 5
	} else {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		//fmt.Printf("local n:%p\n", &n)
		if err != nil {
			fmt.Printf("strconv: %s error!\n", os.Args[1])
		}
		fmt.Printf("n=%d\n", n)
	}

	//Counter
	go func() {
		for x := 0; x < n; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				break //channel was chosed and drained
			}
			squares <- x * x
		}
	}()

	//Printer in main goroutine
	for {
		v, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
