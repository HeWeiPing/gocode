package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tmp := make(map[string]int)
	fmt.Println("Enter your text:")
	inputs := bufio.NewScanner(os.Stdin)
	//Set the split function for the scanning operation.
	inputs.Split(bufio.ScanWords)
	for inputs.Scan() {
		tmp[inputs.Text()]++
	}

	for s, n := range tmp {
		fmt.Printf("%s\t%d\n", s, n)
	}
}
