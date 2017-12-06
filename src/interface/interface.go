package main

import (
	"fmt"
)

type Phone interface{
	call()
}

type NokiaPhone struct{}
type IPhone struct{}

func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call(){
	fmt.Println("I am iPhone, I can call you!")
}

func main(){
	var p Phone

	p = new(NokiaPhone)
	p.call()

	p = new(IPhone)
	p.call()
}

