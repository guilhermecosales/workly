package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpening(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"message": "Create Opening",
	})
}
