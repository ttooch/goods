package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

		panic(fmt.Sprintf("Fail to init new engine: %v", err))

	}

	Engine.ShowSQL(true)
}

type Model interface {
	AfterFind()
}

func GetModel(model Model, session *xorm.Session) error {
	has, err := session.Get(model)

	model.AfterFind()

	if err != nil {
		return err
	} else if !has {
		return ErrNotExist
	}

	return nil
}

func AddModel(model *Goods) error {
	effect, err := Engine.InsertOne(model)

	if err != nil {
		return err
	} else if effect == 0 {
		return ErrInsert
	}
	return nil
}

func UpdateModel(model *Goods, session *xorm.Session) error {
	_, err := session.Update(model)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func DelModel(model Goods, session *xorm.Session, safe ...bool) (err error) {

	if len(safe) > 0 && safe[0] == false {
		_, err = session.Unscoped().Delete(&model)
	} else {
		_, err = session.Delete(&model)
	}

	return err
}
