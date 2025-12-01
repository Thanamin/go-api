package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/domains/service/auth"
	staffUseCases "github.com/thanamin/go-api/internal/domains/usecases/staff"
)

type StaffController struct {
	createUseCase *staffUseCases.CreateStaffUseCase
	loginUseCase  *staffUseCases.LoginStaffUseCase
	authService   *auth.AuthService
}

func NewStaffController(
	createUseCase *staffUseCases.CreateStaffUseCase,
	loginUseCase *staffUseCases.LoginStaffUseCase,
	authService *auth.AuthService,
) *StaffController {
	return &StaffController{
		createUseCase: createUseCase,
		loginUseCase:  loginUseCase,
		authService:   authService,
	}
}

func (c *StaffController) Create(ctx *gin.Context) {
	var req struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		HospitalID   int    `json:"hospital_id"`
		FirstNameTh  string `json:"first_name_th"`
		MiddleNameTh string `json:"middle_name_th"`
		LastNameTh   string `json:"last_name_th"`
		FirstNameEn  string `json:"first_name_en"`
		MiddleNameEn string `json:"middle_name_en"`
		LastNameEn   string `json:"last_name_en"`
		Email        string `json:"email"`
		Position     string `json:"position"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Map request to Staff entity
	staff := &entity.Staff{
		Username:     req.Username,
		PasswordHash: req.Password, // Will be hashed in usecase
		HospitalID:   req.HospitalID,
		FirstNameTh:  req.FirstNameTh,
		MiddleNameTh: req.MiddleNameTh,
		LastNameTh:   req.LastNameTh,
		FirstNameEn:  req.FirstNameEn,
		MiddleNameEn: req.MiddleNameEn,
		LastNameEn:   req.LastNameEn,
		Email:        req.Email,
		Position:     req.Position,
	}

	if err := c.createUseCase.Execute(ctx.Request.Context(), staff); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, staff)
}

// Login handles staff authentication and returns JWT token
func (c *StaffController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	staff, err := c.loginUseCase.Execute(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := c.authService.GenerateToken(staff.ID, staff.HospitalID, staff.Username, staff.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":      token,
		"token_type": "Bearer",
	})
}

// SetupRoutes configures staff routes
func (c *StaffController) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/create", c.Create)
	router.POST("/login", c.Login)
}
