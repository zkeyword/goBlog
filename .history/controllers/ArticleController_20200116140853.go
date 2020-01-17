package controllers

import (
	"BLOG/services"
	"BLOG/util/result"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// ArticleController 首页
type ArticleController struct {
	Ctx iris.Context
}

// Get 首页Get请求
func (ctx *ArticleController) Get() mvc.Result {
	var results = make(map[string]interface{})
	results["Title"] = "文章页"
	fmt.Println("1", services.UsersService.CheckUserIsLockByID("1"))
	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}

func (ctx *ArticleController) GetBy(id string) mvc.Result {
	var results = make(map[string]interface{})
	results["Title"] = "首页22"
	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}
