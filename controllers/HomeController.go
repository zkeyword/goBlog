package controllers

import (
	"BLOG/model"
	"BLOG/services"
	"BLOG/util/db"
	"BLOG/util/result"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"encoding/csv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gomodule/redigo/redis"
)

// HomeController 首页
type HomeController struct {
	Ctx iris.Context
}

type user struct {
	Name string
}

// Get 首页Get请求
func (ctx *HomeController) Get() mvc.Result {
	results := make(map[string]interface{})
	list, err := services.NewArticleService.GetList(1, 10)

	results["Title"] = "首页"
	if err == nil {
		results["Articles"] = list
	}

	c := db.GetRedis()

	defer c.Close()

	length, err := c.Do("LLEN", "h:l:t")
	if err != nil {
		fmt.Println("err", err)
	}

	if length.(int64) < 10 {
		//写入数据
		_, err := c.Do("LPUSH", "h:l:t", "1")
		if err != nil {
			fmt.Println("err while setting:", err)
		}
	} else {
		//读取数据
		val, err := redis.Values(c.Do("LRANGE", "h:l:t", 0, 10))
		if err != nil {
			fmt.Println("err while setting:", err)
		}
		for _, v := range val {
			fmt.Printf("%s ", v.([]uint8))
		}
	}

	// //判断key是否存在
	// isKeyExit, err := redis.Bool(c.Do("EXISTS", "go_key"))
	// if err != nil {
	// 	fmt.Println("err while checking keys:", err)
	// } else {
	// 	fmt.Println(isKeyExit)
	// }

	// fmt.Println(ctx.Ctx.GetLocale().Language(), ctx.Ctx.Tr("hi", &user{
	// 	Name: "John Doe",
	// }))

	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}

// GetLogin 登录页面
func (ctx *HomeController) GetLogin() mvc.Result {
	// c.loginOut()
	var (
		results = make(map[string]interface{})
	)

	results["Title"] = "登录"

	return mvc.View{
		Name: "login.html",
		Data: result.Map(results),
	}
}

// PostLogin 登录页面
func (ctx *HomeController) PostLogin() {
	var (
		results  = make(map[string]interface{})
		username = ctx.Ctx.FormValue("username")
		password = ctx.Ctx.FormValue("password")
	)

	user, err := services.NewUserService.Login(username, password)

	fmt.Println(user, err)

	// user, err := c.UsersService.Login(username, password)

	// if err != nil {
	// 	results["success"] = false
	// 	results["message_status"] = err.Error()
	// 	_, _ = ctx.Ctx.JSON(results)
	// 	return
	// }
	// if user.ID <= 0 {
	// 	results["success"] = false
	// 	results["message_status"] = "登录失败"
	// 	_, _ = ctx.Ctx.JSON(results)
	// }

	// c.Sessions.Set(conf.SystemConfig.UserIDKey, int(user.ID))
	// results["success"] = true
	// results["message_status"] = "登录成功"
	// results["return_url"] = "/"
	_, _ = ctx.Ctx.JSON(results)
}

// GetSignup 注册
func (ctx *HomeController) GetSignup() mvc.Result {
	var (
		results = make(map[string]interface{})
	)

	// results["Hots"] = hots
	// results["HotLabels"] = hotLabels
	results["Title"] = "注册"

	return mvc.View{
		Name: "signup.html",
		Data: result.Map(results),
	}
}

// PostSignup 注册
func (ctx *HomeController) PostSignup() {
	results := make(map[string]interface{})
	user := model.User{}
	username := ctx.Ctx.FormValue("username")
	password := ctx.Ctx.FormValue("password")
	email := ctx.Ctx.FormValue("email")

	// status := true
	// message := ""

	// if len(username) == 0 {
	// 	status = false
	// 	message = "用户名不能为空!"
	// }
	// if len(password) == 0 {
	// 	status = false
	// 	message = "密码不能为空!"
	// }
	// if len(email) == 0 {
	// 	status = false
	// 	message = "邮箱不能为空!"
	// }

	// if !status {
	// 	results["success"] = status
	// 	results["message_status"] = message
	// 	_, _ = c.Ctx.JSON(results)
	// 	return
	// }

	// // 生成用户头像
	// avatar.GenerateAvatarFromUsername(username)
	user.Username = username
	user.Password = password
	user.Email = email
	// user.Avatar = "/public/avatar/" + username + ".png"

	err := services.NewUserService.Create(&user)

	fmt.Println(err)
	// if err != nil {
	// 	status = false
	// 	message = err.Error()
	// 	results["success"] = status
	// 	results["message_status"] = message
	// 	_, _ = c.Ctx.JSON(results)
	// 	return
	// }

	// results["success"] = status
	// results["message_status"] = "注册成功"
	// results["return_url"] = "/login"
	_, _ = ctx.Ctx.JSON(results)
	return
}

// GetCsv test csv
func (ctx *HomeController) GetCsv() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"`
	r := csv.NewReader(strings.NewReader(in))

	ctx.Ctx.ContentType("text/csv")

	record, err := r.Read()
	i := 0

	ctx.Ctx.StreamWriter(func(w io.Writer) bool {
		// time.Sleep(500 * time.Millisecond) // simulate delay.
		if i == len(record)-1 {
			return false //关闭并刷新
		}
		i++
		if err == io.EOF {
			return false
		}
		if err != nil {
			log.Fatal(err)
		}
		return true //继续写入数据
	})
}

// GetXlsx 导出xslx文件
func (ctx *HomeController) GetXlsx() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
	// ctx.Ctx.Write(f.WriteToBuffer())
	buffer, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
	}
	filename := "results.xlsx"

	http.ServeContent(ctx.Ctx.ResponseWriter(), ctx.Ctx.Request(), filename, time.Now(), strings.NewReader(buffer.String()))
}

// GetRedis redis操作
func (ctx *HomeController) GetRedis() {
	c := db.GetRedis()

	//记得销毁本次链连接
	defer c.Close()

	//写入数据
	_, err := c.Do("SET", "go_key", "redigo")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	//判断key是否存在
	isKeyExit, err := redis.Bool(c.Do("EXISTS", "go_key"))
	if err != nil {
		fmt.Println("err while checking keys:", err)
	} else {
		fmt.Println(isKeyExit)
	}

	//获取value并转成字符串
	accountBalance, err := redis.String(c.Do("GET", "go_key"))
	if err != nil {
		fmt.Println("err while getting:", err)
	} else {
		fmt.Println(accountBalance)
	}

	//删除key
	_, err = c.Do("DEL", "go_key")
	if err != nil {
		fmt.Println("err while deleting:", err)
	}

	//设置key过期时间
	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("err while setting:", err)
	}

	//对已有key设置5s过期时间
	n, err := c.Do("EXPIRE", "go_key", 5)
	if err != nil {
		fmt.Println("err while expiring:", err)
	} else if n != int64(1) {
		fmt.Println("failed")
	}
	return
}
