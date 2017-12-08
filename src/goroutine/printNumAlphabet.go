//使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。

package main

import "fmt"
//import "sync"
//import "time"


var ch chan string
var n int = 0
var c byte = 'A'

func printNum(ch* chan string, num *int) {
	for {
		v, ok := <-*ch
		if !ok {
			fmt.Println("channel CLOSED")
			//time.Sleep(1*time.Second)
		}
		if v == "NT"{
			fmt.Printf("%d", *num)
			fmt.Printf("%d", *num)
			*ch <- "AT" 
			break
		}
	}
}

func printAlphabet(ch* chan string, alp* byte) {
	for {
		v, ok := <-*ch
		if !ok {
			fmt.Println("channel CLOSED")
			//time.Sleep(1*time.Second)
		}
		if v == "AT"{
			fmt.Printf("%c", *alp)
			fmt.Printf("%c", *alp)
			*ch <- "NT" 
			break
		}
	}
}

func main() {
	counter := 0
	ch = make(chan string, 2)
	ch <- "NT"

	for {
		counter ++
		if counter > 20{
			close(ch)
			break
		}
		printNum(&ch, &n)
		printAlphabet(&ch, &c)
	}
}
