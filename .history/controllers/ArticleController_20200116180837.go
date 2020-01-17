package controllers

import (
	"BLOG/model"
	"BLOG/services"
	"BLOG/util/result"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// ArticleController 首页
type ArticleController struct {
	Ctx iris.Context
}

// Get 文章列表Get请求
func (ctx *ArticleController) Get() mvc.Result {
	var results = make(map[string]interface{})

	results["Title"] = "文章页"
	results["Articles"] = services.ArticleService.GetList()

	return mvc.View{
		Name: "article.html",
		Data: result.Map(results),
	}
}

// Post 新增文章
func (ctx *ArticleController) Post() {
	title := ctx.Ctx.PostValue("title")
	content := ctx.Ctx.PostValue("content")

	var Article = &model.Article{
		Content:  content,
		Title:    title,
		AuthorID: 1,
	}
	services.ArticleService.Create(Article)
}
