package utils

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

/**
格式化时间
*/
func GetDay() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

/**
纳秒
*/
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

/**
当前时间戳
*/
func GetUnix() int64 {
	return time.Now().Unix()
}

/**
json返回
*/
func ToJson(controller beego.Controller, data interface{}, msg interface{}, code interface{}) {
	controller.Data["json"] = map[string]interface{}{
		"code":      code,
		"data":      data,
		"message":   msg,
		"timestamp": time.Now().Unix(),
	}
	controller.ServeJSON()
}
