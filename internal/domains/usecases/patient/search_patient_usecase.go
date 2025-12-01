package patient

import (
	"context"
	"fmt"
	"time"

	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/model"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories"
)

type SearchPatientUseCase struct {
	PatientRepo   *repositories.PatientRepository
	PatientMapper *dataMapper.PatientMapper
}

func InitSearch(patientRepo *repositories.PatientRepository, patientMapper *dataMapper.PatientMapper) *SearchPatientUseCase {
	return &SearchPatientUseCase{PatientRepo: patientRepo, PatientMapper: patientMapper}
}

type SearchParams struct {
	HospitalID  string
	NationalID  string
	PassportID  string
	FirstName   string
	MiddleName  string
	LastName    string
	DateOfBirth string
	PhoneNumber string
	Email       string
	Limit       int
	Offset      int
}

func (uc *SearchPatientUseCase) Execute(ctx context.Context, params *SearchParams) ([]*entity.Patient, error) {
	patientRepo := uc.PatientRepo.
		ModelWithContext(ctx).
		Where("deleted_at IS NULL")

	limit := 10
	offset := 0

	if params == nil {
		params = &SearchParams{}
	}

	if params.Limit > 0 {
		limit = params.Limit
	}

	if params.Offset > 0 {
		offset = params.Offset
	}

	if params.NationalID != "" {
		patientRepo = patientRepo.Where("national_id = ?", params.NationalID)
	}

	if params.PassportID != "" {
		patientRepo = patientRepo.Where("passport_id = ?", params.PassportID)
	}

	if params.PhoneNumber != "" {
		patientRepo = patientRepo.Where("phone_number = ?", params.PhoneNumber)
	}

	if params.Email != "" {
		patientRepo = patientRepo.Where("email = ?", params.Email)
	}

	if params.FirstName != "" {
		like := "%" + params.FirstName + "%"
		patientRepo = patientRepo.Where(
			"(first_name_th ILIKE ? OR first_name_en ILIKE ?)",
			like, like,
		)
	}

	if params.LastName != "" {
		like := "%" + params.LastName + "%"
		patientRepo = patientRepo.Where(
			"(last_name_th ILIKE ? OR last_name_en ILIKE ?)",
			like, like,
		)
	}

	if params.MiddleName != "" {
		like := "%" + params.MiddleName + "%"
		patientRepo = patientRepo.Where(
			"(middle_name_th ILIKE ? OR middle_name_en ILIKE ?)",
			like, like,
		)
	}

	if params.DateOfBirth != "" {
		dob, err := time.Parse("2006-01-02", params.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("invalid date format (must be YYYY-MM-DD): %w", err)
		}
		patientRepo = patientRepo.Where("date_of_birth = ?", dob)
	}

	if params.HospitalID != "" {
		patientRepo = patientRepo.Where("hospital_id = ?", params.HospitalID)
	}

	var models []*model.PatientModel
	if err := patientRepo.Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, err
	}

	return uc.PatientMapper.ToEntityList(models), nil
}
