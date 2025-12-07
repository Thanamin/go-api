package staff

import (
	"context"
	"testing"

	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/domains/service/auth"
)

func TestCreateStaff_Validate_Fails(t *testing.T) {
	ctx := context.Background()
	authService := auth.NewAuthService()

	uc := &CreateStaffUseCase{StaffRepo: nil, HospitalRepo: nil, authService: authService}

	// Missing username
	s := &entity.Staff{PasswordHash: "pass", Email: "a@b.com"}
	if err := uc.validate(ctx, s); err == nil {
		t.Fatalf("expected validation error for missing username")
	}

	// Missing password
	s2 := &entity.Staff{Username: "u", Email: "a@b.com"}
	if err := uc.validate(ctx, s2); err == nil {
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
