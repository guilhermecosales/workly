package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Update Opening",
	})
}
