package repository

import (
	"BLOG/model"
	"BLOG/util/business_errors"
	"BLOG/util/db"
	"fmt"

	"github.com/jinzhu/gorm"
)

// UserRepository user
type UserRepository struct {
}

// NewUserRepository user DAO
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Get 根据用户id查询
func (r *UserRepository) Get(id int64) *model.User {
	ret := &model.User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

//
func (r *UserRepository) Create(user *model.User) error {
	if err := db.GetMysql().Create(user).Error; err != nil {
		fmt.Printf("CreateUserError:%s", err)
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user = new(model.User)
	result := db.GetMysql().Where("id=?", id).First(user)
	err := result.Error
	return user, err
}

func (r *UserRepository) FindByIDs(ids []uint) ([]model.User, error) {
	var users = make([]model.User, 0)
	err := db.GetMysql().Where("id in (?)", ids).Find(&users).Error
	for idx, _ := range users {
		users[idx].Password = ""
	}
	return users, err
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user = new(model.User)
	result := db.GetMysql().Where("username = ?", username).Find(&user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, business_errors.UsernameNotExist
	}
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user = new(model.User)
	result := db.GetMysql().Where("email=?", email).Find(user)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, business_errors.UsernameNotExist
	}
	return user, err
}

func (r *UserRepository) FindAllUsers(limit, offset int) ([]model.User, error) {
	var users = make([]model.User, 0)
	err := db.GetMysql().Find(&users).Error
	for idx, _ := range users {
		users[idx].Password = ""
	}
	return users, err
}

func (r *UserRepository) Update(users *model.User) error {
	return db.GetMysql().Save(users).Error
}

func (r *UserRepository) Delete(user *model.User) error {
	err := db.GetMysql().Unscoped().Delete(user).Error
	return err
}

func (r *UserRepository) DeleteByID(userID uint) error {
	var user = new(model.User)
	user.ID = userID
	return db.GetMysql().Unscoped().Delete(user).Error
}
