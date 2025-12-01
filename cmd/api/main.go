package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/thanamin/go-api/cmd/api/bootstrap"
	"github.com/thanamin/go-api/internal/infrastructures/config/environments"
	"github.com/thanamin/go-api/internal/presenters/controllers"
)

func main() {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Initialize database connection
	databaseInstance := bootstrap.InitDatabase()

	// Initialize mappers
	mapperInstance := bootstrap.InitMappers()

	// Initialize services
	serviceInstance := bootstrap.InitServices()

	// Initialize use cases (after services so we can inject them)
	useCaseInstance := bootstrap.InitUseCases(databaseInstance, mapperInstance, serviceInstance)

	// Initialize controllers
	controllerInstance := bootstrap.InitControllers(useCaseInstance, serviceInstance)

	// Setup routes
	controllers.SetupRoutes(router, controllerInstance)

	// Start server
	port := environments.GetPort()
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📍 API endpoints available at http://localhost:%s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
