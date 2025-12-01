package bootstrap

import (
	bt "github.com/thanamin/go-api/cmd/api/bootstrap/type"
	patientController "github.com/thanamin/go-api/internal/presenters/controllers/patient"
	staffController "github.com/thanamin/go-api/internal/presenters/controllers/staff"
)

// InitControllers initializes and returns all controllers
func InitControllers(useCases *bt.UseCases, services *Services) *bt.Controllers {
	return &bt.Controllers{
		Staff:   staffController.NewStaffController(useCases.StaffCreate, useCases.StaffLogin, services.Auth),
		Patient: patientController.NewPatientController(useCases.PatientSearch, services.Auth),
	}
}
