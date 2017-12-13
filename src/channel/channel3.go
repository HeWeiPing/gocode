package main

import "fmt"
import "time"

func Producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("P:", i)
	}
}

func Consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("C:", <-ch)
	}
}

func main() {
	ch := make(chan int)

	go Producer(ch)
	go Consumer(ch)

	//互换下面两行会出错，为什么?`
	fmt.Println("main:", <-ch)
	time.Sleep(2 * time.Second)
}
