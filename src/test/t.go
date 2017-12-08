package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 999
	fmt.Println(len(ch))
	fmt.Println(<-ch)

}
