package domain

import "time"

type PermitApproval struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PermitID  uint      `json:"permit_id"`
	ApprovedBy uint     `json:"approved_by"`
	Status    string    `json:"status"` // Approved / Rejected / Pending
	Note      string    `json:"note"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Permit Permit `json:"permit" gorm:"foreignKey:PermitID"`
}


func (PermitApproval) TableName() string {
	return "tbl_permit_approval"
}