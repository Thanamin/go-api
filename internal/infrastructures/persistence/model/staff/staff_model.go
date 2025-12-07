package staff

import (
	"time"
)

type StaffModel struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string     `gorm:"column:username;unique;not null" json:"username"`
	PasswordHash string     `gorm:"column:password_hash;not null" json:"-"`
	HospitalID   int        `gorm:"column:hospital_id;index" json:"hospital_id"`
	FirstNameTh  string     `gorm:"column:first_name_th" json:"first_name_th"`
	MiddleNameTh string     `gorm:"column:middle_name_th" json:"middle_name_th"`
	LastNameTh   string     `gorm:"column:last_name_th" json:"last_name_th"`
	FirstNameEn  string     `gorm:"column:first_name_en" json:"first_name_en"`
	MiddleNameEn string     `gorm:"column:middle_name_en" json:"middle_name_en"`
	LastNameEn   string     `gorm:"column:last_name_en" json:"last_name_en"`
	Email        string     `gorm:"column:email;unique" json:"email"`
	Position     string     `gorm:"column:position" json:"position"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (StaffModel) TableName() string {
	return "staffs"
}
