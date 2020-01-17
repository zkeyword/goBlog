package helper

import (
	"context"
	"strconv"
	"sync"
	"time"

	"IRIS_WEB/route"
	"github.com/kataras/iris"
)

// RunIris 启动iris
func RunIris(port int) {
	app := iris.New()

	app.Use(NewRecoverMdw())
	app.Use(NewAccessLogMdw())

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

	// 注册路由
	route.InnerRoute(app)

	// server配置
	c := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           true,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: true,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
		IgnoreServerErrors:                []string{iris.ErrServerClosed.Error()},
		RemoteAddrHeaders:                 map[string]bool{"X-Real-Ip": true, "X-Forwarded-For": true},
	})

	app.Run(iris.Addr(":"+strconv.Itoa(port)), c)
}
