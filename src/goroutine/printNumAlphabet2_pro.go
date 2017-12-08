//使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。

package main

import "fmt"
import "sync"

func printNum(ch chan bool, wg *sync.WaitGroup) {
	var num int = 1
	for num < 10 {
		fmt.Printf("%d%d", num, num+1)
		num += 2
		ch <- true
		<-ch
	}
	defer wg.Done()
}

func printAlphabet(ch chan bool, wg *sync.WaitGroup) {
	var alp int = 'A'
	for alp < 'A'+10 {
		<-ch
		fmt.Printf("%c%c", alp, alp+1)
		alp += 2
		ch <- false
	}
	defer wg.Done()
}

func main() {
	var ch chan bool
	ch = make(chan bool)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go printNum(ch, &wg)
	go printAlphabet(ch, &wg)

	wg.Wait()
	fmt.Println()
}
