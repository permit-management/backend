package domain

import "time"

type Departement struct {
	ID              uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	DepartementID   string     `json:"departements_id" gorm:"column:departements_id;size:50;not null;unique"`
	DepartementName string     `json:"departements_name" gorm:"column:departements_name;size:100;not null"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Departement) TableName() string {
	return "tbl_departements"
}
