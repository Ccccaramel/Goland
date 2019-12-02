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
*/
