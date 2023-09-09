package repository

import (
	"errors"

	"github.com/bricsport/entity"
	"gorm.io/gorm"
)

type UserRepositry interface {
	// Login(usr entity.User) (bool, error)
	Create(usr *entity.User) (*entity.User, error)
}

type postgresUserRepository struct {
	Tx *gorm.DB
}

func NewUserRepository(db *gorm.DB) *postgresUserRepository {
	return &postgresUserRepository{Tx: db}
}

// func (repo *postgresUserRepository) Cre(user entity.User) (bool, error) {
// 	dbUser := tx
// }

func (repo *postgresUserRepository) Create(usr *entity.User) (*entity.User, error) {
	result := repo.Tx.Create(usr)

	if result.Error != nil || usr.ID <= 0 {
		return nil, errors.New("ERROR CREATING RECORD" + result.Error.Error())
	}

	return usr, nil
}
