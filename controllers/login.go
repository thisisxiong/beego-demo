package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"time"
	"user/models"
)

type LoginController struct {
	beego.Controller
}



// @Title Register
// @Description 用户注册
// @Success 200
// @router  /register  [post]
func (l *LoginController) Register() {
	valid := validation.Validation{}
    user  := models.User{}
    user.Username = l.GetString("username")
    user.Password = l.GetString("password")
    b,_ := valid.Valid(&user)
    if !b {
		l.Data["json"] = map[string]interface{}{
			"code":    400,
			"data":    valid.Errors,
			"message": valid.Errors,
		}
		l.ServeJSON()
		return
	}


	has := md5.Sum([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", has)
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		l.Data["json"] = map[string]interface{}{
			"code":    417,
			"data":    err,
			"message": err,
		}
		l.ServeJSON()
		return
	}

	l.Data["json"] = map[string]interface{}{
		"code":    200,
		"data":    user,
		"message": "注册成功",
	}
	l.ServeJSON()
	return

}

// @Title Login
// @Description 用户登录
// @Success 200
// @router /login [post]
func (l *LoginController) Login() {
	username := l.GetString("username")
	password := l.GetString("password")
	if username == "" || password == "" {
		l.Data["json"] = map[string]interface{}{
			"code":    400,
			"data":    "用户名密码不能为空",
			"message": "用户名密码不能为空",
		}
		l.ServeJSON()
		return
	}
	var user models.User
	user.Username = username
	has := md5.Sum([]byte(password))
	user.Password = fmt.Sprintf("%x", has)

	o := orm.NewOrm()
	err := o.QueryTable(&user).Filter("username", user.Username).Filter("password", user.Password).One(&user)
	if err != nil {
		l.Data["json"] = map[string]interface{}{
			"code":    400,
			"data":    err,
			"message": "用户名密码错误",
		}
		l.ServeJSON()
		return
	}

	secret, _ := config.String("secret")
	token, err := CreateToken(string(user.Id), secret)
	if err != nil {
		l.Data["json"] = map[string]interface{}{
			"code":    500,
			"data":    "系统错误",
			"message": "系统错误",
		}
		l.ServeJSON()
		return
	}

	l.Data["json"] = map[string]interface{}{
		"code":    200,
		"data":    token,
		"message": "登录成功",
	}
	l.ServeJSON()
	return

}

//生成token
func CreateToken(uid, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
