package main

import(
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func main(){
	fmt.Println(a, b, c)

	const(
		e = iota
		f
		g
		h = "haha"
		i
		j = 100
		k
		l = iota
		m
		n
	)
	fmt.Println(e, f, g, h, i, j, k, l, m, n)

	const(
		a1 = 1<<iota //1
		a2 = 3<<iota //6
		a3           //12
		a4           //24
	)
	fmt.Println(a1, a2, a3, a4)

}
