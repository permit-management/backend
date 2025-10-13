package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type IncidentReportRepository struct {
	db *gorm.DB
}

func NewIncidentReportRepository(db *gorm.DB) *IncidentReportRepository {
	return &IncidentReportRepository{db: db}
}

func (r *IncidentReportRepository) Create(report *domain.IncidentReport) error {
	return r.db.Create(report).Error
}
