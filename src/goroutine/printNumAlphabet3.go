//使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。

// 别人的写法
package main

import (
	"fmt"
	"sync"
)

func PrintNums(printChar chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", 2*i+j+1)
		}

		printChar <- 1
		<-printChar
	}
}

func PrintChars(printChar chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 5; i++ {

		<-printChar

		for j := 0; j < 2; j++ {
			fmt.Printf("%c", 'A'+(2*i+j))
		}

		printChar <- 1
	}
}

func main() {

	flag := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go PrintNums(flag, &wg)
	go PrintChars(flag, &wg)

	wg.Wait()
}
