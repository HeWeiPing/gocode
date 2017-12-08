package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	//ok
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	//Not ok
	z := make(bar)
}
