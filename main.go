package main

import (
	"net/http"
)

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		//解析参数
		request.ParseForm()
		mobile := request.PostForm.Get("mobile")
		passwd := request.PostForm.Get("passwd")

		loginok := false
		if mobile == "18600000000" && passwd == "123456" {
			loginok = true
		}

		//curl http://127.0.0.1:8989/user/login -X POST -d "mobile=18600000000&passwd=12345"
		str := `{"code":0,"data":{"id":1,"token":'test'}}`
		if !loginok {
			str = `{"code":-1,"msg":"123123"}`
		}
		//设置header为JSON（默认是text/html）
		writer.Header().Set("Content-Type", "appliction/json")
		writer.WriteHeader(http.StatusOK)
		//输出
		writer.Write([]byte(str))
	})

	//搭建web服务器
	http.ListenAndServe(":8989", nil)
}
