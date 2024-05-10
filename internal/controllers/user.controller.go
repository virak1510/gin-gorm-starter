package controllers

import (
	"myFirstGoGin/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userServices *services.UserServices
}

func New(db *gorm.DB) *UserController {
	return &UserController{
		userServices: services.New(db),
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users, err = uc.userServices.GetUsers()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "users": users})
}
