package controllers

import (
	"BLOG/middleware"
	"BLOG/model"

	"github.com/kataras/iris/v12"
)

// LoginController 首页
type LoginController struct {
	Ctx iris.Context
}

// Post /login/ 登录
func (ctx *LoginController) Post() {
	token, err := middleware.GetJWTString("111", 1)
	if err != nil {
		ctx.Ctx.JSON(new(model.ResModel).WithError("-1", err.Error()))
	}
	ctx.Ctx.JSON(new(model.ResModel).WithData(token))
}
