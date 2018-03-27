package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
)

func init() {

	var err error
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"goods",
		"Ttouch2016",
		"47.100.10.111",
		"goods"))

	if err != nil {
		fmt.Printf("Fail to init new engine: %v", err)
	}

	Engine.ShowSQL(true)
}
