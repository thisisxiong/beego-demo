// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"user/controllers"
	"user/controllers/admin"
	"user/middlerware"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/admin",
			//路由中间件 用于验证登录
			beego.NSBefore(middlerware.AdminAuth),
			beego.NSInclude(
				&admin.UserController{},
			),
		),
		beego.NSInclude(
			&controllers.LoginController{},
		),
	)
	beego.AddNamespace(ns)
}
