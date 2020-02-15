package model

import (
	"github.com/jinzhu/gorm"
)

// ArticleTag 文章标签中间表
type ArticleTag struct {
	gorm.Model
	TagID     int `gorm:"int"`
	ArticleID int `gorm:"int"`
}
