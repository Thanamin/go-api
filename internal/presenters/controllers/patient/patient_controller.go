package patient

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanamin/go-api/internal/domains/service/auth"
	patientUseCases "github.com/thanamin/go-api/internal/domains/usecases/patient"
	"github.com/thanamin/go-api/internal/presenters/middlewares"
)

type PatientController struct {
	searchUseCase *patientUseCases.SearchPatientUseCase
	authService   *auth.AuthService
}

func NewPatientController(searchUseCase *patientUseCases.SearchPatientUseCase, authService *auth.AuthService) *PatientController {
	return &PatientController{
		searchUseCase: searchUseCase,
		authService:   authService,
	}
}

// Search searches for patients by query parameters
func (c *PatientController) Search(ctx *gin.Context) {
	// Get authenticated staff info from context
	authClaims, exists := middlewares.GetAuthClaims(ctx)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	hospitalID := authClaims.HospitalID

	// Build search parameters from query
	params := patientUseCases.SearchParams{
		HospitalID:  strconv.Itoa(hospitalID),
		NationalID:  ctx.Query("national_id"),
		PassportID:  ctx.Query("passport_id"),
		FirstName:   ctx.Query("first_name"),
		LastName:    ctx.Query("last_name"),
		DateOfBirth: ctx.Query("date_of_birth"),
		PhoneNumber: ctx.Query("phone_number"),
		Email:       ctx.Query("email"),
	}

	// Support a generic `name` parameter as a shortcut (search across first/last)
	if name := ctx.Query("name"); name != "" {
		params.FirstName = name
	}

	// Parse pagination with defaults
	if limit := ctx.Query("limit"); limit != "" {
		if val, err := strconv.Atoi(limit); err == nil {
			params.Limit = val
		}
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	if offset := ctx.Query("offset"); offset != "" {
		if val, err := strconv.Atoi(offset); err == nil {
			params.Offset = val
		}
	}

	// Execute search use case
	patients, err := c.searchUseCase.Execute(ctx.Request.Context(), &params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search patients"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   patients,
		"limit":  params.Limit,
		"offset": params.Offset,
	})
}

// SetupRoutes configures patient routes
func (c *PatientController) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/search", middlewares.AuthMiddleware(c.authService), c.Search)
}
