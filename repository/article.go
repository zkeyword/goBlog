package repository

import (
	"BLOG/model"
	"BLOG/util/db"
	"fmt"
	"time"
)

// ArticleRepository article DAO
type ArticleRepository struct {
}

// ArticleList 文章response
type ArticleList struct {
	Data     []model.Article
	Total    int
	PageSize int
	Page     int
}

// Article 类型
type Article struct {
	ID        uint
	Title     string
	Content   string
	UpdatedAt time.Time
	TagID     int
	TagName   string
}

// NewArticleRepository 实例化DAO
func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

// Get 获取文章详情
func (r *ArticleRepository) Get(id int64) *Article {
	ret := &Article{}

	db.GetMysql().
		// Debug().
		Table("articles a").
		Select("a.id, a.title, a.content, a.updated_at, t.id tag_id, t.title tag_name").
		Joins("left join article_tags r on a.id = r.article_id").
		Joins("left join tags t on t.id = r.tag_id").
		Where("a.id = ?", id).
		Find(&ret)

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
func (r *ArticleRepository) GetList(page int, pageSize int) (ArticleList, error) {
	ret := make([]model.Article, 0)
	orm := db.GetMysql() //.Debug()
	list := orm.Order("created_at DESC").Offset((page - 1) * pageSize).Limit(10).Find(&ret)
	total := 0
	if err := orm.Model(&model.Article{}).Count(&total).Error; err != nil {
		fmt.Println(err)
		return ArticleList{}, err
	}
	if err := list.Error; err != nil {
		fmt.Println(err)
		return ArticleList{}, err
	}
	return ArticleList{
		Data:     ret,
		PageSize: pageSize,
		Page:     page,
		Total:    total,
	}, nil
}

// Create 创建文章
func (r *ArticleRepository) Create(t *model.Article) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}
