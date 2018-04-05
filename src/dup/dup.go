package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tmp := make(map[string]int)
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		tmp[inputs.Text()]++
	}

	for s, n := range tmp {
		//if n > 1 {
		fmt.Printf("%s\t%d\n", s, n)
		//}
	}
}
