package controllers

import (
	"BLOG/model"
	"BLOG/services"
	"BLOG/util/result"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// HomeController 首页
type HomeController struct {
	Ctx iris.Context
}

type user struct {
	Name string
}

// Get 首页Get请求
func (ctx *HomeController) Get() mvc.Result {
	results := make(map[string]interface{})
	list, err := services.NewArticleService.GetList(1, 10)

	results["Title"] = "首页"
	if err == nil {
		results["Articles"] = list
	}

	// fmt.Println(ctx.Ctx.GetLocale().Language(), ctx.Ctx.Tr("hi", &user{
	// 	Name: "John Doe",
	// }))

	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}

// GetLogin 登录页面
func (ctx *HomeController) GetLogin() mvc.Result {
	// c.loginOut()
	var (
		results = make(map[string]interface{})
	)

	results["Title"] = "登录"

	return mvc.View{
		Name: "login.html",
		Data: result.Map(results),
	}
}

// PostLogin 登录页面
func (ctx *HomeController) PostLogin() {
	var (
		results  = make(map[string]interface{})
		username = ctx.Ctx.FormValue("username")
		password = ctx.Ctx.FormValue("password")
	)

	user, err := services.NewUserService.Login(username, password)

	fmt.Println(user, err)

	// user, err := c.UsersService.Login(username, password)

	// if err != nil {
	// 	results["success"] = false
	// 	results["message_status"] = err.Error()
	// 	_, _ = ctx.Ctx.JSON(results)
	// 	return
	// }
	// if user.ID <= 0 {
	// 	results["success"] = false
	// 	results["message_status"] = "登录失败"
	// 	_, _ = ctx.Ctx.JSON(results)
	// }

	// c.Sessions.Set(conf.SystemConfig.UserIDKey, int(user.ID))
	// results["success"] = true
	// results["message_status"] = "登录成功"
	// results["return_url"] = "/"
	_, _ = ctx.Ctx.JSON(results)
}

// GetSignup 注册
func (ctx *HomeController) GetSignup() mvc.Result {
	var (
		results = make(map[string]interface{})
	)

	// results["Hots"] = hots
	// results["HotLabels"] = hotLabels
	results["Title"] = "注册"

	return mvc.View{
		Name: "signup.html",
		Data: result.Map(results),
	}
}

// PostSignup 注册
func (ctx *HomeController) PostSignup() {
	results := make(map[string]interface{})
	user := model.User{}
	username := ctx.Ctx.FormValue("username")
	password := ctx.Ctx.FormValue("password")
	email := ctx.Ctx.FormValue("email")

	// status := true
	// message := ""

	// if len(username) == 0 {
	// 	status = false
	// 	message = "用户名不能为空!"
	// }
	// if len(password) == 0 {
	// 	status = false
	// 	message = "密码不能为空!"
	// }
	// if len(email) == 0 {
	// 	status = false
	// 	message = "邮箱不能为空!"
	// }

	// if !status {
	// 	results["success"] = status
	// 	results["message_status"] = message
	// 	_, _ = c.Ctx.JSON(results)
	// 	return
	// }

	// // 生成用户头像
	// avatar.GenerateAvatarFromUsername(username)
	user.Username = username
	user.Password = password
	user.Email = email
	// user.Avatar = "/public/avatar/" + username + ".png"

	err := services.NewUserService.Create(&user)

	fmt.Println(err)
	// if err != nil {
	// 	status = false
	// 	message = err.Error()
	// 	results["success"] = status
	// 	results["message_status"] = message
	// 	_, _ = c.Ctx.JSON(results)
	// 	return
	// }

	// results["success"] = status
	// results["message_status"] = "注册成功"
	// results["return_url"] = "/login"
	_, _ = ctx.Ctx.JSON(results)
	return
}
