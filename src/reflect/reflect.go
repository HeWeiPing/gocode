package main

import "fmt"
import "reflect"


func main(){
	var x float64 = 3.14

	v := reflect.ValueOf(x)

	fmt.Println(v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
