package keycloakmiddleware

import (
	"github.com/gin-gonic/gin"
)

func (middleware *middleware) wrapper(status int, context *gin.Context, message interface{}, data interface{}) map[string]interface{} {
	if middleware.wrapperCode == 2 {
		return middleware.traceableWrapper(context, message, data)
	} else if middleware.wrapperCode == 1 {
		return middleware.standardWrapper(message, data)
	} else {
		return middleware.defaultWrapper(status, message, data)
	}
}

func (middleware *middleware) defaultWrapper(status int, message interface{}, data interface{}) map[string]interface{} {
	return gin.H{
		"status":        status,
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
	var traceID = context.GetString("X-Trace-Id")
	return gin.H{
		"id":      traceID,
		"appName": getEnvOrDefault("APP_NAME", nil),
		"version": getEnvOrDefault("APP_VERSION", nil),
		"build":   getEnvOrDefault("BUILD", nil),
		"message": message,
		"data":    data,
	}
}
