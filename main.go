package main

import (
	"./model"
	"./service"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var userService service.UserService

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/register", userRegister)

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

func userLogin(writer http.ResponseWriter, request *http.Request) {
	//解析参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if mobile == "18600000000" && passwd == "123456" {
		loginok = true
	}

	//curl http://127.0.0.1:8989/user/login -X POST -d "mobile=18600000000&passwd=12345"
	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "success"
		Resp(writer, 0, data, "登录成功")
	} else {
		Resp(writer, 0, nil, "账号或者密码错误")
	}

}

//curl http://127.0.0.1:8989/user/register -d "mobile=13828748468&passwd=123123"
func userRegister(writer http.ResponseWriter, request *http.Request) {
	//解析参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		Resp(writer, -1, nil, err.Error())
	} else {
		Resp(writer, 0, user, "")
	}
}

type H struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header为JSON（默认是text/html）
	w.Header().Set("Content-Type", "appliction/json")
	w.WriteHeader(http.StatusOK)
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转换成JSON输出
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}
	//输出
	w.Write(ret)
}
