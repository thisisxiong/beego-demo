package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"user/utils"
)

func dbHandler(rw http.ResponseWriter, r *http.Request) {
	//rw.WriteHeader(http.StatusBadRequest)
	rw.Header().Set("errMsg", "db service failed")
	rw.Write([]byte("db error msg"))

}

func init() {
	web.ErrorHandler("db", dbHandler)
}

type ErrorController struct {
	web.Controller
}

func (c *ErrorController) Error404() {
	utils.ToJson(c.Controller, "", "未找到", 404)
	return
}

func (c *ErrorController) Error500() {
	utils.ToJson(c.Controller, "", "服务器错误", 500)
	return
}

func (c *ErrorController) Error403() {
	utils.ToJson(c.Controller, "", "未登录", 403)
	return
}

func (c *ErrorController) Error408() {
	utils.ToJson(c.Controller, "", "自定义错误信息", 200)
	return
}

//func (c *ErrorController) Errordb() {
//	utils.ToJson(c.Controller,"","数据库错误",400)
//	return
//}
