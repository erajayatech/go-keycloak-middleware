package keycloakmiddleware

import (
	"github.com/gin-gonic/gin"
)

func (middleware *middleware) wrapper(context *gin.Context, message interface{}, data interface{}) map[string]interface{} {
	if middleware.wrapperCode == 2 {
		return middleware.traceableWrapper(context, message, data)
	} else if middleware.wrapperCode == 1 {
		return middleware.standardWrapper(message, data)
	} else {
		return middleware.defaultWrapper(message, data)
	}
}

func (middleware *middleware) defaultWrapper(message interface{}, data interface{}) map[string]interface{} {
	return gin.H{
		"error_message": message,
		"data":          data,
	}
}

func (middleware *middleware) standardWrapper(message interface{}, data interface{}) map[string]interface{} {
	return gin.H{
		"message": message,
		"data":    data,
	}
}

func (middleware *middleware) traceableWrapper(context *gin.Context, message interface{}, data interface{}) map[string]interface{} {
	return gin.H{
		"id":      context.GetHeader("X-Trace-Id"),
		"appName": getEnvOrDefault("APP_NAME", nil),
		"version": getEnvOrDefault("APP_VERSION", nil),
		"build":   getEnvOrDefault("BUILD", nil),
		"message": message,
		"data":    data,
	}
}
