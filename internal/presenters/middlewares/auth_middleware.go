package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thanamin/go-api/internal/domains/service/auth"
)

// AuthMiddleware validates JWT token and stores claims in context
func AuthMiddleware(authService *auth.AuthService) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check Bearer prefix
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store claims in context
		c.Set("auth", claims)
		c.Set("staff_id", claims.StaffID)
		c.Set("hospital_id", claims.HospitalID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// GetAuthClaims retrieves JWT claims from context
func GetAuthClaims(c *gin.Context) (*auth.JWTClaims, bool) {
	claims, exists := c.Get("auth")
	if !exists {
		return nil, false
	}

	authClaims, ok := claims.(*auth.JWTClaims)
	return authClaims, ok
}

// GetStaffID retrieves staff ID from context
func GetStaffID(c *gin.Context) (int, bool) {
	staffID, exists := c.Get("staff_id")
	if !exists {
		return 0, false
	}

	id, ok := staffID.(int)
	return id, ok
}

// GetHospitalID retrieves hospital ID from context
func GetHospitalID(c *gin.Context) (int, bool) {
	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		return 0, false
	}

	id, ok := hospitalID.(int)
	return id, ok
}
