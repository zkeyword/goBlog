package services

import (
	"BLOG/model"
	"BLOG/repository"
)

// ArticleService 文章服务
type ArticleService interface {
	Get(id int64) *model.Article
	GetPrev(id int64) *model.Article
	GetNext(id int64) *model.Article
	GetList() []model.Article
	Create(Article *model.Article)
}

type articleService struct {
	repo *repository.ArticleRepository
}

// NewArticleService 实例化ArticleService
var NewArticleService = newArticleService()

func newArticleService() ArticleService {
	return &articleService{
		repo: repository.NewArticleRepository(),
	}
}

func (s *articleService) Get(id int64) *model.Article {
	return s.repo.Get(id)
}

func (s *articleService) GetPrev(id int64) *model.Article {
	return s.repo.GetPrev(id)
}

func (s *articleService) GetNext(id int64) *model.Article {
	return s.repo.GetNext(id)
}

func (s *articleService) GetList() []model.Article {
	return s.repo.GetList()
}

func (s *articleService) Create(Article *model.Article) {
	s.repo.Create(Article)
}
