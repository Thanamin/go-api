package staff

import (
	"context"
	"fmt"

	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/domains/service/auth"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories"
)

type LoginStaffUseCase struct {
	StaffRepo   *repositories.StaffRepository
	authService *auth.AuthService
}

func InitLogin(authService *auth.AuthService, staffRepo *repositories.StaffRepository) *LoginStaffUseCase {
	return &LoginStaffUseCase{
		authService: authService,
		StaffRepo:   staffRepo,
	}
}

func (uc *LoginStaffUseCase) Execute(ctx context.Context, email, password string) (*entity.Staff, error) {
	staff, err := uc.StaffRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := uc.verifyPassword(staff, password); err != nil {
		return nil, err
	}

	return staff, nil
}

func (uc *LoginStaffUseCase) verifyPassword(staff *entity.Staff, password string) error {
	if err := uc.authService.DecryptPassword(password, staff.PasswordHash); err != nil {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}
