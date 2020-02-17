package controllers

import (
	"BLOG/model"
	"BLOG/services"
	"BLOG/util/helper"
	"BLOG/util/result"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// ArticleController 首页
type ArticleController struct {
	Ctx iris.Context
}

// Get 文章列表Get请求
func (ctx *ArticleController) Get() mvc.Result {
	page, err := strconv.Atoi(ctx.Ctx.URLParam("page"))
	results := make(map[string]interface{})
	list, err := services.NewArticleService.GetList(page, 10)

	if err == nil {
		results["Title"] = "文章页"
		results["Articles"] = list
		results["paging"] = helper.CreatePaging(int64(list.Page), int64(list.PageSize), int64(list.Total))
	}

	return mvc.View{
		Name: "article.html",
		Data: result.Map(results),
	}
}

// Post 新增文章
func (ctx *ArticleController) Post() {
	title := ctx.Ctx.PostValue("title")
	content := ctx.Ctx.PostValue("content")
	tagID, _ := strconv.Atoi(ctx.Ctx.PostValue("tagId"))

	Article := &model.Article{
		Content:  content,
		Title:    title,
		AuthorID: 1,
	}

	ArticleID, _ := services.NewArticleService.Create(Article)

	if tagID != 0 {
		ArticleTag := &model.ArticleTag{
			TagID:     uint(tagID),
			ArticleID: ArticleID,
		}
		services.NewArticleTagService.Create(ArticleTag)
	}
}

// GetBy 文章详情 /article/123
func (ctx *ArticleController) GetBy(articleID int64) mvc.Result {
	var results = make(map[string]interface{})
	var article = services.NewArticleService.Get(articleID)

	if article != nil {
		author, _ := services.NewUserService.FindByID(article.AuthorID)
		results["Title"] = article.Title
		results["Article"] = article
		results["Prev"] = services.NewArticleService.GetPrev(articleID)
		results["Next"] = services.NewArticleService.GetNext(articleID)
		results["Author"] = author
	}

	return mvc.View{
		Name: "articleDetail.html",
		Data: result.Map(results),
	}
}
