package app

import (
	"html/template"
	"time"

	"BLOG/config"
	"BLOG/controllers"
	"BLOG/middleware"
	"BLOG/util/strtime"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

// Application app
type Application struct {
	*iris.Application
	Sessions *sessions.Sessions
}

// InnerRoute 注入路由
func InnerRoute(app *Application) {

	SetupViews(app, "./views")
	SetupErrorHandlers(app)
	SetupSessions(app)

	// 设置静态路径
	app.HandleDir("/public", "./public")

	// blog
	mvc.Configure(app.Party("/", middleware.Cors()), func(m *mvc.Application) {
		m.Register(app.Sessions.Start) // 保持在controller之前，否则无效
		m.Party("/").Handle(new(controllers.HomeController))
		m.Party("/article").Handle(new(controllers.ArticleController))
		m.Party("/upload", iris.LimitRequestBodySize(5<<20)).Handle(new(controllers.UploadController))
		m.Party("/captcha").Handle(new(controllers.CaptchaController))
		// m.Party("/tag").Handle(new(controllers.HomeController))
		// m.Party("/user").Handle(new(controllers.HomeController))
		// m.Party("/category").Handle(new(controllers.HomeController))
	})

	// api
	mvc.Configure(app.Party("/api/"), func(m *mvc.Application) {
		// m.Router.Use(middleware.AdminAuth)
		// m.Party("/user").Handle(new(admin.UserController))
	})

}

// SetupViews 设置Views
func SetupViews(app *Application, viewsDir string) {
	htmlEngine := iris.HTML(viewsDir, ".html").Layout("shared/layout.html")
	htmlEngine.Reload(true)
	// 给模板内置方法
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int64) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(config.SysTimeformShort)
	})

	htmlEngine.AddFunc("FromTimeString", func(t time.Time) string {
		dt := time.Unix(t.Unix(), int64(0))
		return dt.Format(config.SysTimeform)
	})

	htmlEngine.AddFunc("FromUnixtime", func(t int64) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(config.SysTimeform)
	})

	htmlEngine.AddFunc("FromStrTime", func(t time.Time) string {
		return strtime.StrTime(t.Unix())
	})

	htmlEngine.AddFunc("Html", func(t string) template.HTML {
		return template.HTML(t)
	})

	app.RegisterView(htmlEngine)
}

// SetupErrorHandlers 错误处理
func SetupErrorHandlers(app *Application) {
	app.OnAnyErrorCode(func(ctx iris.Context) {
		errorMsg := ctx.Values().GetString("message_status")
		errorCode := ctx.GetStatusCode()
		if len(errorMsg) == 0 {
			if errorCode == 404 {
				errorMsg = "你已闯入非法领域!"
			}
		}
		err := iris.Map{
			"status":         errorCode,
			"message_status": errorMsg,
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			_, _ = ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.ViewLayout("shared/layout_error.html")
		_ = ctx.View("shared/error.html")
	})
}

//SetupSessions 设置Session
func SetupSessions(app *Application) {
	app.Sessions = sessions.New(sessions.Config{
		Cookie:  "ssid",
		Expires: 24 * time.Hour,
	})
}
