package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int
	Username string
	Password string
	AddTime  time.Time `orm:"auto_now_add;type(date)"`
}

func (User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}
