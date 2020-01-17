package model

import (
	"github.com/jinzhu/gorm"
)

// Article 文章表
type Article struct {
	gorm.Model
	Title   string `gorm:"varchar(255)"`
	Content string `gorm:"LONGTEXT"`
}
