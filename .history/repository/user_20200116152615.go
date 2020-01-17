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
	fmt.Println(11111, DB.First(ret, "id = ?", id))

	if err := DB.First(ret, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
