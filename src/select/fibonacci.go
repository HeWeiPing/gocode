package main

import "fmt"
import "time"

func fibonacci(c, quit chan int){
	x, y := 0, 1

	for{
		select{
		case c<-x:
			x, y = y, x+y
			time.Sleep(500*time.Millisecond)
		case <-quit:
			fmt.Println("\nquit")
			return
		}
	}
}

func main(){
	c := make(chan int)
	quit := make(chan int)

	go func(){
		for i:=0; i<10; i++{
			fmt.Printf("%-3d ", <-c)
		}
		quit <- 1
	}()

	fibonacci(c, quit)
}
