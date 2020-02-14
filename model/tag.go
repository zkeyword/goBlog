package model

import (
	"github.com/jinzhu/gorm"
)

// Tag 标签表
type Tag struct {
	gorm.Model
	ArticleID int    `gorm:"int"`
	Title     string `gorm:"size:128;not null;"`
}
