package controllers

import (
	"myFirstGoGin/internal/models"
	"myFirstGoGin/internal/services"
	"net/http"

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
