package db

import (
	"BLOG/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var mysqlDB *gorm.DB

// StartMysql 初始化mysql
func StartMysql(dsn string, maxIdle, maxOpen int) (err error) {
	mysqlDB, err = gorm.Open("mysql", dsn)

	if err == nil {
		mysqlDB.DB().SetMaxIdleConns(maxIdle)
		mysqlDB.DB().SetMaxOpenConns(maxOpen)
		mysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	}

	return
}

// GetMysql 获取mysql连接
func GetMysql() *gorm.DB {
	mysqlDB.Set("gorm:table_options", "ENGINE=utf8mb4_general_ci")
	mysqlDB.AutoMigrate(&model.Article{}, &model.User{})
	return mysqlDB
}

// CloseMysql 关闭mysql
func CloseMysql() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}
