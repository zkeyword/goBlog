package controllers

import (
	// "BLOG/util/captcha"

	"github.com/dchest/captcha"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	"BLOG/model"
)

// CaptchaController 验证码控制器
type CaptchaController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

// Get 获取验证码
func (ctx *CaptchaController) Get() {
	ctx.Ctx.Header("Content-Type", "image/png")
	captchaID := captcha.NewLen(4)
	ctx.Session.Set("captchaID", captchaID)
	captcha.WriteImage(ctx.Ctx.ResponseWriter(), captchaID, 100, 40)
}

// Post 校验验证码
func (ctx *CaptchaController) Post() {
	code := ctx.Ctx.PostValue("code")
	isOK := captcha.VerifyString(ctx.Session.GetString("captchaID"), code)
	if isOK {
		ctx.Ctx.JSON(new(model.ResModel).WithData(isOK))
		return
	}
	ctx.Ctx.JSON(new(model.ResModel).WithError("-1", "验证码错误"))

}
