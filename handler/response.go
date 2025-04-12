package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendErrorResponse(context *gin.Context, statusCode int, message string) {
	context.Header("Content-Type", "application/json")

	context.JSON(statusCode, gin.H{
		"message":   message,
		"errorCode": statusCode,
	})
}

func sendSuccessResponse(context *gin.Context, message string, data interface{}) {
	context.Header("Content-Type", "application/json")

	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
