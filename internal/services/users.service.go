package services

import (
	"github.com/virak1510/gin-gorm-starter/internal/schemas"

	"gorm.io/gorm"
)

type UserServices struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserServices {
	return &UserServices{db: db}
}

func (us *UserServices) GetUsers() ([]*schemas.User, error) {
	var users []*schemas.User
	err := us.db.Raw("SELECT * FROM users").Scan(&users).Error
	if err != nil {
		// Handle other errors
		return nil, err
	}

	return users, nil
}
