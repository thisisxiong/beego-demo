package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["user/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["user/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/article/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers/admin:ArticleController"] = append(beego.GlobalControllerRouter["user/controllers/admin:ArticleController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: "/article/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers/admin:UserController"] = append(beego.GlobalControllerRouter["user/controllers/admin:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/userinfo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:ListController"] = append(beego.GlobalControllerRouter["user/controllers:ListController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/article/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:LoginController"] = append(beego.GlobalControllerRouter["user/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:LoginController"] = append(beego.GlobalControllerRouter["user/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "IsExist",
            Router: "/cache/exist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/cache/get",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "Set",
            Router: "/cache/set",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "Context",
            Router: "/context/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "HttpGet",
            Router: "/http/get",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "HttpPost",
            Router: "/http/post",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user/controllers:TestController"] = append(beego.GlobalControllerRouter["user/controllers:TestController"],
        beego.ControllerComments{
            Method: "Output",
            Router: "/logs",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
