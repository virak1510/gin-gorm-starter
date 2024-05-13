package utils

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func (*ApiResponse) Response(c *gin.Context, data interface{}, message string, status int) {
	response := ApiResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
	c.JSON(status, response)
}
