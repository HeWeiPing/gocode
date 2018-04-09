package main

import "fmt"

var stack []int

func push(v int) []int {
	stack = append(stack, v)
	return stack
}

func pop() (v int) {
	v = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return v
}

func remove(k int) ([]int, error) {
	for i, v := range stack {
		if k == v {
			copy(stack[i:], stack[i+1:])
			return stack[:len(stack)-1], nil
		}
	}
	err := fmt.Errorf("%d not found!\n", k)
	return stack, err
}

func main() {
	for i := 0; i < 10; i++ {
		push(i)
	}
	fmt.Println(stack)

	fmt.Printf("pop:%d\n", pop())
	fmt.Println(stack)
	stack, err := remove(5)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(stack)
}
