package route

import (
	"BLOG/controllers"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

// 定义500错误处理函数
func err500(ctx iris.Context) {
	ctx.WriteString("CUSTOM 500 ERROR")
}

// 定义404错误处理函数
func err404(ctx iris.Context) {
	ctx.WriteString("CUSTOM 404 ERROR")
}

// InnerRoute 注入路由
func InnerRoute(app *iris.Application) {
	app.StaticWeb("/static", "./assets")
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))
	app.OnErrorCode(iris.StatusInternalServerError, err500)
	app.OnErrorCode(iris.StatusNotFound, err404)
	app.Get("/", controllers.HomeController)
}
