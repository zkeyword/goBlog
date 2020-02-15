package services

import (
	"BLOG/model"
	"BLOG/repository"
)

// ArticleTagService 文章服务
type ArticleTagService interface {
	Get(id int64) *model.ArticleTag
	GetList() []model.ArticleTag
	Create(ArticleTag *model.ArticleTag)
}

type articleTagService struct {
	repo *repository.ArticleTagRepository
}

// NewArticleTagService 实例化ArticleTagService
var NewArticleTagService = newArticleTagService()

func newArticleTagService() ArticleTagService {
	return &articleTagService{
		repo: repository.NewArticleTagRepository(),
	}
}

func (s *articleTagService) Get(id int64) *model.ArticleTag {
	return s.repo.Get(id)
}

func (s *articleTagService) GetList() []model.ArticleTag {
	return s.repo.GetList()
}

func (s *articleTagService) Create(ArticleTag *model.ArticleTag) {
	s.repo.Create(ArticleTag)
}
