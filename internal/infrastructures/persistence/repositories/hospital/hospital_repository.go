package hospital

import (
	"context"

	"gorm.io/gorm"

	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/model"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/base"
)

type HospitalRepository struct {
	*base.BaseSQLRepository
	Mapper *dataMapper.HospitalMapper
}

func (r *HospitalRepository) DBWithContext(ctx context.Context) *gorm.DB {
	return r.BaseSQLRepository.DBWithContext(ctx)
}

func (r *HospitalRepository) ModelWithContext(ctx context.Context) *gorm.DB {
	return r.DBWithContext(ctx).Model(&model.HospitalModel{})
}

func NewHospitalRepository(db *gorm.DB) *HospitalRepository {
	mapper := dataMapper.NewHospitalMapper()
	return &HospitalRepository{
		BaseSQLRepository: base.NewBaseSQLRepository(db.Model(&model.HospitalModel{}), mapper),
		Mapper:            mapper,
	}
}
