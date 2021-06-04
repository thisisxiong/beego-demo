package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "user/init"
	_ "user/routers"
	_ "user/utils"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
