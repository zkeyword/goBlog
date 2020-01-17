package controllers

import (
	"BLOG/util/result"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// ArticleController 首页
type ArticleController struct {
	Ctx iris.Context
}

// Get 首页Get请求
func (ctx *HomeController) Get() mvc.Result {
	var results = make(map[string]interface{})
	results["Title"] = "首页"
	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}

func (ctx *HomeController) GetBy(id string) mvc.Result {
	var results = make(map[string]interface{})
	results["Title"] = "首页22"
	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}
