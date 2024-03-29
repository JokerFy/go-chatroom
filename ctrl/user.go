package ctrl

import (
	"../model"
	"../service"
	"../util"
	"fmt"
	"math/rand"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	//解析参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	//curl http://127.0.0.1:8989/user/login -X POST -d "mobile=18600000000&passwd=12345"
	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}

//curl http://127.0.0.1:8989/user/register -d "mobile=13828748468&passwd=123123"
var userService service.UserService

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	//解析参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}
