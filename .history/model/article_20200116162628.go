package model

import (
	"github.com/jinzhu/gorm"
)

// Article 文章表
type Article struct {
	gorm.Model
	Title    string `gorm:"TEXT"`
	Content  string `gorm:"LONGTEXT"`
	AuthorId int    `gorm:"INT"`
}
