package main

import "fmt"

var ch chan bool

func Go(ch chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	ch <- true
}

func main() {
	ch = make(chan bool)
	var pch *chan bool
	pch = &ch
	fmt.Printf("ch type:%T\n", ch)
	fmt.Printf("pch type:%T\n", pch)
	fmt.Printf("cap(ch)=%d\n", cap(ch))

	go Go(*pch)
	//time.Sleep(2 * time.Second)
	<-ch //Be blocked until GO() finish
	//<-ch //dead lock
}
