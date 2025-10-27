package domain

import "time"

type Worker struct {
	ID          uint      `json:"-" gorm:"primaryKey;autoIncrement"` // abaikan ID
	PermitID    uint      `json:"permit_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	NumberPhone string    `json:"phone" gorm:"column:phone_number"`
	NIK         string    `json:"nik"`

	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}



func (Worker) TableName() string {
    return "tbl_worker"
}