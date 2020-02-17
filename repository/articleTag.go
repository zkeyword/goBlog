package repository

import (
	"BLOG/model"
	"BLOG/util/db"
	"fmt"
)

// ArticleTagRepository DAO
type ArticleTagRepository struct {
}

// NewArticleTagRepository 实例化DAO
func NewArticleTagRepository() *ArticleTagRepository {
	return &ArticleTagRepository{}
}

// Get 获取标签关联的文章
func (r *ArticleTagRepository) Get(id int64) *model.ArticleTag {
	ret := &model.ArticleTag{}

	fmt.Println(r)

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

// GetList 获取标签列表
func (r *ArticleTagRepository) GetList() []model.ArticleTag {
	ret := make([]model.ArticleTag, 0)
	if err := db.GetMysql().Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}

// Create 创建标签
func (r *ArticleTagRepository) Create(t *model.ArticleTag) (err error) {
	err = db.GetMysql().Debug().Create(t).Error
	return
}
