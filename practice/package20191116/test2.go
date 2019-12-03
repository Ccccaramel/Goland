package main

import "fmt"

func main() {
	var s string = "Golang" // 不使用就会报错?!全局变量不会
	var num1 int
	/*
		: 等价于 var ,用于声明新的变量
		:= 是一个声明语句
		此行等价于 var num2 = 1,但无法指定变量类型
		这种不带声明格式的只能在函数体中出现
	*/
	num2 := 1
	num3, s1, bool1 := 5, "AI", true
	var num4 int = 6
	var num5 int = num4
	fmt.Println(s, &s) // 值类型的变量的值存储在栈中
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3, s1, bool1)
	fmt.Println(num4, num5)
	num4 = 7
	fmt.Println(num4, num5)
	num4, num5 = num5, num4 // 并行赋值
	var num6 int
	_, num6 = num4, num5 // _ 不会报错
	fmt.Println(num4, num5, num6)

}

/*
变量声明:
	1.var name[,name1,name2,...] type = value[,value1,value2,...]
	2.var name[,name1,name2,...] type  // 如果没有初始化,则变量默认为零值
	  name[,name1,name2,...] = value[,value1,value2,...]
	3.var name[,name1,name2,...] = value[,value1,value2,...]  // 若未指明变量类型则会根据值来判断
	4.name[,name1,name2,...] := value[,value1,value2,...]  // !慎用

全局变量声明:
var(  // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

零值:
	布尔类型: false
	数值类型(包括complex64/128): 0
	字符串: ""
	指针/数组/map/chan/func/error(接口): nil

并行赋值
	a,b = b,a

空白标识符
	 _ 被用于抛弃值,如 _,b = 1,2 中值 1 会被抛弃
	 _ 实际上是一个只写变量,你不能得到它的值
*/
