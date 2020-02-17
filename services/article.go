package services

import (
	"BLOG/model"
	"BLOG/repository"
)

// ArticleService 文章服务
type ArticleService interface {
	Get(id int64) *repository.Article
	GetPrev(id int64) *model.Article
	GetNext(id int64) *model.Article
	GetList(page int, pageSize int) (repository.ArticleList, error)
	Create(Article *model.Article) (uint, error)
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

func (s *articleService) Get(id int64) *repository.Article {
	return s.repo.Get(id)
}

func (s *articleService) GetPrev(id int64) *model.Article {
	return s.repo.GetPrev(id)
}

func (s *articleService) GetNext(id int64) *model.Article {
	return s.repo.GetNext(id)
}

func (s *articleService) GetList(page int, pageSize int) (repository.ArticleList, error) {
	return s.repo.GetList(page, pageSize)
}

func (s *articleService) Create(Article *model.Article) (uint, error) {
	ID, err := s.repo.Create(Article)
	return ID, err
}
