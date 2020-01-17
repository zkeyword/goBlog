package model

import (
	"github.com/jinzhu/gorm"
)

// Article 文章表
type Article struct {
	gorm.Model
	Title    string `gorm:"size:128;not null;"`
	Content  string `gorm:"type:longtext;not null;"`
	AuthorID int    `gorm:"INT"`
}
