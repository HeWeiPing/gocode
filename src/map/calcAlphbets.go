//统计字符串里字各个字母出现的次数

package main

import (
	"fmt"
)

func main() {
	ss := "sdkfiejtkcehdgsd"
	cal := make(map[byte]int, len(ss))

	for i := 0; i < len(ss); i++ {
		cal[ss[i]] = cal[ss[i]] + 1
	}
	for k, v := range cal {
		fmt.Printf("%c:%d\n", k, v)
	}
}
