package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetrieveOpeningById(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Retrieve Opening by Identifier",
	})
}
