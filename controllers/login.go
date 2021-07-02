package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"time"
	"user/models"
	"user/utils"
)

type LoginController struct {
	beego.Controller
}

// @Title Register
// @Description 用户注册
// @Param username formData string true 用户名
// @Param password formData string true 密码
// @Success 200
// @router /register [post]
func (l *LoginController) Register() {
	valid := validation.Validation{}
	user := models.User{}
	user.Username = l.GetString("username")
	user.Password = l.GetString("password")
	b, _ := valid.Valid(&user)
	if !b {
		utils.ToJson(l.Controller, valid.Errors, "验证失败", 417)
		return
	}

	has := md5.Sum([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", has)
	o := models.Open(&user)
	_, err := o.Insert(&user)
	if err != nil {
		utils.ToJson(l.Controller, err, "系统错误", 400)
		return
	}

	utils.ToJson(l.Controller, user, "注册成功", 200)
	return

}

// @Title Login
// @Description 用户登录
// @Param username formData string true 用户名
// @Param password formData string true 密码
// @Success 200
// @router /login [post]
func (l *LoginController) Login() {
	username := l.GetString("username")
	password := l.GetString("password")
	if username == "" || password == "" {
		utils.ToJson(l.Controller, "", "用户名或密码错误", 417)
		return
	}
	var user models.User

	user.Username = username
	has := md5.Sum([]byte(password))
	user.Password = fmt.Sprintf("%x", has)

	o := models.Open(&user)
	err := o.QueryTable(&user).Filter("username", user.Username).Filter("password", user.Password).One(&user)
	if err != nil {
		utils.ToJson(l.Controller, err, "系统错误", 500)
		return
	}

	secret, _ := config.String("secret")
	token, err := CreateToken(string(user.Id), secret)
	if err != nil {
		utils.ToJson(l.Controller, err, "创建token失败", 500)
		return
	}

	utils.ToJson(l.Controller, token, "登录成功", 200)
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
