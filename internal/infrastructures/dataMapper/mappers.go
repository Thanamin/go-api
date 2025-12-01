package dataMapper

import (
	hospitalMapper "github.com/thanamin/go-api/internal/infrastructures/dataMapper/hospital"
	patientMapper "github.com/thanamin/go-api/internal/infrastructures/dataMapper/patient"
	staffMapper "github.com/thanamin/go-api/internal/infrastructures/dataMapper/staff"
)

// Type aliases for centralized mapper access
type (
	HospitalMapper = hospitalMapper.HospitalMapper
	StaffMapper    = staffMapper.StaffMapper
	PatientMapper  = patientMapper.PatientMapper
)

// Constructor aliases for centralized mapper creation
var (
	NewHospitalMapper = hospitalMapper.NewHospitalMapper
	NewStaffMapper    = staffMapper.NewStaffMapper
	NewPatientMapper  = patientMapper.NewPatientMapper
)
