package controllers

import (
	"fmt"
	"io"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// UploadController 首页
type UploadController struct {
	Ctx iris.Context
}

// Get 上传页面 GET /upload
func (ctx *UploadController) Get() mvc.Result {

	return mvc.View{
		Name: "upload.html",
	}
}

// Post 上传 POST /upload
func (ctx *UploadController) Post() {
	// ctx.Ctx.SetMaxRequestBodySize(1024 * 1024 * 10)
	file, info, err := ctx.Ctx.FormFile("file")
	if err != nil {
		fmt.Println(ctx.Ctx.GetHeader("Content-Type"), err)
		return
	}
	defer file.Close()
	filename := info.Filename
	out, err := os.OpenFile("./public/uploads/"+filename,
		os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()
	io.Copy(out, file)
}
