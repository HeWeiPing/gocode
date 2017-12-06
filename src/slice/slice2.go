package main

import ("fmt")
import ("reflect")

func main(){

	//array type
	var array [10] byte
	for i:=0; i<10; i++{
		array[i] = (byte)('a' + i)
	}
	fmt.Println(array)
	//array = append(array, 'o')//会报错：first argument to append must be slice

	var array2 = [...]int{1,2,3,4,5,6,7,8,9}
	//array2 = append(array2, 111)//这样也不行，说明以上两种方式定义是数组类型

	s1 := make([]int, 3, 10)
	fmt.Println(s1)

	a := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
	a = append(a, 'o')//可以执行，说明a是切片类型
	fmt.Println(a)
	fmt.Printf("a的数据类型：%T\n", a)
	fmt.Println("a的数据类型:", reflect.TypeOf(a))
}
