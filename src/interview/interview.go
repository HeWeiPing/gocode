package main

import "fmt"

import "sync"

//import "runtime"

/*111111111111111111111111111111111111111111111*/
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

//func main() {
//	//defer_call()
//}
/*111111111111111111111111111111111111111111111*/

/*2222222222222222222222222222222222222222222222*/
type student struct {
	Name string
	age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", age: 24},
		{Name: "li", age: 23},
		{Name: "wang", age: 22},
	}
	/* 下面这种赋值方法是错误的
	   因为stu是for循环的短声明作用域
	   每次循环都是重用了上一次的定义
	   所以最后的map的所有value值都是
	   stu变量最后一次指向的指针，即：
	   {"wang", 22} */
	/*
		for _, stu := range stus {
			fmt.Println(stu)
			fmt.Printf("%p\n", &stu)
			m[stu.Name] = &stu
		}
	*/
	/* 正确的方法如下：*/
	for i, v := range stus {
		fmt.Printf("&i:%p, &v:%p\n", &i, &v)
		fmt.Println(i, ":", v)
		m[v.Name] = &stus[i]
		fmt.Printf("&m[%d]%p\n", i, m[v.Name])
	}
	for k, v := range m {
		fmt.Printf("%s:%s:%d\n", k, v.Name, v.age)
	}
}

//func main() {
//	//pase_student()
//}
/*2222222222222222222222222222222222222222222222*/

/*3333333333333333333333333333333333333333333333*/
/*
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	//ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("for1:&i:%p, i=%d\n", &i, i)
			wg.Done()
			//ch<-1
		}()
		//<-ch
		//runtime.Gosched()
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("outside:&i:%p, i=%d\n", &i, i)
		go func(i int) {//局部作用域，会重新分配地址，进行值拷贝传递
			fmt.Printf("for2:&i:%p, i=%d\n", &i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
*/

/*3333333333333333333333333333333333333333333333*/

/*4444444444444444444444444444444444444444444444*/
/*
type People struct {}
func (p *People)ShowA(){
	fmt.Println("showA")
	p.ShowB()
}
func (p *People)ShowB(){
	fmt.Println("showB")
}
type Teacher struct{
	People
}
func (t *Teacher)ShowB(){
	fmt.Println("teacher showB")
}
func main(){
	t := Teacher{}
	t.ShowA()
}
*/
/*4444444444444444444444444444444444444444444444*/

/*5555555555555555555555555555555555555555555555*/
/*
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"

	select {
	case v := <-int_chan:
		fmt.Println(v)
	case v := <-string_chan:
		panic(v)
	}
}
*/
/*5555555555555555555555555555555555555555555555*/

/*6666666666666666666666666666666666666666666666*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//func main(){
//	a := 1
//	b := 2
//	defer calc("1", a, calc("10", a, b))
//	a = 0
//	defer calc("2", a, calc("20", a, b))
//	b = 1
//}
/*
最开始我认为会这样输出：
20 0 1 1
2 0 1 1
10 0 1 1
1 0 1 1
结果是：
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
*/
/*6666666666666666666666666666666666666666666666*/

/*7777777777777777777777777777777777777777777777*/
//请写出以下输入内容
//func main() {
//	s := make([]int, 5)
//	s = append(s, 1, 2, 3)
//	fmt.Println(s)
//}
/*7777777777777777777777777777777777777777777777*/

/*8888888888888888888888888888888888888888888888*/
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (u *UserAges) Add(name string, age int) {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}
func (u *UserAges) Get(name string) int {
	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}

//func main(){
//	pu := &UserAges{make(map[string]int), sync.Mutex{}}
//	pu.Add("hwp", 18)
//	fmt.Println(pu)
//	age := pu.Get("hwp")
//	fmt.Println(age)
//}
/*8888888888888888888888888888888888888888888888*/

/*9999999999999999999999999999999999999999999999*/
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{})
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		defer set.RUnlock()
		for elem, value := range set.s {
			ch <- value
			fmt.Println("Iter:", elem, value)
		}
		close(ch)
	}()

	return ch
}

//func main() {
//	th := threadSafeSet{s: []interface{}{"aaa", "bbb"}}
//
//	ch := th.Iter()
//	for i := 0; i < len(th.s); i++ {
//		fmt.Printf("%s%v\n", "main r:", <-ch)
//	}
//}

/*9999999999999999999999999999999999999999999999*/

/*10 10 10 10 10 10 10 10 10 10 10 10 10 10 10 10*/
//type People interface {
//	Speak(string) string
//}
//
//type Student struct{}
//
//func (stu *Student) Speak(think string) (talk string) {
//	if think == "bitch" {
//		talk = "You are a good boy"
//	} else {
//		talk = "hi"
//	}
//	return talk
//}
//func main() {
//	//这种写法不可以，因为你是用 *Student 指针类型为接收者参数
//	//而不是Student类型为接收者参数(Speak method has pointer receiver)
//	//var p People = Student{}
//	var p People = &Student{}
//	think := "bitch"
//	fmt.Println(p.Speak(think))
//}

/*10 10 10 10 10 10 10 10 10 10 10 10 10 10 10 10*/

/* 11 11 11 11 11 11 11 11 11 11 11 11 11 11 11 11 */
//type People interface {
//	Show()
//}
//
//type Student struct{}
//
//func (stu *Student) Show() {}
//
//func live() People {
//	var stu *Student
//
//	return stu
//}
//
//func main() {
//	if live() == nil {
//		fmt.Println("aaaaa")
//	} else {
//		fmt.Println("bbbbb")
//	}
//}

/* 11 11 11 11 11 11 11 11 11 11 11 11 11 11 11 11 */

/* 12  12  12  12  12  12  12  12  12  12  12  12 */
//func GetValue()int{
//	return 1
//}
func GetValue() interface{}{
	//return 1
	return "abc"
}
func main(){
	i := GetValue()

	switch i.(type){//type 只能用在interface上
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("Unknown")
	}
}
/* 12  12  12  12  12  12  12  12  12  12  12  12 */
