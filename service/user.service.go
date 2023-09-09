package service

import (
	"github.com/bricsport/common/utils"
	"github.com/bricsport/entity"
	"github.com/bricsport/repository"
)

type UserService struct {
	Repo repository.UserRepositry
}

func NewUserService(repo repository.UserRepositry) *UserService {
	return &UserService{Repo: repo}
}

func (srvc *UserService) CreateUser(usr *entity.User) (*entity.User, error) {
	hash, hashErr := utils.HashPassword(usr.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	usr.Password = hash

	result, err := srvc.Repo.Create(usr)
	if err != nil {
		return nil, err
	}

	return result, nil
}
