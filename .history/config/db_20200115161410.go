package config

// DriverName 数据库名
const DriverName = "mysql"

// DbConf 数据库配置
type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

// MasterDbConfig 主库配置
var MasterDbConfig DbConf = DbConf{
	Host:   "localhost",
	Port:   3306,
	User:   "root",
	Pwd:    "123456zjx",
	DbName: "aggra_forum",
}

// SlaveDbConfig 从库配置
var SlaveDbConfig DbConf = DbConf{
	Host:   "localhost",
	Port:   3306,
	User:   "root",
	Pwd:    "123456zjx",
	DbName: "aggra_forum",
}
