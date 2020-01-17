package main

import (
	"flag"
	"fmt"
	"blog/config"
	"blog/util/db"
)

func main() {
	flag.Parse()
	fmt.Print("InitConfig...\r")
	
	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("StartMysql...\r")
	checkErr("StartMysql", db.StartMysql(Conf.MysqlDsn, Conf.MysqlMaxIdle, Conf.MysqlMaxOpen))
	fmt.Print("StartMysql Success!!!\n")
	config.MasterDbConfig.Host
}
