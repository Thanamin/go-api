package staff

import (
	"testing"

	entstaff "github.com/thanamin/go-api/internal/domains/entity/staff"
	"github.com/thanamin/go-api/internal/domains/service/auth"
)

func TestVerifyPassword_SucceedsAndFails(t *testing.T) {
	authService := auth.NewAuthService()
	plain := "correctpass"
	hashed, err := authService.EncryptPassword(plain)
	if err != nil {
		t.Fatalf("encrypt error: %v", err)
	}

	uc := &LoginStaffUseCase{StaffRepo: nil, authService: authService}

	st := &entstaff.Staff{Email: "x@x", PasswordHash: hashed}
	if err := uc.verifyPassword(st, plain); err != nil {
		t.Fatalf("expected password to verify, got %v", err)
	}

	if err := uc.verifyPassword(st, "wrong"); err == nil {
		t.Fatalf("expected verify to fail for wrong password")
	}
}
