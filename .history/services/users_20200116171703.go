package services

import (
	"BLOG/model"
	"BLOG/repository"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (this *userService) Get(id int64) *model.User {
	return repository.UserRepository.Get(id)
}
