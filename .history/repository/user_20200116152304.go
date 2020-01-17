package repository

import (
	"BLOG/entitys"
	D "BLOG/util/db"
	"fmt"
)

var DB = D.GetMysql()
var LinkRepository = newLinkRepository()

func newLinkRepository() *linkRepository {
	return &linkRepository{}
}

type linkRepository struct {
}

func (this *linkRepository) Get(id int64) *entitys.User {
	ret := &entitys.User{}
	fmt.Println(ret)
	if err := DB.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
