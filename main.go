package main

import (
	"test-app/routes"

	"github.com/gin-gonic/gin"
)

// Main function
func main() {

	// Setup gin default router
	router := gin.Default()

	// Connect API router
	routes.ApiRouter(router.Group("/api"))

	// Run the server
	router.Run(":4000")
}
