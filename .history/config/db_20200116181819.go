package config

// DriverName 数据库名
const DriverName = "mysql"

// DbConf 数据库结构
type DbConf struct {
	Dsn     string
	MaxIdle int
	MaxOpen int
}

// DbConfig 主库配置
var DbConfig DbConf = DbConf{
	Dsn:     "root:123456zjx@tcp(127.0.0.1:3306)/aggra_forum?charset=utf8mb4&parseTime=true&loc=Local",
	MaxIdle: 10,
	MaxOpen: 100,
}
