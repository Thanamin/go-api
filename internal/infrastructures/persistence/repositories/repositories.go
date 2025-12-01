package repositories

import (
	hospitalRepo "github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/hospital"
	patientRepo "github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/patient"
	staffRepo "github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/staff"
)

type (
	HospitalRepository = hospitalRepo.HospitalRepository
	StaffRepository    = staffRepo.StaffRepository
	PatientRepository  = patientRepo.PatientRepository
)

var (
	NewHospitalRepository = hospitalRepo.NewHospitalRepository
	NewStaffRepository    = staffRepo.NewStaffRepository
	NewPatientRepository  = patientRepo.NewPatientRepository
)
