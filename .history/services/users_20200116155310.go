package services

import (
	"BLOG/entitys"
	"BLOG/repository"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (this *userService) Get(id int64) *entitys.User {
	return repository.UserRepository.Get(id)
}
