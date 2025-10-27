package domain

import "time"

type Permit struct {
    ID             uint       `json:"id" gorm:"primaryKey"`
    PermitNumber   string     `json:"permit_number" gorm:"unique;not null"`
    WorkName       string     `json:"work_name"`
    WorkTypeID     uint       `json:"work_type_id"`
    WorkType       WorkType   `json:"work_type" gorm:"foreignKey:WorkTypeID"`
    WorkingStart   time.Time  `json:"working_start" gorm:"column:working_start_datetime"`
    WorkingEnd     time.Time  `json:"working_end" gorm:"column:working_end_datetime"`
    WorkingArea    string     `json:"working_area" gorm:"column:working_area"`
    Risk           string     `json:"risk"`
    Status         string     `json:"status"`
    SubmitDate     time.Time  `json:"submit_date"`
    JsaText        string     `json:"jsa_text" gorm:"column:jsa_text"`
    CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime;column:created_at"`
    UpdatedAt      time.Time  `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`

    // relasi
    Activities []Activity `json:"activities" gorm:"foreignKey:PermitID"`
    Workers    []Worker   `json:"workers" gorm:"foreignKey:PermitID"`
}


func (Permit) TableName() string {
	return "tbl_permit"
}
