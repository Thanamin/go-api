package base

import (
	"context"

	"github.com/thanamin/go-api/internal/infrastructures/dataMapper/base"
	"gorm.io/gorm"
)

type BaseSQLRepository struct {
	DB         *gorm.DB
	DataMapper base.IDataMapper
}

func NewBaseSQLRepository(db *gorm.DB, mapper base.IDataMapper) *BaseSQLRepository {
	return &BaseSQLRepository{
		DB:         db,
		DataMapper: mapper,
	}
}

func (r *BaseSQLRepository) DBWithContext(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}
