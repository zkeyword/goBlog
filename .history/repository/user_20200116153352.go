package repository

import (
	"BLOG/entitys"
	"BLOG/util/db"
	"fmt"
)

type linkRepository struct {
}

var LinkRepository = newLinkRepository()

func newLinkRepository() *linkRepository {
	return &linkRepository{}
}

func (r *linkRepository) Get(id int64) *entitys.User {
	ret := &entitys.User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
