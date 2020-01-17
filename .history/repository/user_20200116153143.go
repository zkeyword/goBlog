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

func (r *linkRepository) Get(id int64) *entitys.User {

	// res := db.GetMysql().Where("id = ?", userId).Find(&users)

	ret := &entitys.User{}
	fmt.Println(11111, r)

	if err := DB.First(ret, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
