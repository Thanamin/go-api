package patient

import "time"

type Patient struct {
	ID           int        `json:"id"`
	HospitalID   int        `json:"hospital_id"`
	NationalID   string     `json:"national_id"`
	PassportID   string     `json:"passport_id"`
	FirstNameTh  string     `json:"first_name_th"`
	MiddleNameTh string     `json:"middle_name_th"`
	LastNameTh   string     `json:"last_name_th"`
	FirstNameEn  string     `json:"first_name_en"`
	MiddleNameEn string     `json:"middle_name_en"`
	LastNameEn   string     `json:"last_name_en"`
	DateOfBirth  time.Time  `json:"date_of_birth"`
	PatientHN    string     `json:"patient_hn"`
	PhoneNumber  string     `json:"phone_number"`
	Email        string     `json:"email"`
	Gender       string     `json:"gender"`
	StaffID      int        `json:"staff_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
