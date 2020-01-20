package controllers

import (
	"BLOG/model"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// UserController 首页
type UserController struct {
	Ctx iris.Context
}

// Post /user/
func (ctx *UserController) Post() {
	token := ctx.Ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	ctx.Ctx.JSON(new(model.ResModel).WithData(token))
}
