package main

import "fmt"
import "os"

func main() {
	f, err := os.Open("log.log")
	fmt.Printf("err1:%p\n", &err)
	if err != nil {
		fmt.Printf("%p", err)
	}
	d, err := os.Stat("log.log")//此时的err变量重用了7行定义的err，只是对其重新赋值而已
	fmt.Printf("err2:%p\n", &err)//与8行打印的地址一样
	if err != nil {
		fmt.Println(d)
		f.Close()
		fmt.Printf("%p", err)
	}

	for i := 0; i < 2; i++ {
		fmt.Printf("i1:%p\n", &i)
	}
	for i := 0; i < 2; i++ {
		fmt.Printf("i2:%p\n", &i)
	}

}
