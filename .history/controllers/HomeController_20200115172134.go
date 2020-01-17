package controllers

import (
	"github.com/kataras/iris"
)

// HomeController 首页
func HomeController(ctx iris.Context) {
	userList := []string{
		"Alice",
		"Bob",
		"Tom",
	}

	userList2 := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	ctx.ViewData("userList", userList)
	ctx.ViewData("userList2", userList2)
	ctx.ViewData("message", "xxx1111")
	ctx.View("home.html")
}
