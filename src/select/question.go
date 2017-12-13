package main

import "fmt"
import "time"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
			time.Sleep(1 * time.Second)
		case <-quit:
			fmt.Println("\nquit")
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func main() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			<-c //在这执行读取10次都没事
		}
	}()

	time.Sleep(1 * time.Second)
	//<-c //放在这就会就执行不了

}
