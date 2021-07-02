package models

import "github.com/beego/beego/v2/client/orm"

type UserLevel struct {
	Id          int
	Name        string
	Description string
	Img         string
	Status      int8
	CreateTime  int
	UpdateTime  int
	DeleteTime  int
}

func (u *UserLevel) TableName() string {
	return "user_level"
}

func (u *UserLevel) DbName() string {
	return "beego-admin"
}

func init() {
	orm.RegisterModel(new(UserLevel))
}
