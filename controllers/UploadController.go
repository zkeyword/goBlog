package controllers

import (
	"fmt"
	"io"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// UploadController 首页
type UploadController struct {
	Ctx iris.Context
}

// Get 文章列表Get请求
func (ctx *UploadController) Get() mvc.Result {

	return mvc.View{
		Name: "upload.html",
	}
}

// Post 新增文章
func (ctx *UploadController) Post() {
	// ctx.Ctx.SetMaxRequestBodySize(1024 * 1024 * 10)
	file, info, err := ctx.Ctx.FormFile("file")
	if err != nil {
		fmt.Println(ctx.Ctx.GetHeader("Content-Type"), err)
		return
	}
	defer file.Close()
	filename := info.Filename
	out, err := os.OpenFile("./uploads/"+filename,
		os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()
	io.Copy(out, file)
}
