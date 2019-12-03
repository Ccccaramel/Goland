package main

import (
	"fmt"
)

const (
	IMAGE_PATH = "F:\\dada\\image"
	SOUND_PATH = "F:\\dada\\sound"
	STATE      = true
)

func main() {
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

	if 1 == 2 {
		fmt.Println("1 = 2")
	} else {
		fmt.Println("1 != 2")
	}

	var x interface{}
	switch i := x.(type) { // 判断某个 interface 变量中实际存储得变量类型
	case nil:
		fmt.Println("x 的类型:%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("unknown")
	}
	switch {
	case false:
		fmt.Println("1-false")
		fallthrough // 继续判断下一个
	case true:
		fmt.Println("2-true")
		fallthrough
	default:
		fmt.Println("7-unknown")
	}
}

/*
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

条件语句:
	if
	if..else..
	switch 与 type switch
	select:
		select 是 Go 中的一个控制结构,类似于用于通信得 Switch 语句
		每个 case 必须是一个通信操作,要么是发送要么是接收
		select 随即执行一个可运行的 case
		如果没有 case 可运行它将阻塞直到有 case 可运行
		一个默认的子句应该总是可运行的
		语法:
			每个 case 都必须是一个通信
			所有 channel 表达式都会被求值
			所有被发送的表达式都会被求值
			如果任意某个通信可以进行,它就执行,其它被忽略
			如果有多个 case 都可以运行, select 会随机公平地选出一个执行,其它不执行否则:
				1.如果有 default 子句,则执行该语句
				2.如果没有 select 将阻塞直到某个通信可以运行; Go 不会重新对 channel 或值进行求值
*/
