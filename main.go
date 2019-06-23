package main

import (
	"./ctrl"
	"./service"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

var userService service.UserService

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/chat", ctrl.Chat)
	//提供静态资源文件
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	//template处理
	RegisterView()

	//搭建web服务器
	http.ListenAndServe(":8989", nil)
}

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}
