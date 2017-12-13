package main

import "fmt"
import "strconv"

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name:" + p.name + "-age:" + strconv.Itoa(p.age) + "years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1       //int
	list[1] = "Hello" //string
	list[2] = Person{"hwp", 30}

	for index, element := range list {
		if v, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int, and value is:%d\n", index, v)
		} else if v, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string, and value is:%s\n", index, v)
		} else if v, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an Person, and value is:%s\n", index, v)
		} else {
			fmt.Printf("lise[%d] is of a different type\n", index)
		}
	}
}
