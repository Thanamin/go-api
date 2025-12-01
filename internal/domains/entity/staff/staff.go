package staff

import "time"

type Staff struct {
	ID           int        `json:"id"`
	Username     string     `json:"username"`
	PasswordHash string     `json:"-"`
	HospitalID   int        `json:"hospital_id"`
	FirstNameTh  string     `json:"first_name_th"`
	MiddleNameTh string     `json:"middle_name_th"`
	LastNameTh   string     `json:"last_name_th"`
	FirstNameEn  string     `json:"first_name_en"`
	MiddleNameEn string     `json:"middle_name_en"`
	LastNameEn   string     `json:"last_name_en"`
	Email        string     `json:"email"`
	Position     string     `json:"position"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
