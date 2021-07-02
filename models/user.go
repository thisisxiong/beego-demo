package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"strings"
	"time"
)

type User struct {
	Id int `json:"id"`
	//StructTag 校验
	Username string    `json:"username" valid:"Required;Match(/^\\w{5,18}$/)"`
	Password string    `json:"password"`
	AddTime  time.Time `json:"add_time" orm:"auto_now_add;type(date)"`
}

//数据校验
func (u *User) Valid(v *validation.Validation) {
	//reg, _ := regexp.Compile(`\w{5,18}`)
	//if b:=reg.FindString(u.Username); b == "" {
	//	v.SetError("username","用户名为5-18位字母数字下划线组成")
	//}
	if strings.Index(u.Username, "admin") != -1 {
		v.SetError("username", "用户名不能含有 admin")
	}
	//if b:=reg.FindString(u.Password); b == "" {
	//	v.SetError("password","密码为5-18位字母数字下划线组成")
	//}
}

func (User) TableName() string {
	return "user"
}

func (User) DbName() string {
	return "blog"
}

func init() {
	orm.RegisterModel(new(User))
}
