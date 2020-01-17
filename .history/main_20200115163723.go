package main

import (
	"blog/config"
	"blog/util/db"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	fmt.Print("InitConfig...\r")

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("StartMysql...\r")
	checkErr("StartMysql", db.StartMysql(config.DbConfig.Dsn, config.DbConfig.MaxIdle, config.DbConfig.MaxOpen))
	fmt.Print("StartMysql Success!!!\n")
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s Error: %v\n", errMsg, err)
		os.Exit(1)
	}
}
