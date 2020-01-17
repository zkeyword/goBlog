package repository

import (
	"BLOG/entitys"
	D "BLOG/util/db"

	"github.com/jinzhu/gorm"
)

var DB = D.GetMysql()
var LinkRepository = newLinkRepository()

func newLinkRepository() *linkRepository {
	return &linkRepository{}
}

type linkRepository struct {
}

func (this *linkRepository) Get(db *gorm.DB, id int64) *entitys.User {
	ret := &entitys.User{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
