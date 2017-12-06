package main

import(
	"fmt"
)

func main(){
	a, b := 1, 2

	swap1(a, b)
	fmt.Println("After swap1:")
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println("After swap2:")
	fmt.Println(a, b)
}

func swap1(n1, n2 int){
	tmp := n1
	n1 = n2
	n2 = tmp
}

func swap2(n1 *int, n2 *int){
	tmp := *n1
	*n1 = *n2
	*n2 = tmp
}
