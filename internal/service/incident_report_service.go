package service

import (
	"time"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type IncidentReportService struct {
	repo *repository.IncidentReportRepository
}

func NewIncidentReportService(repo *repository.IncidentReportRepository) *IncidentReportService {
	return &IncidentReportService{repo: repo}
}

func (s *IncidentReportService) Create(permitID uint, description string, photo string, date *time.Time) error {
	report := &domain.IncidentReport{
		PermitID:    permitID,
		Description: description,
		Photo:       photo,
		Date:        date,
	}

	return s.repo.Create(report)
}