package repository

import (
	"BLOG/model"
	"BLOG/util/db"
	"fmt"
)

// TagRepository article DAO
type TagRepository struct {
}

// NewTagRepository 实例化DAO
func NewTagRepository() *TagRepository {
	return &TagRepository{}
}

// Get 获取文章详情
func (r *TagRepository) Get(id int64) *model.Tag {
	ret := &model.Tag{}

	fmt.Println(r)

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

// GetList 获取文章列表
// func (r *TagRepository) GetList() ([]model.Tag, error) {
// 	ret := make([]model.Tag, 0)
// 	err := db.GetMysql().Find(ret).Error
// 	return ret, err
// }
func (r *TagRepository) GetList() []model.Tag {
	ret := make([]model.Tag, 0)
	if err := db.GetMysql().Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}

// Create 创建文章
func (r *TagRepository) Create(t *model.Tag) (err error) {
	err = db.GetMysql().Create(t).Error
	return
}
