package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:123456@(127.0.0.1:3306)/chat?charset=utf8"
	DbEngin, err := xorm.NewEngine(drivename, DsName)
	if err != nil {
		log.Fatal(err.Error())
	}

	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(2)
	//DbEngin.sync2(new(User))
	fmt.Println("init data base ok")
}
