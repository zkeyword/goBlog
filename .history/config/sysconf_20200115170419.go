package config

import "time"

// LogPath 日志路径
var LogPath = "runtime/logs/iris_web_%Y%m%d.log"

// ServerPort web服务端口
var ServerPort = 80

// SysTimeform 时间格式化字符串
const SysTimeform string = "2006-01-02 15:04:05"

// SysTimeformShort 日期时间格式化字符串
const SysTimeformShort string = "2006-01-02"

// SysTimeLocation 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

type System struct {
	AppTitle      string
	UserIDKey     string
	UserStructKey string
}

// SystemConfig web系统配置
var SystemConfig = &System{
	AppTitle:      "BlockWallet",
	UserIDKey:     "UserID",
	UserStructKey: "User",
}
