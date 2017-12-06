package main

import(
	"fmt"
)

/* 阶乘 */
func Factorial(n uint64)(result uint64){
	if(n > 0){
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

/* 斐波那契数列 */
func fibonacci(n uint64) uint64{
	if (n < 2){
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func printFibonacci(n uint64){
	var i uint64
	for i=0; i < n; i++{
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

func main(){
	var i uint64 = 5
	
	fmt.Println(i, "的阶乘=", Factorial(i))
	fmt.Println(i, "的斐波那契数列=")
	printFibonacci(i)
}


