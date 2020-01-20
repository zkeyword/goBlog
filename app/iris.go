package app

import (
	"context"
	"html/template"
	"strconv"
	"sync"
	"time"

	"BLOG/config"
	"BLOG/middleware"
	"BLOG/util/strtime"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	i18n "github.com/iris-contrib/middleware/go-i18n"
)

// Application app
type Application struct {
	*iris.Application
	Sessions *sessions.Sessions
}

// New 初始化
func New() {
	b := &Application{
		Application: iris.New(),
	}
	b.RunIris()
}

// RunIris 启动iris
func (app *Application) RunIris() *Application {

	// 错误拦截以及访问日志
	app.Use(middleware.Recover())
	app.Use(middleware.AccessLog())

	// 优雅的关闭程序
	serverWG := new(sync.WaitGroup)
	defer serverWG.Wait()

	iris.RegisterOnInterrupt(func() {
		serverWG.Add(1)
		defer serverWG.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		app.Shutdown(ctx)
	})

	// 国际化
	app.I18n.Reset(i18n.NewLoader("./locales/*.yaml"))
	app.I18n.SetDefault("en-US")

	// 设置view路径
	SetupViews(app, "./views")

	// 设置错误处理
	SetupErrorHandlers(app)

	// 设置错误处理
	SetupSessions(app)

	// 注册路由
	InnerRoute(app)

	// server配置
	c := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           true,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: true,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        config.SysTimeform,
		Charset:                           "UTF-8",
		IgnoreServerErrors:                []string{iris.ErrServerClosed.Error()},
		RemoteAddrHeaders:                 map[string]bool{"X-Real-Ip": true, "X-Forwarded-For": true},
	})

	app.Run(iris.Addr(":"+strconv.Itoa(config.ServerPort)), c)

	return app
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
