package controllers

import (
	"BLOG/model"
	"BLOG/services"
	"BLOG/util/result"
	"fmt"
	"unicode/utf8"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// ArticleController 首页
type ArticleController struct {
	Ctx iris.Context
}

// Get 文章列表Get请求
func (ctx *ArticleController) Get() mvc.Result {
	var results = make(map[string]interface{})

	results["Title"] = "文章页"
	results["Articles"] = services.NewArticleService.GetList()

	return mvc.View{
		Name: "article.html",
		Data: result.Map(results),
	}
}

// Post 新增文章
func (ctx *ArticleController) Post() {
	title := ctx.Ctx.PostValue("title")
	content := ctx.Ctx.PostValue("content")

	fmt.Println(utf8.ValidString(title))

	var Article = &model.Article{
		Content:  content,
		Title:    title,
		AuthorID: 1,
	}
	services.NewArticleService.Create(Article)
}

// GetBy 文章详情 /article/123
func (ctx *ArticleController) GetBy(articleID int64) mvc.Result {
	var results = make(map[string]interface{})
	var article = services.NewArticleService.Get(articleID)

	if article != nil {
		results["Title"] = article.Title
		results["Article"] = article
		results["Prev"] = services.NewArticleService.GetPrev(articleID)
		results["Next"] = services.NewArticleService.GetNext(articleID)
	}

	return mvc.View{
		Name: "articleDetail.html",
		Data: result.Map(results),
	}
}
