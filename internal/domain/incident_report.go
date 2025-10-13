package domain

import "time"

type IncidentReport struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	PermitID    uint       `json:"permit_id"`
	Description string     `json:"description"`
	Photo       string     `json:"photo"`
	Date        *time.Time `json:"date" gorm:"type:datetime;default:null"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (IncidentReport) TableName() string {
	return "tbl_incident_report"
}
