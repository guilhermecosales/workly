package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermecosales/workly/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:openingId", handler.RetrieveOpeningById)

		v1.GET("/openings", handler.RetrieveAllOpenings)

		v1.POST("/openings", handler.CreateOpening)

		v1.PUT("/openings/:openingId", handler.UpdateOpening)

		v1.DELETE("/openings/:openingId", handler.DeleteOpening)
	}

}
