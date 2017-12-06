package main

import (
	"fmt"
)

type Book struct{
	title	string
	author	string
	subject	string
	book_id	string
}
func main(){
	var Book1 Book
	var Book2 Book

	/* book1 描述 */
	Book1.title	= "Go 语言"
	Book1.author = "HWP"
	Book1.subject = "Go 语言教程"
	Book1.book_id = "1234567890"

	/* book2 描述 */
	Book2.title	= "Python 语言"
	Book2.author = "HWP"
	Book2.subject = "Python 语言教程"
	Book2.book_id = "9876543210"

	/* 打印 Book1 信息 */
	fmt.Printf("Book1 title: %s\n", Book1.title)
	fmt.Printf("Book1 author: %s\n", Book1.author)
	fmt.Printf("Book1 subject: %s\n", Book1.subject)
	fmt.Printf("Book1 book_id: %s\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book2 title: %s\n", Book2.title)
	fmt.Printf("Book2 author: %s\n", Book2.author)
	fmt.Printf("Book2 subject: %s\n", Book2.subject)
	fmt.Printf("Book2 book_id: %s\n", Book2.book_id)

	//struct 作为参数
	fmt.Println("Struct paramter:")
	printBook(Book1)
	printBook(Book2)
}


func printBook(b Book){
	/* 打印 Book 信息 */
	fmt.Printf("Book title: %s\n", b.title)
	fmt.Printf("Book author: %s\n", b.author)
	fmt.Printf("Book subject: %s\n", b.subject)
	fmt.Printf("Book book_id: %s\n", b.book_id)
}
