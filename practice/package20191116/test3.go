package main

import "fmt"

func main() {
	if 1 == 2 {
		fmt.Println("1 = 2")
	} else {
		fmt.Println("1 != 2")
	}

	var x interface{}
	switch i := x.(type) { // 判断某个 interface 变量中实际存储得变量类型
	case nil:
		fmt.Printf("x 的类型:%T", i)
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

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1: // 这行什么意思? ---> 通道
		fmt.Println("received", i1, " from c1")
	case c2 <- i2: // 这行又是什么意思? ---> 通道
		fmt.Println("sent ", i2, "to c2")
	case i3, ok := (<-c3): // 几个意思? ---> 通道
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

	strings := []string{"google", "golang"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 4, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	num1 := 4
	goto Loop
	num1++
	num1++
Loop:
	num1++
	fmt.Printf("num1: %d\n", num1)

	num2 := swap(&num1)
	fmt.Printf("num1: %d ,num2: %d ,&num2: %d\n", num1, *num2, num2)

	f := func() string {
		return "why"
	}
	num3 := f()
	fmt.Printf("num3: %s\n", num3)
}
func swap(num *int) *int {
	*num++
	return num
}

/*
条件语句:
	if
	if..else..
	switch 与 type switch
	select:
		* 在此之前需要了解 go 通信 ---> test4.go
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

for 循环:
	for init; condition; post{}: for 循环的青春版
	for condition{}:当 condition 条件成立时
	for{}:死循环

for-each range 循环:
	这种格式的循环可以对字符串、数组、切片等进行迭代输出元素

goto 语句:
	go 语言的 goto 语句可以无条件地转移到过程中指定的行
	goto 语句通常与条件语句配合使用,可以来实现条件转移,构成循环,跳出虚幻提等功能
	但是,在结构化程序设计中一般不主张使用 goto 语句,以免造成程序流程的混乱,使理解和调试程序都产生困难
	goto label;
	...
	label:statement;

函数:
	func function_name([parameter list])[return_types]{
		函数体
	}
	可以指定参数类型和返回值类型,以及可以返回多个值:func fun(num1,num2 int)(int, string){ return num3,str1 }
	另外, go 不支持默认参数也不支持重载

函数参数:
	值传递:浅复制
	引用传递:将参数地址传入,对其进行修改会影响实际参数

函数用法:
	函数作实参
	闭包(匿名函数)
	普通函数
*/
