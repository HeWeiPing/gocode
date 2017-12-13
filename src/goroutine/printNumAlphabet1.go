//使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。

package main

import "fmt"
import "sync"

var num uint32 = 1
var alp uint32 = 'A'
var b bool = true

func printNum(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for num < 10 {
		if b == true {
			fmt.Printf("%d%d", num, num+1)
			num += 2
		}
		ch <- false
		b = <-ch
	}
}

func printAlphabet(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for alp < 'A'+10 {
		b = <-ch
		if b == false {
			fmt.Printf("%c%c", alp, alp+1)
			alp += 2
		}
		ch <- true
	}
}

func main() {
	ch := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)
	go printAlphabet(ch, &wg)
	go printNum(ch, &wg)

	wg.Wait()
	fmt.Println()

}
