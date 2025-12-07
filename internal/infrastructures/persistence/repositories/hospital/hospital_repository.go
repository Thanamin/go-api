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

// Exists checks if a hospital exists by ID
func (r *HospitalRepository) Exists(ctx context.Context, id int) (bool, error) {
	var count int64
	err := r.ModelWithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
