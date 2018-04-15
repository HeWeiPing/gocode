package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//ncpu := runtime.NumCPU()
	//fmt.Println("CPU numbers:", ncpu)

	ncpu := flag.Int("n", 1, "CPU numbers")
	flag.Parse()

	fmt.Println("CPU numbers:", *ncpu)
	for i := 0; i < *ncpu; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				for k := 0; k < 8; k++ {
					fmt.Printf("%d", j%2)
				}
			}
		}()
	}
	time.Sleep(1 * time.Second)
}
