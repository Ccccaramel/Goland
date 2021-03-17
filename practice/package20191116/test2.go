package main

import (
	"fmt"
	"reflect"
)

const (
	IMAGE_PATH = "F:\\dada\\image"
	SOUND_PATH = "F:\\dada\\sound"
	STATE      = true
)

func main() {
	var v1 *int
	var v2 []int
	var v3 map[string]int
	var v4 chan int
	var v5 func(string) int
	var v6 error
	fmt.Println(v1, v2, v3, v4, v5, v6)
	fmt.Printf("%T", reflect.TypeOf(v4))

	var s1 string = "Golang" // 不使用就会报错?!全局变量不会
	var num1 int
	/*
		: 等价于 var ,用于声明新的变量
		:= 是一个声明语句使代码更加简洁
		这种形式称作 简短声明
		此行等价于 var num2 = 1
		这种不带声明格式的只能在函数体中出现
		故 var 适合定义全局变量
		而 ：= 适合定义函数内部变量
	*/
	num2 := 1
	num3, s2, bool1 := 5, "AI", true
	var num4 int = 6
	var num5 int = num4
	var num6 int

	fmt.Println(s1, &s1) // 值类型的变量的值存储在栈中
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3, s2, bool1)
	fmt.Println(num4, num5)
	num4 = 7
	fmt.Println(num4, num5)
	num4, num5 = num5, num4 // 并行赋值

	_, num6 = num4, num5 // _ 不会报错
	fmt.Println(num4, num5, num6)

	const PI = 3.14
	const (
		a = iota      // 0
		b = 4 << iota // iota += 1,4 -> 100,<< 左移 iota 位,1000 -> 8
		c             // 2
		d = "ha"      // 独立值, iota += 1
		e             // 未赋值则与上一个常量等值,所以第一个常量必须赋值."ha", iota += 1
		f = 100       // iota += 1
		g             // 100, iota += 1
		h = iota      // 7,恢复计数
		i             // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

}

/*
变量声明:
	1.var name[,name1,name2,...] type = value[,value1,value2,...]
	2.var name[,name1,name2,...] type  // 如果没有初始化,则变量默认为零值
	  name[,name1,name2,...] = value[,value1,value2,...]
	3.var name[,name1,name2,...] = value[,value1,value2,...]  // 若未指明变量类型则会根据值来判断
	4.name[,name1,name2,...] := value[,value1,value2,...]

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

常量:
	1.const name[,nam1,name2,...] type = value[value1,value2,...]  // 显示类型定义
	2.const name[,nam1,name2,...] = value[value1,value2,...]  // 隐式类型定义

iota:
	特殊常量,可以认为是一个可以被编译器修改的常量
	iota 在 const 关键字出现时将被重置为 0( const 内部的第一行之前)
	const 中每新增一行常量声明将使 iota 计数一次( iota 可理解为 const 语块中的行索引)

内置运算符:
	算数运算符: + - * / % ++ --
	关系运算符: == != > < >= <=
	逻辑运算符: && || !
	位运算符: & | ^ << >>
	赋值运算符: = += -= *= /= %= <<= >>= &= ^= \=
	其他运算符: &(返回变量存储地址) *(指针变量)

*/
