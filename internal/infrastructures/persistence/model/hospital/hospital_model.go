package hospital

import (
	"time"
)

type HospitalModel struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	NameTh    string     `gorm:"column:name_th" json:"name_th"`
	NameEn    string     `gorm:"column:name_en" json:"name_en"`
	Email     string     `gorm:"column:email" json:"email"`
	Tel       string     `gorm:"column:tel" json:"tel"`
	Address   string     `gorm:"column:address" json:"address"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (HospitalModel) TableName() string {
	return "hospitals"
}
