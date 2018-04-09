//求两个数的最大公约数

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Printf("Enter two integer:")
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n%d & %d  gcd=%d\n", x, y, gcd(x, y))
}
