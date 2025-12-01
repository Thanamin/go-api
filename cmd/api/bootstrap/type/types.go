package types

import (
	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"

	patientUseCase "github.com/thanamin/go-api/internal/domains/usecases/patient"
	staffUseCase "github.com/thanamin/go-api/internal/domains/usecases/staff"

	"github.com/thanamin/go-api/internal/presenters/controllers"
)

// Mappers holds all data mapper instances
type Mappers struct {
	Hospital *dataMapper.HospitalMapper
	Staff    *dataMapper.StaffMapper
	Patient  *dataMapper.PatientMapper
}

// UseCases holds all use case instances
type UseCases struct {
	StaffCreate   *staffUseCase.CreateStaffUseCase
	StaffLogin    *staffUseCase.LoginStaffUseCase
	PatientSearch *patientUseCase.SearchPatientUseCase
}

// Controllers is an alias for the controllers package type
type Controllers = controllers.Controllers
