package helpers

import "github.com/gin-gonic/gin"

func ResponseJson(c *gin.Context, code int, status bool, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"status":  status,
		"message": msg,
		"data":    data,
	})
}
