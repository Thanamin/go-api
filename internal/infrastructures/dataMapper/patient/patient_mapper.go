package patient

import (
	"github.com/thanamin/go-api/internal/domains/entity/patient"
	model "github.com/thanamin/go-api/internal/infrastructures/persistence/model/patient"
)

type PatientMapper struct{}

func NewPatientMapper() *PatientMapper {
	return &PatientMapper{}
}

func (m *PatientMapper) ToEntity(model *model.PatientModel) *patient.Patient {
	if model == nil {
		return nil
	}

	return &patient.Patient{
		HospitalID:   model.HospitalID,
		ID:           model.ID,
		NationalID:   model.NationalID,
		PassportID:   model.PassportID,
		FirstNameTh:  model.FirstNameTh,
		MiddleNameTh: model.MiddleNameTh,
		LastNameTh:   model.LastNameTh,
		FirstNameEn:  model.FirstNameEn,
		MiddleNameEn: model.MiddleNameEn,
		LastNameEn:   model.LastNameEn,
		DateOfBirth:  model.DateOfBirth,
		PatientHN:    model.PatientHN,
		PhoneNumber:  model.PhoneNumber,
		Email:        model.Email,
		Gender:       model.Gender,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
	}
}

func (m *PatientMapper) ToModel(entity *patient.Patient) *model.PatientModel {
	if entity == nil {
		return nil
	}

	return &model.PatientModel{
		HospitalID:   entity.HospitalID,
		ID:           entity.ID,
		NationalID:   entity.NationalID,
		PassportID:   entity.PassportID,
		FirstNameTh:  entity.FirstNameTh,
		MiddleNameTh: entity.MiddleNameTh,
		LastNameTh:   entity.LastNameTh,
		FirstNameEn:  entity.FirstNameEn,
		MiddleNameEn: entity.MiddleNameEn,
		LastNameEn:   entity.LastNameEn,
		DateOfBirth:  entity.DateOfBirth,
		PatientHN:    entity.PatientHN,
		PhoneNumber:  entity.PhoneNumber,
		Email:        entity.Email,
		Gender:       entity.Gender,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
	}
}

func (m *PatientMapper) ToEntityList(models []*model.PatientModel) []*patient.Patient {
	entities := make([]*patient.Patient, len(models))
	for i, model := range models {
		entities[i] = m.ToEntity(model)
	}
	return entities
}

func (m *PatientMapper) ToModelList(entities []*patient.Patient) []*model.PatientModel {
	models := make([]*model.PatientModel, len(entities))
	for i, entity := range entities {
		models[i] = m.ToModel(entity)
	}
	return models
}
