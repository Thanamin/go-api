package patient

import (
	"context"

	"gorm.io/gorm"

	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/model"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/base"
)

type PatientRepository struct {
	*base.BaseSQLRepository
	Mapper *dataMapper.PatientMapper
}

func (r *PatientRepository) DBWithContext(ctx context.Context) *gorm.DB {
	return r.BaseSQLRepository.DBWithContext(ctx)
}

func (r *PatientRepository) ModelWithContext(ctx context.Context) *gorm.DB {
	return r.DBWithContext(ctx).Model(&model.PatientModel{})
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	mapper := dataMapper.NewPatientMapper()
	return &PatientRepository{
		BaseSQLRepository: base.NewBaseSQLRepository(db.Model(&model.PatientModel{}), mapper),
		Mapper:            mapper,
	}
}
