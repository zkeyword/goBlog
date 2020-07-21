package main

import (
	"BLOG/app"
	"BLOG/config"
	"BLOG/util/db"
	"flag"
	"fmt"
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	// 创建文件日志，按天分割，日志文件仅保留一周
	w, err := rotatelogs.New(config.LogPath)
	checkErr("CreateRotateLog", err)

	// 设置日志
	logrus.SetOutput(w)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("Start Mysql...\r")
	checkErr("Start Mysql", db.StartMysql(config.DbConfig.Dsn, config.DbConfig.MaxIdle, config.DbConfig.MaxOpen))
	fmt.Print("Start Mysql Success!!!\n")

	// 启动redis
	defer db.CloseRedis()
	checkErr("Start Redis", db.StartRedis(config.RedisDbConfig.Addr, config.RedisDbConfig.Password, config.RedisDbConfig.DB, config.RedisDbConfig.MaxIdle, config.RedisDbConfig.MaxOpen))
	fmt.Print("Start Redis Success!!!\n")

	// 开始运行iris框架
	app.New()
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s Error: %v\n", errMsg, err)
		os.Exit(1)
	}
}
