package domain

import "time"

type Worker struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	PermitID uint   `json:"permit_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	NIK      string `json:"nik"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Worker) TableName() string {
	return "tbl_worker"
}
