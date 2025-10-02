package domain

import "time"

type DailyWorkCheck struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	PermitID    uint      `json:"permit_id"`
	ActivityID  uint      `json:"activity_id"`
	Date        time.Time `json:"date"`
	Nik         string    `json:"nik"`
	Status      string    `json:"status"`
	Description string    `json:"description"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Permit   Permit   `json:"permit" gorm:"foreignKey:PermitID"`
	Activity Activity `json:"activity" gorm:"foreignKey:ActivityID"`
	NIK      string   `json:"nik" gorm:"foreignKey:Nik;references:Nik"` // ini aku tambahin sendiri, karna menurutku relasi ini agar lebih jelas nik ini milik siapa / relasi ke worker
}

func (DailyWorkCheck) TableName() string {
	return "tbl_daily_work_check"
}
