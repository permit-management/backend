package domain

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        string    `json:"user_id" gorm:"unique;not null"`
	Name          string    `json:"name" gorm:"not null"`
	NumberPhone   string    `json:"number_phone" gorm:"not null"`
	Email         string    `json:"email" gorm:"not null;unique"`
	DepartementID string    `json:"departement_id" gorm:"column:departements_id"`
	RoleID        string    `json:"role_id" gorm:"column:role_id"`
	Password      string    `json:"password" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "tbl_users"
}
