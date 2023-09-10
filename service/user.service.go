package service

import (
	"net/http"

	"github.com/bricsport/common/exceptions"
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

func (srvc *UserService) CreateUser(usr *entity.User) (*entity.User, *exceptions.Exception) {
	hash, hashErr := utils.HashPassword(usr.Password)
	if hashErr != nil {
		return nil, exceptions.Throw(http.StatusInternalServerError, "ERROR HASHING PASSWORD: "+hashErr.Error())
	}
	usr.Password = hash

	result, err := srvc.Repo.Create(usr)
	if err != nil {
		return nil, exceptions.Throw(http.StatusExpectationFailed, "USER REGISTER WAS FAILED: "+err.Error())
	}

	return result, nil
}
