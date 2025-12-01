package bootstrap

import (
	"github.com/thanamin/go-api/internal/domains/service/auth"
)

// Services holds long-lived service instances
type Services struct {
	Auth *auth.AuthService
}

// InitServices initializes shared services
func InitServices() *Services {
	return &Services{
		Auth: auth.NewAuthService(),
	}
}
