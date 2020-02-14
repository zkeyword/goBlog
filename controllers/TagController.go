package controllers

import (
	"BLOG/services"
	"BLOG/util/result"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// TagController 首页
type TagController struct {
	Ctx iris.Context
}

// Get 文章列表Get请求
func (ctx *TagController) Get() mvc.Result {
	var results = make(map[string]interface{})

	results["Title"] = "文章页"
	results["Articles"] = services.NewTagService.GetList()

	return mvc.View{
		Name: "article.html",
		Data: result.Map(results),
	}
}
