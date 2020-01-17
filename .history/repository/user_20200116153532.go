package repository

import (
	"BLOG/entitys"
	"BLOG/util/db"
	"fmt"
)

type userRepository struct {
}

// UserRepository user DAO
var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Get(id int64) *entitys.User {
	ret := &entitys.User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
