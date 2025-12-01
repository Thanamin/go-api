package bootstrap

import (
	bt "github.com/thanamin/go-api/cmd/api/bootstrap/type"
	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"
)

// InitMappers initializes and returns all data mappers
func InitMappers() *bt.Mappers {
	return &bt.Mappers{
		Hospital: dataMapper.NewHospitalMapper(),
		Staff:    dataMapper.NewStaffMapper(),
		Patient:  dataMapper.NewPatientMapper(),
	}
}
