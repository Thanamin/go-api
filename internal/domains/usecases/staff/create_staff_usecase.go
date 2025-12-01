package staff

import (
	"context"
	"fmt"

	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/domains/service/auth"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories"
)

type CreateStaffUseCase struct {
	StaffRepo   *repositories.StaffRepository
	authService *auth.AuthService
}

func InitCreate(authService *auth.AuthService, staffRepo *repositories.StaffRepository) *CreateStaffUseCase {
	return &CreateStaffUseCase{
		authService: authService,
		StaffRepo:   staffRepo,
	}
}

func (uc *CreateStaffUseCase) Execute(ctx context.Context, staff *entity.Staff) error {
	if err := uc.validate(staff); err != nil {
		return err
	}

	return uc.createInDB(ctx, staff)
}

func (uc *CreateStaffUseCase) validate(staff *entity.Staff) error {
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
	return nil
}

func (uc *CreateStaffUseCase) createInDB(ctx context.Context, staff *entity.Staff) error {
	hashedPassword, err := uc.authService.EncryptPassword(staff.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	staff.PasswordHash = hashedPassword

	return uc.StaffRepo.Create(ctx, staff)
}
