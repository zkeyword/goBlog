package services

import (
	"BLOG/model"
	"BLOG/repository"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

func (s *articleService) Get(id int64) *model.Article {
	return repository.ArticleRepository.Get(id)
}

func (s *articleService) GetList() []model.Article {
	return repository.ArticleRepository.GetList()
}

func (s *articleService) Create(Article *model.Article) {
	repository.ArticleRepository.Create(Article)
}
