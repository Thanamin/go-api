package hospital

import "time"

type Hospital struct {
	ID          int        `json:"id"`
	NameTh      string     `json:"name_th"`
	NameEn      string     `json:"name_en"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
