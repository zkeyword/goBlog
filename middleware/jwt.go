package middleware

import (
	"BLOG/model"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"time"
)

var (
	// JWT JWT Middleware
	JWT *jwt.Middleware
)

func initJWT() {
	JWT = jwt.New(jwt.Config{
		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			}

			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(model.ResModel{
				Code: "501",
				Msg:  err.Error(),
			})
		},
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("7a0CyHC3O5v0KpeHH4asKFpf80wLru2I"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

// GetJWTString get jwt string with expiration time 20 minutes
func GetJWTString(name string, id uint) (string, error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 根据需求，可以存一些必要的数据
		"userName": name,
		"userId":   id,

		// 签发时间
		"iat": time.Now().Unix(),
		// 设定过期时间，设置20分钟过期
		"exp": time.Now().Add(20 * time.Minute * time.Duration(1)).Unix(),
	})

	// 使用设置的秘钥，签名生成jwt字符串
	tokenString, err := token.SignedString([]byte("7a0CyHC3O5v0KpeHH4asKFpf80wLru2I"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
