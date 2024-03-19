package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Username      string
	Password      string
	NickName      string
	Status        string
	LastLoginTime time.Time
}

type Role struct {
	gorm.Model
	Name   string
	Remark string
	Status string
}

type SysLog struct {
	Operator     string
	OperatorTime time.Time
	Status       string
	Desc         string
	Content      string
	Response     string
}
