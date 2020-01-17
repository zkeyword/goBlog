package controllers

import (
	"BLOG/util/result"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// HomeController 首页
func HomeController(ctx iris.Context) {
	// func HomeController(ctx iris.Context) mvc.Result {
	// var results = make(map[string]interface{})
	// results["Title"] = "首页"
	// return mvc.View{
	// 	Name: "index.html",
	// 	Data: result.Map(results),
	// }

	var results = make(map[string]interface{})
	results["Title"] = "首页"
	result.Map(results)
	mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}

	fmt.Println(mvc.View)

	ctx.ViewData("message", "xxx1111")
	ctx.View("index.html")
}

// func HomeController(ctx iris.Context) mvc.Result {
// 	var results = make(map[string]interface{})
// 	results["Title"] = "首页"
// 	return mvc.View{
// 		Name: "index.html",
// 		Data: result.Map(results),
// 	}
// }
