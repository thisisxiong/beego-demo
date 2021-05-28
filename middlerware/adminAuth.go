package middlerware

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
)

//登录校验
func AdminAuth(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	secret, _ := config.String("secret")
	_, err := ParseToken(token, secret)
	if err != nil {
		data := map[string]interface{}{
			"code":    403,
			"data":    "请先登录",
			"message": "请先登录",
		}
		str, _ := json.Marshal(data)
		ctx.Output.Body(str)
	}

}

func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
