package routes

import (
	"test-app/controllers"

	"github.com/gin-gonic/gin"
)

// Routers
func ApiRouter(router *gin.RouterGroup) *gin.RouterGroup {

	// API Routes
	apiController := controllers.APIController{}
	router.POST("/images/transform", apiController.TransformImage)
	router.POST("/images/filter", apiController.ApplyFilter)
	router.POST("/images/meta", apiController.Meta)

	return router
}
