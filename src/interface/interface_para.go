package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) String() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + "years - phone_number:" + h.phone + ">"
}

func main() {
	Bob := Human{"Bob", 32, "15330273585"}
	fmt.Println("This human is:", Bob)
}
