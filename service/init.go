package service

import (
	"../model"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:root@(localhost:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}

	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(2)
	//自动创建User表结构
	DbEngin.Sync2(new(model.User))
	fmt.Println("init data base ok")
}
