package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
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
