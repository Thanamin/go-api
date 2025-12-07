package staff

import (
	"context"
	"fmt"
	"strings"

	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/domains/service/auth"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories"
)

type CreateStaffUseCase struct {
	StaffRepo    *repositories.StaffRepository
	HospitalRepo *repositories.HospitalRepository
	authService  *auth.AuthService
}

func InitCreate(authService *auth.AuthService, staffRepo *repositories.StaffRepository, hospitalRepo *repositories.HospitalRepository) *CreateStaffUseCase {
	return &CreateStaffUseCase{
		authService:  authService,
		StaffRepo:    staffRepo,
		HospitalRepo: hospitalRepo,
	}
}

func (uc *CreateStaffUseCase) Execute(ctx context.Context, staff *entity.Staff) error {
	if err := uc.validate(ctx, staff); err != nil {
		return err
	}

	return uc.createInDB(ctx, staff)
}

func (uc *CreateStaffUseCase) validate(ctx context.Context, staff *entity.Staff) error {
	fmt.Println(staff)
	if staff.Username == "" {
		return fmt.Errorf("username is required")
	}
	if staff.PasswordHash == "" {
		return fmt.Errorf("password is required")
	}
	if staff.Email == "" {
		return fmt.Errorf("email is required")
	}
	if staff.HospitalID == 0 {
		return fmt.Errorf("hospital_id is required")
	}

	// Check if hospital exists
	exists, err := uc.HospitalRepo.Exists(ctx, staff.HospitalID)
	if err != nil {
		return fmt.Errorf("failed to validate hospital")
	}
	if !exists {
		return fmt.Errorf("invalid hospital_id")
	}

	return nil
}

func (uc *CreateStaffUseCase) createInDB(ctx context.Context, staff *entity.Staff) error {
	hashedPassword, err := uc.authService.EncryptPassword(staff.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}
	staff.PasswordHash = hashedPassword

	err = uc.StaffRepo.Create(ctx, staff)
	if err != nil {
		// Check for unique constraint violations
		errMsg := err.Error()
		if strings.Contains(errMsg, "staffs_username_key") || (strings.Contains(errMsg, "duplicate key") && strings.Contains(errMsg, "username")) {
			return fmt.Errorf("username already exists")
		}
		if strings.Contains(errMsg, "staffs_email_key") || (strings.Contains(errMsg, "duplicate key") && strings.Contains(errMsg, "email")) {
			return fmt.Errorf("email already exists")
		}
		return fmt.Errorf("failed to create staff")
	}

	return nil
}
