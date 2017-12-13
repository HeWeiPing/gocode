package main

import "fmt"
import "time"

//Switch 在每一个case后面默认加上了break，不用担心往后执行

func main(){
	i := 2
	fmt.Println("i=", i)
	switch i{
	case 1:
		fmt.Println("i is one")
	case 2:
		fmt.Println("i is two")
	case 3:
		fmt.Println("i is three")
	default:
		fmt.Println("unkown value")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend -->", time.Now().Weekday())
	default:
		fmt.Println("It's the weekday -->", time.Now().Weekday())
	}

	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon:", t.Hour())
	default:
		fmt.Println("It's after noon:", t.Hour())
	}

	whatAmI := func (i interface{}){
		switch t:=i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type:%T\n", t)
		}
	}
	whatAmI(1)
	whatAmI(true)
	whatAmI("aaaaaa")
}
