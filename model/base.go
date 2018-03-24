package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

func init() {

	var err error
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"Ttouch2016",
		"47.100.10.111",
		"youdian"))

	if err != nil {
		fmt.Println("Fail to init new engine: %v", err)
	}

	engine.ShowSQL(true)
}
