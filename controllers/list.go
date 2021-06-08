package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"user/models"
	"user/utils"
)

type ListController struct {
	beego.Controller
}

// @Title List
// @Description 文章列表
// @Success 200
// @router /article/list [post]
func (this *ListController) List() {
	var article []models.Article
	o := orm.NewOrm()
	_, err := o.QueryTable("article").RelatedSel().All(&article)
	if err != nil {
		utils.ToJson(this.Controller, err.Error(), "读取失败", "400")
		return
	}
	utils.ToJson(this.Controller, article, "", "200")
	return

}
