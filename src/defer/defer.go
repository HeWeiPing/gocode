package main

import "fmt"

func f1(a *int, b *int, c *int){
	fmt.Printf("&a:%p a=%d &b:%p b=%d &c:%p c=%d\n", a, *a, b, *b, c, *c)
}
func f2(a int, b int, c int){
	fmt.Printf("&a:%p a=%d &b:%p b=%d &c:%p c=%d\n", &a, a, &b, b, &c, c)
}


func main() {
	//上来就使用函数作用域的变量是不可以的
	//所以局部作用域的变量必须先声明后使用
	//defer fmt.Println(a, b, c)
	a := 1
	b := 2
	c := 3
	//defer fmt.Println("aaa:", a, b, c)
	fmt.Printf("&a:%p a=%d &b:%p b=%d &c:%p c=%d\n", &a, a, &b, b, &c, c)
	defer f1(&a, &b, &c)
	//defer f2(a, b, c)
	a = 4
	b = 5
	c = 6
	//defer fmt.Println("bbb:", a, b, c)
	//defer p(a, b, c)
	fmt.Printf("&a:%p a=%d &b:%p b=%d &c:%p c=%d\n", &a, a, &b, b, &c, c)
	defer f1(&a, &b, &c)
	//defer f2(a, b, c)
}
