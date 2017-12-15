package main

import "fmt"

func foo(a int, b int, c int) {
	fmt.Println(a, b, c)
}

//func main() {
//	a := 1
//	b := 2
//	c := 3
//	defer foo(a, b, c)
//	a = 4
//	b = 5
//	c = 6
//	defer foo(a, b, c)
//}

func defer_when() int {
	i := 1

	fmt.Println("111111111111111")
	defer fmt.Println("***************")
	fmt.Println("222222222222222")

	return i

	fmt.Println("333333333333333")
	fmt.Println("444444444444444")

	return i
}
func crash() (p *int) {
	p = new(int)
	defer func() {
		*p = 1
	}()
	return nil
}

func main() {
	//defer_when()
	//crash()

	//list := new([]int)
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)

	s1:= []int{1,2,3}
	s2:= []int{4,5}
	s1 = append(s1, s2...)
	fmt.Println(s1)


}
