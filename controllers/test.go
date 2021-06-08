package controllers

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"time"
	"user/utils"
)

type TestController struct {
	beego.Controller
}

var cacheObj cache.Cache

func init() {
	//目前支持 file、memcache、memory 和 redis 四种引擎
	redisConf, err := config.NewConfig("ini", "conf/redis.conf")
	if err != nil {
		logs.Error("redis init conf err")
	}
	key, _ := redisConf.String("prefix")
	host, _ := redisConf.String("host")
	port, _ := redisConf.String("port")
	dbName, _ := redisConf.String("dbName")
	password, _ := redisConf.String("auth")
	config := fmt.Sprintf(`{"key":"%s","conn":"%s:%s","dbName":"%s","password":"%s"}`, key, host, port, dbName, password)
	obj, err := cache.NewCache("redis", config)
	if err != nil {
		logs.Info(err)
	}
	cacheObj = obj
}

// @Title Set
// @Description 缓存设置
// @Success 200
// @router /cache/set [get]
func (t *TestController) Set() {
	err := cacheObj.Put(context.TODO(), "cacheKey", "cacheValue", time.Second*60)
	if err != nil {
		t.Data["json"] = map[string]interface{}{
			"data": err,
			"code": 400,
		}
	}

	t.Data["json"] = "设置成功"
	t.ServeJSON()
}

// @Title Get
// @Description 缓存获取
// @Success 200
// @router /cache/get [get]
func (t *TestController) Get() {
	value, err := cacheObj.Get(context.TODO(), "cacheKey")
	if err != nil {
		t.Data["json"] = map[string]interface{}{
			"data": err,
			"code": 400,
		}
		t.ServeJSON()
	}
	if value == nil {
		t.Data["json"] = map[string]interface{}{
			"data": "缓存失效",
			"code": 400,
		}
		t.ServeJSON()
	}

	t.Data["json"] = string(value.([]byte))
	t.ServeJSON()
}

// @Title Exist
// @Description 缓存是否失效
// @Success 200
// @router /cache/exist [get]
func (t *TestController) IsExist() {
	isExist, err := cacheObj.IsExist(context.TODO(), "cacheKey")
	if err != nil {
		t.Data["json"] = map[string]interface{}{
			"data": err,
			"code": 400,
		}
		t.ServeJSON()
	}
	if isExist {
		t.Data["json"] = map[string]interface{}{
			"data": "缓存未失效",
			"code": 200,
		}
		t.ServeJSON()
	}

	t.Data["json"] = map[string]interface{}{
		"data": "缓存已失效",
		"code": 200,
	}
	t.ServeJSON()
}

// @Title Output
// @Description 日志
// @Success 200
// @router /logs [post]
func (t *TestController) Output() {
	typeName := t.GetString("type", "cmd")
	//日志自定义格式化
	f := &logs.PatternLogFormatter{
		Pattern:    "文件全路径：%F 行数：%n | 时间：%w 消息级别：%T 》》》 %m \r\n",
		WhenFormat: "2006-01-02 15:04:05",
	}
	logs.RegisterFormatter("myformat", f)
	logs.SetGlobalFormatter("myformat")

	log := logs.NewLogger()
	log.SetLogFuncCallDepth(2) //这个层级刚好调用的地方

	//包括：console、file、conn、smtp、es、multifile
	var err error
	switch typeName {
	case "cmd":
		err = log.SetLogger(logs.AdapterConsole, `{"level":7,"color":true,"formatter":"myformat"}`)
	case "file":
		err = log.SetLogger(logs.AdapterFile, `{"formatter":"myformat","filename":"app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":3,"color":true}`)
	case "multifile":
		err = log.SetLogger(logs.AdapterMultiFile, `{"filename":"access.log","separate":["info","debug","error"]}`)
	}
	if err != nil {
		utils.ToJson(t.Controller, err.Error(), err, "500")
		return
	}

	log.Info(typeName)
	log.Warn("this is warning")
	log.Error("this is error")
	log.Debug("this is debug")
	log.Critical("this is crash")
	utils.ToJson(t.Controller, typeName, "logs print", "200")
	return

}
