package main // 定义包名

import "fmt" // 告诉 Go 编译器这个程序需要用 fmt 包(的函数,或其它元素), fmt 包实现了格式化 IO(输入/输出) 的函数

func main() { // 入口函数.若有 init() 函数则会先执行 init() 函数
	fmt.Println("Hello!") // fmt.Print(...) 可以将字符串输出到控制台
	fmt.Print("Hi!" + "Golang!")
}

/*
标识符:
	由字母、数字、下划线组成,不可以以数字开头

关键字:
	break       default     func        interface   select
	case        defer       go          map         struct
	chan        else        goto        package     switch
	const       fallthrough if          range       type
	continue    for         import      return      var

预定义字符:
	append      bool        byte        cap         close       complex     complex64   complex128  uint16
	copy        false       float32     float64     imag        int         int8        int16       uint32
	int32       int64       iota        len         make        new         nil         panic       uint64
	print       println     real        recover     string      true        uint        uint8       uintptr

数据类型:
	布尔型: true/false
	数字类型: int、float32、float64
	字符串类型: Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本
	派生类型:指针类型、数组类型、结构化类型、 Channel 类型、函数类型、切片类型、接口类型、 Map 类型

数字类型:
	Go 也有基于架构的类型。例如 int 、 uint 、 uintptr
	uint8:无符号 8 位整型( 0 到 2^8-1 )
	uint16:无符号 16 位整型( 0 到 2^16-1 )
	uint32:无符号 32 位整型( 0 到 2^32-1 )
	uint64:无符号 64 位整型( 0 到 2^64-1 )
	int8:有符号 8 位整型( -2^8 到 2^7-1 )
	int16:有符号 16 位整型( -2^16 到 2^15-1 )
	int32:有符号 32 位整型( -2^32 到 2^31-1 )
	int64:有符号 64 位整型( -2^64 到 2^63-1 )

浮点型:
	float32: IEEE-754 32位浮点型数
	float64: IEEE-754 64位浮点型数
	complex64: 32 位实数和虚数
	complex128: 64 位实数和虚数
注: IEEE 二进制浮点数算术标准(IEEE 754)是 20 世纪 80 年代以来最广泛使用的浮点数运算标准,为许多 CPU 与浮点运算器所采用
*/
