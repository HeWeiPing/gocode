package main

import "fmt"

func sum(a []int, ch chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	for j := 0; j < 2; j++ {
		ch <- j
	}
	//ch <- sum
	//fmt.Println(len(ch))
	//ch <- 111
	//ch <- 222
	//ch <- 333
	//fmt.Println("gggg")
	//fmt.Println(len(ch))
}

func main() {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	ch := make(chan int, 2)
	fmt.Println(len(ch))

	go sum(a, ch)
	//go sum(a[:5], ch)
	//x, y := <-ch, <-ch
	//fmt.Println(x, y)
	//for {
	//if len(ch) == 0 {
	//	fmt.Println("channel is empty")
	//	break
	//}

	//fmt.Println(<-ch)
	//}
	var d int = 0
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Scanf("%d", &d)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


