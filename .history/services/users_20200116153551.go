package services

import (
	"BLOG/entitys"
	"BLOG/repository"
	"fmt"
)

var LinkService = newLinkService()

func newLinkService() *linkService {
	return &linkService{}
}

type linkService struct {
}

func (this *linkService) Get(id int64) *entitys.User {
	fmt.Println(repository.UserRepository.Get(id))
	return repository.UserRepository.Get(id)
}
