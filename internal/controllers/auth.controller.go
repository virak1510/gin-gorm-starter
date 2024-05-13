package controllers

import (
	"net/http"

	"github.com/virak1510/gin-gorm-starter/internal/models"
	"github.com/virak1510/gin-gorm-starter/internal/services"
	"github.com/virak1510/gin-gorm-starter/pkg/utils"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
	res         *utils.ApiResponse
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		authService: services.NewAuthService(db),
	}
}

func (a *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	returnUser := a.authService.Login(&user)

	a.res.Response(c, returnUser, "success", http.StatusOK)
}

func (a *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	registerdUser, err := a.authService.Register(user)
	if err != nil {
		a.res.Response(c, nil, err.Error(), http.StatusBadRequest)
	}

	a.res.Response(c, registerdUser, "success", http.StatusOK)

}
