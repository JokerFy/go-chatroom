package main

import (
	"io"
	"net/http"
)

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world")
	})

	//搭建web服务器
	http.ListenAndServe(":8989", nil)
}
