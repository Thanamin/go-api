package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thanamin/go-api/internal/infrastructures/adapters/hash"
	"github.com/thanamin/go-api/internal/infrastructures/config/environments"
)

type AuthService struct {
	passwordHasher hash.PasswordHasher
}

type JWTClaims struct {
	StaffID    int    `json:"staff_id"`
	HospitalID int    `json:"hospital_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthService() *AuthService {
	return &AuthService{
		passwordHasher: hash.NewPasswordHasher(),
	}
}

func (s *AuthService) EncryptPassword(plainPassword string) (string, error) {
	if plainPassword == "" {
		return "", fmt.Errorf("password cannot be empty")
	}

	hashedPassword, err := s.passwordHasher.Hash(plainPassword)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt password: %w", err)
	}

	return hashedPassword, nil
}

func (s *AuthService) DecryptPassword(plainPassword, hashedPassword string) error {
	if plainPassword == "" {
		return fmt.Errorf("password cannot be empty")
	}
	if hashedPassword == "" {
		return fmt.Errorf("hashed password cannot be empty")
	}

	if err := s.passwordHasher.Verify(plainPassword, hashedPassword); err != nil {
		return fmt.Errorf("password verification failed: %w", err)
	}

	return nil
}

func (s *AuthService) GenerateToken(staffID, hospitalID int, username, email string) (string, error) {
	jwtSecret := environments.GetJWTSecret()
	jwtExpiration := time.Duration(environments.GetJWTExpirationHours()) * time.Hour

	now := time.Now()
	claims := JWTClaims{
		StaffID:    staffID,
		HospitalID: hospitalID,
		Username:   username,
		Email:      email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtExpiration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*JWTClaims, error) {
	jwtSecret := environments.GetJWTSecret()

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
