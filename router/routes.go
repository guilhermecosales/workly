package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:openingId", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET Opening by Identifier",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET Openings",
			})
		})

		v1.POST("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
			})
		})

		v1.PUT("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "PUT Opening",
			})
		})

		v1.DELETE("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "DELETE Opening",
			})
		})
	}

}
