package staff

import (
	"testing"

	entstaff "github.com/thanamin/go-api/internal/domains/entity/staff"
	"github.com/thanamin/go-api/internal/domains/service/auth"
)

func TestCreateStaff_Validate_Fails(t *testing.T) {

	authService := auth.NewAuthService()

	uc := &CreateStaffUseCase{StaffRepo: nil, authService: authService}

	// Missing username
	s := &entstaff.Staff{PasswordHash: "pass", Email: "a@b.com"}
	if err := uc.validate(s); err == nil {
		t.Fatalf("expected validation error for missing username")
	}

	// Missing password
	s2 := &entstaff.Staff{Username: "u", Email: "a@b.com"}
	if err := uc.validate(s2); err == nil {
		t.Fatalf("expected validation error for missing password")
	}
}

func TestCreateStaff_EncryptPassword(t *testing.T) {
	authService := auth.NewAuthService()
	plain := "secret"
	hashed, err := authService.EncryptPassword(plain)
	if err != nil {
		t.Fatalf("encrypt error: %v", err)
	}
	if hashed == plain {
		t.Fatalf("expected hashed password to differ from plain")
	}
}
