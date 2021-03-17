package main

import "fmt"

func main() {
	var balance [10]float32
	//var balance [] float32  也可以不设 size , go 会根据元素个数设置大小
	balance = [10]float32{12.0, 23.0, 34.0, 45.0, 56.0, 67.0, 78.0, 89.0, 90, 100.0}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d:%f\n", i, balance[i])
	}

	num1 := 20
	var ip *int
	ip = &num1
	fmt.Printf(" num1 的变量地址: %x\n ip 变量储存的指针地址: %x\n使用指针 *ip 访问地址内的值: %d\n", &num1, ip, *ip)

	arr1 := []int{10, 100, 1000}
	var i int
	var ptr [3]*int
	for i = 0; i < 3; i++ {
		ptr[i] = &arr1[i] //把地址赋值给指针数组
	}
	for i = 0; i < 3; i++ {
		fmt.Printf("arr[%d] = %d\n", i, *ptr[i])
	}

	var num2 int
	var ptr1 *int
	var pptr **int
	num2 = 3000
	ptr1 = &num2
	pptr = &ptr1

	fmt.Printf("变量 num2 = %d\n:", num2)
	fmt.Printf("指针变量 *ptr1 = %d\n:", *ptr1)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n:", **pptr)

	num3 := 12
	num4 := 34
	swap1(&num3, &num4)
	fmt.Printf("num3: %d ,num4:%d\n", num3, num4)
}
func swap1(a *int, b *int) {
	*a, *b = *b, *a
}

/*
作用域:
	局部变量:函数内定义的变量 :=
	全局变量:函数外定义的变量 var
	形式参数:函数定义中的变量

数组:
	pass

向函数传递数组:
	void fun(param [10]int){}
	void fun(param []int){}

指针:
	pass

空指针:
	它的值为 nil
	nil 指针也称为空指针
	一个指针变量通常缩写为 ptr

指针数组:
	pass

指针指向指针:
	一个指针变量存放的又是另一个指针变量的地址,则称这个指针变量为指向指针的指针变量
	var ptr **int

指针作为函数参数:
	pass
*/
