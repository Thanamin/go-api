package bootstrap

import (
	"gorm.io/gorm"

	bt "github.com/thanamin/go-api/cmd/api/bootstrap/type"
	patientUseCase "github.com/thanamin/go-api/internal/domains/usecases/patient"
	staffUseCase "github.com/thanamin/go-api/internal/domains/usecases/staff"
	repositories "github.com/thanamin/go-api/internal/infrastructures/persistence/repositories"
)

// InitUseCases initializes and returns all use cases
func InitUseCases(db *gorm.DB, mappers *bt.Mappers, services *Services) *bt.UseCases {
	patientRepo := repositories.NewPatientRepository(db)
	staffRepo := repositories.NewStaffRepository(db)

	return &bt.UseCases{
		StaffCreate:   staffUseCase.InitCreate(services.Auth, staffRepo),
		StaffLogin:    staffUseCase.InitLogin(services.Auth, staffRepo),
		PatientSearch: patientUseCase.InitSearch(patientRepo, mappers.Patient),
	}
}
