package domain

import "time"

type DailyCheckIn struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	PermitID   uint      `json:"permit_id"`
	WorkerName string    `json:"worker_name"`
	Nik        string    `json:"nik"`
	PhotoURL   string    `json:"photo_url"`
	Status     string    `json:"status"` // contoh: CHECKIN, CHECKOUT
	Date       time.Time `json:"date"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// relasi
	Permit Permit `json:"permit" gorm:"foreignKey:PermitID"`
	Worker Worker `json:"worker" gorm:"foreignKey:Nik"`
}

func (DailyCheckIn) TableName() string {
	return "tbl_daily_check_in"
}
