package route

import (
	"BLOG/controllers"

	"github.com/kataras/iris"
	"time"
	"BLOG/config"
	"BLOG/util/strtime"
)

// InnerRoute 注入路由
func InnerRoute(app *iris.Application) {
	app.StaticWeb("/static", "./assets")
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))
	app.OnErrorCode(iris.StatusInternalServerError, err500)
	app.OnErrorCode(iris.StatusNotFound, err404)
	app.Get("/", controllers.HomeController)
}

func (b *Bootstrapper) SetupViews(viewsDir string) {
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
	b.RegisterView(htmlEngine)
}

// 定义500错误处理函数
func err500(ctx iris.Context) {
	ctx.WriteString("CUSTOM 500 ERROR")
}

// 定义404错误处理函数
func err404(ctx iris.Context) {
	ctx.WriteString("CUSTOM 404 ERROR")
}

