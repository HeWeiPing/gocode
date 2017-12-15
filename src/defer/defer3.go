package main

import "fmt"

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

//对下面这个坑的理解：
//按网上据看过官方文档的网友解释，return 的值是要被压栈，
//然后去执行defer func，因为是无名接收返回值，所以在defer
//语句执行前return 的值已经保存即：return value = i已经入栈
//有可能是保存了一个立即数。而此时defer func 里去修改t对
//return velue来说没什么关系了
func DeferFunc2(i int) int {
	t := i
	fmt.Printf("f2,outside closure:&t=%p, t=%d\n", &t, t)
	defer func() {
		t += 3
		fmt.Printf("f2, inside closure:&t=%p, t=%d\n", &t, t)
	}()
	fmt.Println("after fun:t=", t)
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
func main() {
	DeferFunc1(1)
	DeferFunc2(1)
	DeferFunc3(1)
}
