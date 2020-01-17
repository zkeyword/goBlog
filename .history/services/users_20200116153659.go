package services

import (
	"BLOG/entitys"
	"BLOG/repository"
	"fmt"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (this *userService) Get(id int64) *entitys.User {
	fmt.Println(repository.UserRepository.Get(id))
	return repository.UserRepository.Get(id)
}
