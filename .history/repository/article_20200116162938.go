package repository

import (
	"BLOG/model"
	"BLOG/util/db"
)

type userRepository struct {
}

// UserRepository user DAO
var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Get(id int64) *model.User {
	ret := &model.User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
