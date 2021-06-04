package admin

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"user/models"
	"user/utils"
)

type ArticleController struct {
	beego.Controller
}

// @Title Create
// @Description 发布文章
// @Success 200
// @router /article/add  [post]
func (a *ArticleController) Create() {
	article := models.Article{}
	user := models.User{}
	user.Id, _ = a.GetInt("user_id")
	article.User = &user
	err := a.ParseForm(&article)
	if err != nil {
		a.Data["json"] = map[string]interface{}{
			"code" : 400,
			"data" : err,
			"message" : "接受参数失败",
		}
		a.ServeJSON()
		return
	}
	orm.Debug = true
	o := orm.NewOrm()
	_, err = o.Insert(&article)
	if err != nil {
		a.Data["json"] = map[string]interface{}{
			"code" : 400,
			"data" : err.Error(),
			"message" : article,
		}
		a.ServeJSON()
		return
	}
	a.Data["json"] = map[string]interface{}{
		"code" : 200,
		"data" : article,
		"message" : "发布成功",
	}
	a.ServeJSON()
	return
}

// @Title Edit
// @Description 更新文章
// @Success 200
// @router /article/update  [post]
func (a *ArticleController) Edit() {
	id, _ := a.GetInt("id")
	if id == 0 {
		utils.ToJson(a.Controller,"参数错误","文章id有误","400")
		return
	}
	o := orm.NewOrm()
	article := models.Article{
		Id: id,
	}
	if o.Read(&article) == nil {
		err := a.ParseForm(&article)
		if err != nil {
			utils.ToJson(a.Controller,err.Error(),"参数错误","400")
			return
		}
		_, err = o.Update(&article)
		if err != nil {
			utils.ToJson(a.Controller,err.Error(),"更新失败","400")
			return
		}

		utils.ToJson(a.Controller,article,"更新成功","400")
		return
	}
}