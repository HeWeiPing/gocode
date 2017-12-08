//使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。

package main

import "fmt"
import "sync"
import "time"
import "runtime"

var num uint32 = 1
var alp uint32 = 'A'
var b bool

//var ch chan byte
var ch chan bool
var mutex sync.Mutex

//func printNum(ch chan byte, m *sync.Mutex) {
//	for {
//			m.Lock()
//			num++
//			ch <- num
//			num++
//			ch <- num
//			m.Unlock()
//			runtime.Gosched()
//		}
//}
//
//func printAlphabet(ch chan byte, m *sync.Mutex) {
//	for {
//			m.Lock()
//			alp++
//			ch <- alp
//			alp++
//			ch <- alp
//			m.Unlock()
//			runtime.Gosched()
//	}
//}

func printNum(ch chan bool, m *sync.Mutex) {
	for num < 10 {
		b = <-ch
		if b == true {
			m.Lock()
			fmt.Printf("%d%d", num, num+1)
			num += 2
			m.Unlock()
		}
		ch <- false
		runtime.Gosched()
	}
}

func printAlphabet(ch chan bool, m *sync.Mutex) {
	for alp < 'A'+10 {
		b = <-ch
		//fmt.Println("Alp:", b)
		if b == false {
			m.Lock()
			fmt.Printf("%c%c", alp, alp+1)
			alp += 2
			m.Unlock()
		}
		ch <- true
		runtime.Gosched()
	}
}

func main() {
	//ch = make(chan byte, 100)
	ch = make(chan bool)
	//ch <- true
	//fmt.Printf("%c", alp)

	go printNum(ch, &mutex)
	go printAlphabet(ch, &mutex)
	ch <- true
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		runtime.Gosched()
		//fmt.Println(<-ch)
	}
		fmt.Println()
}
