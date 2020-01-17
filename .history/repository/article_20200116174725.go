package repository

import (
	"BLOG/model"
	"BLOG/util/db"
	"fmt"
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

// func (r *articleRepository) GetList() ([]model.Article, error) {
// 	ret := make([]model.Article, 0)
// 	err := db.GetMysql().Find(ret).Error
// 	return ret, err
// }
func (r *articleRepository) GetList() []model.Article {
	ret := make([]model.Article, 0)
	if err := db.GetMysql().Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
