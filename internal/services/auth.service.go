package services

import (
	"github.com/virak1510/gin-gorm-starter/internal/models"
	"github.com/virak1510/gin-gorm-starter/internal/schemas"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}
func (a *AuthService) Login(user *models.User) *models.User {
	// user.Username = "modify"
	return user
}

func (a *AuthService) Register(params models.User) (uint, error) {
	var user = schemas.User{
		Username: params.Username,
		Password: params.Password,
	}
	result := a.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
