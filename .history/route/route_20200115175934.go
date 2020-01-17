package route

import (
	"BLOG/controllers"

	"BLOG/config"
	"BLOG/util/strtime"
	"github.com/kataras/iris"
	"time"
)

// InnerRoute 注入路由
func InnerRoute(app *iris.Application) {

	SetupViews(app, "./views")
	SetupErrorHandlers(app)

	// 设置静态路径
	app.StaticWeb("/public", "./public")

	app.Get("/", controllers.HomeController)
}

// SetupViews 设置Views
func SetupViews(app *iris.Application, viewsDir string) {
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

	app.RegisterView(htmlEngine)
}

// SetupErrorHandlers 错误处理
func SetupErrorHandlers(app *iris.Application) {
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
