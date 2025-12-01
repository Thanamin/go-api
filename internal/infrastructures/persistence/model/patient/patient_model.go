package patient

import "time"

type PatientModel struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	HospitalID   int        `gorm:"column:hospital_id;index;uniqueIndex:idx_patients_hospital_patient_hn" json:"hospital_id"`
	NationalID   string     `gorm:"column:national_id;index" json:"national_id"`
	PassportID   string     `gorm:"column:passport_id;index" json:"passport_id"`
	FirstNameTh  string     `gorm:"column:first_name_th" json:"first_name_th"`
	MiddleNameTh string     `gorm:"column:middle_name_th" json:"middle_name_th"`
	LastNameTh   string     `gorm:"column:last_name_th" json:"last_name_th"`
	FirstNameEn  string     `gorm:"column:first_name_en" json:"first_name_en"`
	MiddleNameEn string     `gorm:"column:middle_name_en" json:"middle_name_en"`
	LastNameEn   string     `gorm:"column:last_name_en" json:"last_name_en"`
	DateOfBirth  time.Time  `gorm:"column:date_of_birth" json:"date_of_birth"`
	PatientHN    string     `gorm:"column:patient_hn;uniqueIndex:idx_patients_hospital_patient_hn" json:"patient_hn"`
	PhoneNumber  string     `gorm:"column:phone_number" json:"phone_number"`
	Email        string     `gorm:"column:email" json:"email"`
	Gender       string     `gorm:"column:gender" json:"gender"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (PatientModel) TableName() string {
	return "patients"
}
