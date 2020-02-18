package controllers

import (
	"BLOG/model"
	"BLOG/services"
	"BLOG/util/result"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// TagController 首页
type TagController struct {
	Ctx iris.Context
}

// Get 标签列表Get请求
func (ctx *TagController) Get() mvc.Result {
	var results = make(map[string]interface{})

	results["Title"] = "文章页"
	results["Tags"] = services.NewTagService.GetList()

	return mvc.View{
		Name: "tag.html",
		Data: result.Map(results),
	}
}

// Post 标签列表新增
func (ctx *TagController) Post() {
	title := ctx.Ctx.PostValue("title")

	Tag := &model.Tag{
		Title: title,
	}

	services.NewTagService.Create(Tag)
}
