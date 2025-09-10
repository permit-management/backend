package domain

import "time"

type WorkType struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:id"`
	WorkType  string    `json:"work_type" gorm:"column:work_type"`
	Approval1 uint      `json:"approval_1" gorm:"column:approval_1"`
	Approval2 uint      `json:"approval_2" gorm:"column:approval_2"`
	Approval3 uint      `json:"approval_3" gorm:"column:approval_3"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (WorkType) TableName() string {
	return "tbl_work_type"
}