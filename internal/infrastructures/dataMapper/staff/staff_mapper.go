package staff

import (
	"github.com/thanamin/go-api/internal/domains/entity/staff"
	model "github.com/thanamin/go-api/internal/infrastructures/persistence/model/staff"
)

type StaffMapper struct{}

func NewStaffMapper() *StaffMapper {
	return &StaffMapper{}
}

func (m *StaffMapper) ToEntity(model *model.StaffModel) *staff.Staff {
	if model == nil {
		return nil
	}

	return &staff.Staff{
		ID:           model.ID,
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
		HospitalID:   model.HospitalID,
		FirstNameTh:  model.FirstNameTh,
		MiddleNameTh: model.MiddleNameTh,
		LastNameTh:   model.LastNameTh,
		FirstNameEn:  model.FirstNameEn,
		MiddleNameEn: model.MiddleNameEn,
		LastNameEn:   model.LastNameEn,
		Email:        model.Email,
		Position:     model.Position,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
	}
}

func (m *StaffMapper) ToModel(entity *staff.Staff) *model.StaffModel {
	if entity == nil {
		return nil
	}

	return &model.StaffModel{
		ID:           entity.ID,
		Username:     entity.Username,
		PasswordHash: entity.PasswordHash,
		HospitalID:   entity.HospitalID,
		FirstNameTh:  entity.FirstNameTh,
		MiddleNameTh: entity.MiddleNameTh,
		LastNameTh:   entity.LastNameTh,
		FirstNameEn:  entity.FirstNameEn,
		MiddleNameEn: entity.MiddleNameEn,
		LastNameEn:   entity.LastNameEn,
		Email:        entity.Email,
		Position:     entity.Position,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
	}
}

func (m *StaffMapper) ToEntityList(models []*model.StaffModel) []*staff.Staff {
	entities := make([]*staff.Staff, len(models))
	for i, model := range models {
		entities[i] = m.ToEntity(model)
	}
	return entities
}

func (m *StaffMapper) ToModelList(entities []*staff.Staff) []*model.StaffModel {
	models := make([]*model.StaffModel, len(entities))
	for i, entity := range entities {
		models[i] = m.ToModel(entity)
	}
	return models
}
