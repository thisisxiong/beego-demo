package admin

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"user/models"
)

type UserController struct {
	beego.Controller
}

// @Title GetAll
// @Description 获取用户列表
// @Success 200
// @router /userinfo  [get]
func (uc *UserController) GetAll() {
	var user []models.User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&user, "id", "username")
	if err != nil {
		uc.Data["json"] = map[string]interface{}{
			"code":    400,
			"data":    err,
			"message": err,
		}
		uc.ServeJSON()
		return
	}

	uc.Data["json"] = map[string]interface{}{
		"code":    200,
		"data":    user,
		"message": user,
	}
	uc.ServeJSON()
	return
}
