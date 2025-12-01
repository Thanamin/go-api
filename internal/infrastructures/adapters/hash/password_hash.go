package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/thanamin/go-api/internal/infrastructures/config/environments"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password, hash string) error
}

type BcryptHasher struct {
	cost int
}

func NewPasswordHasher() PasswordHasher {
	method := environments.GetHashMethod()

	switch method {
	case "bcrypt":
		return &BcryptHasher{
			cost: environments.GetBcryptCost(),
		}
	default:
		return &BcryptHasher{
			cost: 10,
		}
	}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

func (h *BcryptHasher) Verify(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}
