package main

import "fmt"

//part 2 每次循环defer都会压栈一次无名函数的实例
//也就是函数入口地址，但没有参数，所以没有参数压栈
//直到for range执行完毕，defer开始Pop压栈的闭包子函数
//此时需要访问子函数外的局部变量i，而此时i已经变为4
//所以会打印：4 4 4 4 4

//part 3 每次defer压栈时，当前运行环境下的上下文i的值
//也会被一起压栈，所以等到pop函数时，函数是带着当运行环境
//下的i的值为参数执行子函数，所以会打印: 4 3 2 1 0

func main() {
	var whatever [5]struct{}

	for i := range whatever {
		fmt.Println("p1:", i)
	} // part 1

	for i := range whatever {
		defer func() { fmt.Println("p2:", i) }()
	} // part 2

	for i := range whatever {
		defer func(n int) { fmt.Println("p3:", n) }(i)
	} // part 3
}
