package main

import ("fmt")
import ("reflect")

func main(){
	s1 := make([]int, 3, 10)
	fmt.Println(s1)

	a := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
	fmt.Printf("a的数据类型：%T\n", a)
	fmt.Println("a的数据类型:", reflect.TypeOf(a))
	sa := a[2:5]
	a = append(a, 'o')
	sa = append(sa, 'o')
	fmt.Println("sa的数据类型:", reflect.TypeOf(sa))
	fmt.Println(len(sa), cap(sa))
	sb := sa[3:5]
	fmt.Println(string(sb))

	s2 := []int{1,2,3,4,5}
	sc := append(s2,6,7,8)
	fmt.Println("s2:",s2, "sc:",sc)
	fmt.Printf("&s2:%p &sc:%p\n", s2, sc)

	ss1 := []int{1,2,3,4,5,6}
	ss2 := []int{7,8,9,0}
	fmt.Println("before copy:",ss1, ss2)
	copy(ss1,ss2)
	fmt.Println("after copy :",ss1, ss2)

	saa := []int{1,2,3,4,5,6}
	saa1 := saa[2:5]
	saa2 := saa[1:3]
	fmt.Println("saa:", saa)
	fmt.Println("saa1:", saa1, "saa2:", saa2)
	//saa1[0] = 9
	fmt.Println("saa1:", saa1, "saa2:", saa2)
	saa2 = append(saa2, 1,1,1,1,1,1,1,1,1)
	saa1[0] = 9
	fmt.Println("saa1:", saa1, "saa2:", saa2)

	//简洁slice拷贝写法
	sb1 := []int{1,2,3,4,5}
	sb2 := sb1[:]
	fmt.Println(sb1, sb2)
	fmt.Println(reflect.TypeOf(sb1))//打印变量的数据类型 
}
