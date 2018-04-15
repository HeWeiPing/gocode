package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	x int
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

//匿名嵌入的特性可以使结构变量直接访问不重名的字段
//而不需要给出完整的路径

//**如果有相同名字字段的子结构体，则会优先同赋值最外层
func main() {
	var w Wheel

	w.Point.x = 1 //如果不完整指明Point的x，则只赋值给Wheel的x
	//w.x = 1
	w.y = 2
	w.Radius = 3
	w.Spokes = 4

	fmt.Printf("%#v\n", w) //{{1 {0 2} 3} 4}
}
