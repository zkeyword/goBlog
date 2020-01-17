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

func (this *articleService) Get(id int64) *model.User {
	return repository.ArticleRepository.Get(id)
}
