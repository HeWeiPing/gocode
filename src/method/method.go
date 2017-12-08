package main

import "fmt"
import "math"

type vertex struct {
	x, y float64
}

func (v *vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
