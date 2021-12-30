package keycloakmiddleware

import (
	"github.com/gin-gonic/gin"
)

func wrapper(message interface{}, data interface{}) map[string]interface{} {
	return gin.H{
		"error_message": message,
		"data":          data,
	}
}
