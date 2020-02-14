package repository

import (
	"BLOG/model"
	"BLOG/util/db"
	"fmt"
)

// ArticleRepository article DAO
type ArticleRepository struct {
}

// NewArticleRepository 实例化DAO
func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

// Get 获取文章详情
func (r *ArticleRepository) Get(id int64) *model.Article {
	ret := &model.Article{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

// GetPrev 获取上一篇文章详情
func (r *ArticleRepository) GetPrev(id int64) *model.Article {
	ret := &model.Article{}

	if err := db.GetMysql().First(ret, "id < ?", id).Order("DESC").Limit(1).Error; err != nil {
		return nil
	}
	return ret
}

// GetNext 获取下一篇文章详情
func (r *ArticleRepository) GetNext(id int64) *model.Article {
	ret := &model.Article{}

	if err := db.GetMysql().First(ret, "id > ?", id).Order("ASC").Limit(1).Error; err != nil {
		return nil
	}
	return ret
}

// GetList 获取文章列表
// func (r *ArticleRepository) GetList() ([]model.Article, error) {
// 	ret := make([]model.Article, 0)
// 	err := db.GetMysql().Find(ret).Error
// 	return ret, err
// }
func (r *ArticleRepository) GetList() []model.Article {
	ret := make([]model.Article, 0)
	if err := db.GetMysql().Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}

// Create 创建文章
func (r *ArticleRepository) Create(t *model.Article) (err error) {
	err = db.GetMysql().Create(t).Error
	return
}
