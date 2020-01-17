package controllers

import (
	"BLOG/util/result"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// HomeController 首页控制器
type HomeController struct {
	Ctx iris.Context
}

// HomeController 首页
// func HomeController(ctx iris.Context) {
// 	// func HomeController(ctx iris.Context) mvc.Result {
// 	// var results = make(map[string]interface{})
// 	// results["Title"] = "首页"
// 	// return mvc.View{
// 	// 	Name: "index.html",
// 	// 	Data: result.Map(results),
// 	// }

// 	var results = make(map[string]interface{})
// 	results["Title"] = "首页"
// 	result.Map(results)
// 	var a = mvc.View{
// 		Name: "index.html",
// 		Data: result.Map(results),
// 	}

// 	fmt.Println(a)

// 	ctx.ViewData("message", "xxx1111")
// 	ctx.View("index.html")
// }
func (this *HomeController) Get() mvc.Result {
	var results = make(map[string]interface{})
	results["Title"] = "首页"
	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}
