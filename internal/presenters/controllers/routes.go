package controllers

import (
	"github.com/gin-gonic/gin"

	patient "github.com/thanamin/go-api/internal/presenters/controllers/patient"
	staff "github.com/thanamin/go-api/internal/presenters/controllers/staff"
)

// Controllers holds all controller instances
type Controllers struct {
	Staff   *staff.StaffController
	Patient *patient.PatientController
}

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, controller *Controllers) {
	// Staff routes
	controller.Staff.SetupRoutes(router.Group("/staff"))

	// Patient routes
	controller.Patient.SetupRoutes(router.Group("/patient"))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Go API is running",
		})
	})
}
