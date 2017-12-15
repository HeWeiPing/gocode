package main

import "fmt"

func main() {
	s := 1

	fmt.Printf("outsdide if:&s=%p\n", &s)

	if s, ok := 8, true; ok {
		fmt.Printf("insdide if:&s=%p\n", &s)
		fmt.Printf("insdide if:s=%d\n", s)
	}

	fmt.Printf("outsdide if:s=%d\n", s)
}
