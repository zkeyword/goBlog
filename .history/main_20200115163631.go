package main

import (
	"blog/config"
	"blog/util/db"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Print("InitConfig...\r")

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("StartMysql...\r")
	checkErr("StartMysql", db.StartMysql(config.DbConfig.Dsn, config.DbConfig.MaxIdle, config.DbConfig.MaxOpen))
	fmt.Print("StartMysql Success!!!\n")
	config.MasterDbConfig.Host
}
