package services

import (
	"BLOG/model"
	"BLOG/repository"
	"BLOG/util/business_errors"
	"BLOG/util/crypto"
)

// UserService 用户服务
type UserService interface {
	Create(user *model.User) error
	Login(username, password string) (*model.User, error)

	FindByID(userID uint) (*model.User, error)
	FindByIDs(ids []uint) []model.User

	FindByUsername(username string) (*model.User, error)
	FindAllUsers(limit, offset int) ([]model.User, error)

	CheckUsernameExist(username string) bool
	CheckEmailExist(email string) bool
	CheckUserIsLockByID(userID uint) bool

	// Actions
	LockUser(user *model.User)
	UnLockUser(user *model.User)
}

type userService struct {
	repo *repository.UserRepository
}

// NewUserService 实例化articleService
var NewUserService = newUserService()

func newUserService() UserService {
	return &userService{
		repo: repository.NewUserRepository(),
	}
}

func (s *userService) Get(id int64) *model.User {
	return s.repo.Get(id)
}


func (s *userService) Create(user *model.User) error {
	if s.CheckEmailExist(user.Username) {
		return business_errors.UsernameAlreadyExists
	}

	if s.CheckEmailExist(user.Email) {
		return business_errors.EmailAlreadyExists
	}

	if len(user.Password) < 8 {
		return business_errors.PasswordLessThanEightCharacters
	}

	pw, err := crypto.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pw
	return s.repo.Create(user)
}

func (s *userService) Login(username, password string) (*model.User, error) {
	user, err := s.repo.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	if user.ID < 0 {
		return nil, business_errors.UsernameNotExist
	}

	isCheck := crypto.CheckPassword(password, user.Password)
	if !isCheck {
		return nil, business_errors.PasswordError
	}
	return user, nil
}

func (s *userService) FindByID(userID uint) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	return user, err
}

func (s *userService) FindByIDs(ids []uint) []model.User {
	users, _ := s.repo.FindByIDs(ids)
	return users
}

func (s *userService) FindByUsername(username string) (*model.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *userService) FindAllUsers(limit, offset int) ([]model.User, error) {
	return s.repo.FindAllUsers(limit, offset)
}

func (s *userService) CheckUsernameExist(username string) bool {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func (s *userService) CheckEmailExist(email string) bool {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func (s *userService) LockUser(user *model.User) {
	user.Lock = true
	_ = s.repo.Update(user)
}

func (s *userService) UnLockUser(user *model.User) {
	user.Lock = false
	_ = s.repo.Update(user)
}

func (s *userService) CheckUserIsLockByID(userID uint) bool {
	user, _ := s.FindByID(userID)
	return user.Lock
}
