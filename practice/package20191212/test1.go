package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb) //注册默认路由,第一个参数为请求路径,第二个参数为函数类型,表示这个请求需要处理的事情,,,
	fmt.Println("服务器即将开启,访问地址 http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误:", err)
	}
}
