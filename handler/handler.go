package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetrieveAllOpenings(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Post Opening",
	})
}

func RetrieveOpeningById(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Retrieve Opening by Identifier",
	})
}

func CreateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Create Opening",
	})
}

func UpdateOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Update Opening",
	})
}

func DeleteOpening(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Delete Opening",
	})
}
