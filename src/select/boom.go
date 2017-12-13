package main

import "fmt"
import "time"

func main() {
	tick := time.Tick(500 * time.Millisecond)
	boom := time.After(1500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("di.")
		case <-boom:
			fmt.Println("Boom!!!")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
