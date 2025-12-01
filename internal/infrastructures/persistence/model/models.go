package model

import (
	hospitalModel "github.com/thanamin/go-api/internal/infrastructures/persistence/model/hospital"
	patientModel "github.com/thanamin/go-api/internal/infrastructures/persistence/model/patient"
	staffModel "github.com/thanamin/go-api/internal/infrastructures/persistence/model/staff"
)

// Type aliases for models
type HospitalModel = hospitalModel.HospitalModel
type StaffModel = staffModel.StaffModel
type PatientModel = patientModel.PatientModel
