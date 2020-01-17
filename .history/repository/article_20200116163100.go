package repository

import (
	"BLOG/model"
	"BLOG/util/db"
)

type articleRepository struct {
}

// UserRepository user DAO
var UserRepository = newArticleRepository()

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (r *articleRepository) Get(id int64) *model.User {
	ret := &model.User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
