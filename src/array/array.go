package main

import("fmt")
import("unsafe")

func main(){
	var balance [0] float32 //如果不带这个0就变成了切片类型
	//balance[0] = 3.14 //可以声明0个元素，但使用不了
	fmt.Println(balance)
	fmt.Printf("balance type=%T\n", balance)
	/*
	if(balance == nil){//切片才可以这样比较
		fmt.Printf("数组是空的\n")
	} 
	*/
	//balance = append(balance, 0)

	//a := [20]int {1,2,3,4,5}
	a := [...]int {18:1}
	fmt.Println("len=", len(a))
	fmt.Println(a)


	b := [2]int {1,2}
	c := [...]int {1,2}
	fmt.Println(c==b)
	fmt.Println(c)

	d := new([2]int)
	d[0] = 1
	d[1] = 2
	fmt.Println(d)
	//fmt.Println(c==d)


	//冒泡排序 
	arr := []int {0,3,5,2,6,7,8,2,9}

	bubble(arr, len(arr))
	
	for k:=0; k<len(arr); k++{
		fmt.Printf("%4d", arr[k])
	}
	fmt.Println()


	var as1 [5]string
	as2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//相同类型数组之间可以互相赋值
	as1 = as2
	fmt.Println(as1)
	fmt.Printf("&as1=%p, &as2=%p,&as1[0]=%p, &as2[0]=%p, sizeof(as1)=0x%X\n", &as1, &as2, &as1[0], &as2[0], unsafe.Sizeof(as1))
	//修改其中一个值会影响到另一个吗？
	as1[0] = "XXX"
	fmt.Println(as1)
	fmt.Println(as2)//值没被改动，说明他们有各自的空间
}


func bubble(array []int, n int){
	for i := 0; i < n; i++{
		for j := i+1; j < n; j++{
			if array[i] < array[j]{
				tmp := array[i]
				array[i]= array[j]
				array[j]= tmp
			}
		}
	}
	
	for k:=0; k<len(array); k++{
		fmt.Printf("%4d", array[k])
	}
	fmt.Println()
}

