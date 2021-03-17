package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book *Books) {
	fmt.Printf("title: %s ;author: %s ;author: %s ;book_id: %d\n", book.title, book.author, book.subject, book.book_id)
}
func main() {

	fmt.Println(Books{"Go 语言", "https://www.github.com/Ccccaramel/Golang", "Go 语言教程", 164})
	fmt.Println(Books{title: "Go 语言", author: "https://www.github.com/Ccccaramel/Python", subject: "Python 语言教程", book_id: 159})
	fmt.Println(Books{title: "Go 语言", author: "https://www.github.com/Ccccaramel"}) // 忽略字段默认置为 0 或 空

	var Book1 Books
	Book1.title = "C"
	Book1.author = "https://www.github.com/Ccccaramel/C"
	Book1.subject = "C 教程"
	Book1.book_id = 560
	fmt.Printf("Book1 title: %s   Book1 author: %s   Book1 author: %s   Book1 book_id: %d\n", Book1.title, Book1.author, Book1.subject, Book1.book_id)

	printBook(&Book1)

	var number1 = make([]int, 3, 5)
	fmt.Printf("number1: len = %d cap = %d slice = %v\n", len(number1), cap(number1), number1)

	var number2 []int
	if number2 == nil {
		fmt.Println("number2 is nil")
	}
	fmt.Printf("number2: len = %d cap = %d slice = %v\n", len(number2), cap(number2), number2)

	number2 = append(number2, 0)
	fmt.Printf("number2: len = %d cap = %d slice = %v\n", len(number2), cap(number2), number2)
	number2 = append(number2, 1, 2, 3, 4)
	fmt.Printf("number2: len = %d cap = %d slice = %v\n", len(number2), cap(number2), number2)
}

/*
结构体:
	访问成员使用 .

切片:
	go 语言切片是对数组的抽象
	go 数组的长度不可改变,在特定场景中这样的集合就不太适用, go 中提供了·一种灵活,功能强悍的内置类型切片("动态数组")
	与数组相比切片的长度是不固定的,可以追加元素,在追加时可能使切片的容量增大
	定义切片:
		你可以声明一个未指定大小的数组来定义切片:
			var identifier []type
		切片不需要说明长度,或使用 make() 函数来创建切片
			var slice []type = make([]type, len)
			slice := make([]type, len)
		也可以指定容量,其中 capacity 为可选参数
			make([]type, len, capacity)
			len 是数组的长度也是切片的初始长度
	切片初始化:
		直接初始化切片,[]表示是切片类型,{1，2，3}初始化值依次是1,2,3.其中 cap = len = 3
			s := [] int {1, 2, 3}
		初始化切片 s ,是数组 arr 的引用
			s := arr[:]
		将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片
			s := arr[startIndex:endIndex]
		默认 endIndex 时将表示一直到 arr 的最后一个元素
			s := arr[startIndex:]
		默认 startIndex 时将表示从 arr 的第一个元素开始
			s := arr[:endIndex]
		通过切片 s 初始化切片 s1
			s1 := s[startIndex:endIndex]
		通过内置函数 make() 初始化切片 s ,[]int 标识为其元素类型为 int 的切片
			s := make([]int, len, cap)
		len() 和 cap() 函数:
			切片是可索引的,并且可以由 len() 方法获取长度
			切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
		空(nil)切片:
			一个切片在未被初始化之前默认为 nil ,长度为 0
		切片截取:
			pass
		append() 与 copy() 函数:
			pass
*/
