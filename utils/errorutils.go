package utils

import "github.com/gin-gonic/gin"

type ErrorUtils struct{}

func (*ErrorUtils) SendJSONError(c *gin.Context, status int, message string, err error) {
	c.JSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}
