package service

import "github.com/bricsport/repository"

type Service struct {
	UserService *UserService
}

func InitServices(repos repository.Repository) *Service {
	return &Service{
		UserService: &UserService{Repo: repos.UserRepository},
	}
}
