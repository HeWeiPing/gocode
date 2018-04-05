//统计从标准输入的各个字符出现的次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tmp := make(map[byte]int)
	//buf := make([]byte, 10)
	fmt.Println("Enter your text:")
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		buf := s.Bytes()
		fmt.Println("***********")

		for _, v := range buf {
			tmp[v]++
		}
	}
	for i, v := range tmp {
		fmt.Printf("%c\t%d\n", i, v)
	}
}
