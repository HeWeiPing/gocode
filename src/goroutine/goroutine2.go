package main

import "fmt"
import "sync"
import "runtime"

var count int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
}

func main(){
	lock := &sync.Mutex{}

	for i:=0;i<10;i++{
		go Count(lock)
	}

	for{
		lock.Lock()
		n := count
		lock.Unlock()
		runtime.Gosched()

		if n >= 10{
			break
		}
	}
}
