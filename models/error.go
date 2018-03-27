package models

import (
	"errors"
)


var (
	ErrNotExist = errors.New("实例不存在")
	ErrInsert = errors.New("没有任何修改")
)
