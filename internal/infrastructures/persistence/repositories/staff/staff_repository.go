package staff

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/thanamin/go-api/internal/domains/entity"
	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/model"
	"github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/base"
)

type StaffRepository struct {
	*base.BaseSQLRepository
	Mapper *dataMapper.StaffMapper
}

func (r *StaffRepository) DBWithContext(ctx context.Context) *gorm.DB {
	return r.BaseSQLRepository.DBWithContext(ctx)
}

func (r *StaffRepository) ModelWithContext(ctx context.Context) *gorm.DB {
	return r.DBWithContext(ctx).Model(&model.StaffModel{})
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	mapper := dataMapper.NewStaffMapper()
	return &StaffRepository{
		BaseSQLRepository: base.NewBaseSQLRepository(db.Model(&model.StaffModel{}), mapper),
		Mapper:            mapper,
	}
}

func (r *StaffRepository) Create(ctx context.Context, staff *entity.Staff) error {
	m := r.Mapper.ToModel(staff)
	if err := r.DBWithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	*staff = *r.Mapper.ToEntity(m)
	return nil
}

func (r *StaffRepository) FindByEmail(ctx context.Context, email string) (*entity.Staff, error) {
	var model model.StaffModel
	err := r.DBWithContext(ctx).
		Where("email = ?", email).
		Where("deleted_at IS NULL").
		First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("staff not found with email: %s", email)
		}
		return nil, fmt.Errorf("failed to find staff: %w", err)
	}

	return r.Mapper.ToEntity(&model), nil
}
