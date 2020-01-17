package repository

import (
	"BLOG/model"
	"BLOG/util/db"
)

type articleRepository struct {
}

// ArticleRepository article DAO
var ArticleRepository = newArticleRepository()

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (r *articleRepository) Get(id int64) *model.Article {
	ret := &model.Article{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *articleRepository) GetList() *model.Article {
	ret := &model.Article{}

	if err := db.GetMysql().Find(ret).Error; err != nil {
		return nil
	}
	return ret
}
