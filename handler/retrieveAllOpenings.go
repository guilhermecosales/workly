package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetrieveAllOpenings(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Retrieve All Opening",
	})
}
