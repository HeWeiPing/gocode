package main

import "fmt"

const (
	a = 100
	b = 105
	c
	d = iota
	e
	f
	g = 111
	h
	i = iota
	j
)

type Student struct {
	Name string
	Age  int
}

//Student method
func (s *Student) GetName() string {
	//str := "student's name:"
	//str += s.Name
	return s.Name
}
func (s *Student) GetAge() int {
	//fmt.Println("student's :age", s.Age)
	return s.Age
}

//一个类型声明不会继承任何的基本类型的方法集合
//以下证明：
type NewStudent Student

func main() {
	fmt.Printf("after call:%d\n", aaa)

	fmt.Printf("c=%d\n", c)
	fmt.Printf("d=%d\n", d)
	fmt.Printf("e=%d\n", e)
	fmt.Printf("f=%d\n", f)
	fmt.Printf("g=%d\n", g)
	fmt.Printf("h=%d\n", h)
	fmt.Printf("i=%d\n", i)
	fmt.Printf("j=%d\n", j)

	//var nstu *NewStudent = new(NewStudent)
	//var nstu *NewStudent = &NewStudent{"hwp", 18}
	//stu := &Student{"张三", 18}
	var stu Student = Student{"hwp", 18}
	//nstu := &NewStudent{"李四", 19}
	fmt.Printf("%s\n", stu.GetName)//刚刚忘了加()，所以一直打出一个地址值，晕...
	fmt.Printf("%s\n", stu.GetName())
	//fmt.Println(nstu.GetName)//报错：没有定义GetName方法
}

var aaa int = 100
