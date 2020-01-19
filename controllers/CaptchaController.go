package controllers

import (
	// "BLOG/util/captcha"

	"github.com/dchest/captcha"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
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
	ctx.Session.Increment("captchaID", captchaID)
	captcha.WriteImage(ctx.Ctx.ResponseWriter(), captchaID, 100, 40)

}

// verify 校验验证码
func (ctx *CaptchaController) verify(code string) {
	// isOK := captcha.VerifyString(session, code)

}
