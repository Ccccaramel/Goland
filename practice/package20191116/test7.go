package main

import "fmt"

type Phone interface {
	call()
}
type IPhone struct {
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you!")
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone,I can call you!")
}
func main() {
	num1 := []int{2, 3, 4}
	for k, v := range num1 {
		fmt.Printf("key: %d   value: %d\n", k, v)
	}

	num2 := map[string]string{"a": "apple", "b": "banana"}
	for _, v := range num2 {
		fmt.Printf("value: %s\n", v)
	}

	var map1 map[string]string
	map1 = make(map[string]string)

	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India"] = "新德里"

	for cou := range map1 {
		fmt.Println(cou, " 首都是 ", map1[cou])
	}

	cou, ok := map1["American"]
	if ok {
		fmt.Printf("American 的首都是 ", cou)
	} else {
		fmt.Printf("American is notyhingness\n")
	}

	delete(map1, "France")
	for cou := range map1 {
		fmt.Println(cou, " 首都是 ", map1[cou])
	}

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

/*
range():
	go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)、集合(map)的元素
	在数组和切片中它返回元素的索引和索引对应的值,在集合中返回 key-value 对

map:
	map 是一种无序的键值对的集合
	map 最重要的一点就是通过 key 来快速检索数据
	key 类似于索引,指向数据的值
	map 是一种集合,所以我们可以像迭代数组和切片那样迭代它
	不过 map 是无序的, map 是使用 hash 表实现的
	定义 map:
		可以使用内建函数 make 也可以使用 map 关键字来定义 map
			var map_name map[key_data_type]value_data_type  //默认 map 是 nil
			map_name := make(map[key_data_type]value_data_type)  //使用 make 函数

delete():
	用于删除集合的元素,参数为 map 和其对应的 key

递归:
	pass

类型转换:
	type_name(expression)

接口:
	接口也属于一种数据类型,它把所有具有共性的方法定义在一起,任何其它类型只要实现这些方法就是实现了这个接口
*/
