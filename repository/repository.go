package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository UserRepositry
}

func InitRepositories(tx *gorm.DB) *Repository {
	return &Repository{
		UserRepository: &postgresUserRepository{Tx: tx},
	}
}
