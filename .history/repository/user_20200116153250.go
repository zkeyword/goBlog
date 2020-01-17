package repository

import (
	"BLOG/entitys"
	"BLOG/util/db"
	"fmt"
)

var LinkRepository = newLinkRepository()

func newLinkRepository() *linkRepository {
	return &linkRepository{}
}

type linkRepository struct {
}

func (r *linkRepository) Get(id int64) *entitys.User {
	ret := &entitys.User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
