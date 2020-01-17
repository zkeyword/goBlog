package services

import (
	"BLOG/entitys"
	"BLOG/repository"
)

var LinkService = newLinkService()

func newLinkService() *linkService {
	return &linkService{}
}

type linkService struct {
}

func (this *linkService) Get(id int64) *entitys.User {
	return repository.LinkRepository.Get(id)
}
