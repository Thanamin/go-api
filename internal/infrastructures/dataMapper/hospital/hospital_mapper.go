package hospital

import (
	"github.com/thanamin/go-api/internal/domains/entity/hospital"
	model "github.com/thanamin/go-api/internal/infrastructures/persistence/model/hospital"
)

type HospitalMapper struct{}

func NewHospitalMapper() *HospitalMapper {
	return &HospitalMapper{}
}

// ToEntity converts database model to domain entity
func (m *HospitalMapper) ToEntity(dbModel *model.HospitalModel) *hospital.Hospital {
	if dbModel == nil {
		return nil
	}

	return &hospital.Hospital{
		ID:        dbModel.ID,
		NameTh:    dbModel.NameTh,
		NameEn:    dbModel.NameEn,
		Email:     dbModel.Email,
		Tel:       dbModel.Tel,
		Address:   dbModel.Address,
		CreatedAt: dbModel.CreatedAt,
		UpdatedAt: dbModel.UpdatedAt,
		DeletedAt: dbModel.DeletedAt,
	}
}

// ToModel converts domain entity to database model
func (m *HospitalMapper) ToModel(entity *hospital.Hospital) *model.HospitalModel {
	if entity == nil {
		return nil
	}

	return &model.HospitalModel{
		ID:        entity.ID,
		NameTh:    entity.NameTh,
		NameEn:    entity.NameEn,
		Email:     entity.Email,
		Tel:       entity.Tel,
		Address:   entity.Address,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func (m *HospitalMapper) ToEntityList(models []*model.HospitalModel) []*hospital.Hospital {
	entities := make([]*hospital.Hospital, len(models))
	for i, model := range models {
		entities[i] = m.ToEntity(model)
	}
	return entities
}

func (m *HospitalMapper) ToModelList(entities []*hospital.Hospital) []*model.HospitalModel {
	models := make([]*model.HospitalModel, len(entities))
	for i, entity := range entities {
		models[i] = m.ToModel(entity)
	}
	return models
}
