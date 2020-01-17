package config

// DriverName 数据库名
const DriverName = "mysql"

// DbConf 数据库结构
type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

// DbConfig 主库配置
var DbConfig DbConf = DbConf{
	Host:   "localhost",
	Port:   3306,
	User:   "root",
	Pwd:    "123456zjx",
	DbName: "aggra_forum",
}
