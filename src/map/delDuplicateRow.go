//删除重复的行

package main

import "fmt"

func main() {
	//md := make(map[int]string)
	md := map[int]string{
		1: "123",
		2: "456",
		3: "789",
		4: "456",
		5: "789",
		6: "123",
	}

	fmt.Println(md)

	tmp := make(map[string]int, len(md))

	for k, v := range md {
		tmp[v] = k
		fmt.Printf("tmp[%s]=%d\n", v, k)
	}
	for k, v := range tmp {
		fmt.Println(k, v)
	}
}
