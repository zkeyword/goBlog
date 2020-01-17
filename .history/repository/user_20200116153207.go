package repository

import (
	"BLOG/entitys"
	D "BLOG/util/db"
	"fmt"
)

var LinkRepository = newLinkRepository()

func newLinkRepository() *linkRepository {
	return &linkRepository{}
}

type linkRepository struct {
}

func (r *linkRepository) Get(id int64) *entitys.User {

	// res := db.GetMysql().Where("id = ?", userId).Find(&users)

	ret := &entitys.User{}
	fmt.Println(11111, D.GetMysql())

	if err := D.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
