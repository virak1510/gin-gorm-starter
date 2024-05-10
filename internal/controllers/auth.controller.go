package controllers

import (
	"net/http"

	"github.com/virak1510/gin-gorm-starter/internal/models"
	"github.com/virak1510/gin-gorm-starter/internal/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	returnUser := services.Login(&user)

	c.JSON(http.StatusOK, gin.H{"message": "hi", "user": returnUser})
}
