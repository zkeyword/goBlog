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

func (this *articleService) Get(id int64) *model.Article {
	return repository.ArticleRepository.Get(id)
}

func (this *articleService) GetList() []model.Article {
	return repository.ArticleRepository.GetList()
}

func (this *articleService) Create(Article model.Article) {
	return repository.ArticleRepository.Create(Article)
}
