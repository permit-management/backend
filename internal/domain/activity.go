package domain

import "time"

type Activity struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PermitID    uint      `json:"permit_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Status      string    `json:"status"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Activity) TableName() string {
	return "tbl_activity"
}
