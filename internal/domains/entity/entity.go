package entity

import (
	hospitalEntity "github.com/thanamin/go-api/internal/domains/entity/hospital"
	patientEntity "github.com/thanamin/go-api/internal/domains/entity/patient"
	staffEntity "github.com/thanamin/go-api/internal/domains/entity/staff"
)

// Type aliases for centralized entity access
type (
	Hospital = hospitalEntity.Hospital
	Staff    = staffEntity.Staff
	Patient  = patientEntity.Patient
)
