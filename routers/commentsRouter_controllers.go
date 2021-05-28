package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["user/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["user/controllers/admin:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/userinfo",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["user/controllers:LoginController"] = append(beego.GlobalControllerRouter["user/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["user/controllers:LoginController"] = append(beego.GlobalControllerRouter["user/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           "/register",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
