package app

import (
	"BLOG/controllers"
	"BLOG/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// InnerRoute 注入路由
func InnerRoute(app *Application) {

	// 设置静态路径
	app.HandleDir("/public", "./public")

	// blog
	mvc.Configure(app.Party("/"), func(m *mvc.Application) {
		m.Register(app.Sessions.Start) // 保持在controller之前，否则无效
		m.Party("/").Handle(new(controllers.HomeController))
		m.Party("/article").Handle(new(controllers.ArticleController))
		// m.Party("/tag").Handle(new(controllers.HomeController))
		// m.Party("/user").Handle(new(controllers.HomeController))
		// m.Party("/category").Handle(new(controllers.HomeController))
	})

	// api
	mvc.Configure(app.Party("/api/", middleware.CORS), func(m *mvc.Application) {
		// m.Router.Use(middleware.AdminAuth)
		m.Party("/upload", iris.LimitRequestBodySize(5<<20)).Handle(new(controllers.UploadController))
		m.Party("/captcha").Register(app.Sessions.Start).Handle(new(controllers.CaptchaController))
		m.Party("/login").Handle(new(controllers.LoginController))
		m.Party("/user", middleware.JWT.Serve).Handle(new(controllers.UserController))
	})

}
