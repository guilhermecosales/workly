package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the Gin router
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	router.Run()
}
