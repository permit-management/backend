package service

import (
	"time"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type DailyWorkCheckService interface {
	MarkDone(input DailyWorkCheckInput) error
	GetActivitiesByWorker(permitID uint, nik string) ([]domain.Activity, error) // << baru
}

type dailyWorkCheckService struct {
	repo         repository.DailyWorkCheckRepository
	activityRepo repository.ActivityRepository
}

func NewDailyWorkCheckService(repo repository.DailyWorkCheckRepository, activityRepo repository.ActivityRepository) DailyWorkCheckService {
	return &dailyWorkCheckService{repo, activityRepo}
}

type DailyWorkCheckInput struct {
	PermitID    uint   `json:"permit_id"`
	ActivityID  uint   `json:"activity_id"`
	Nik         string `json:"nik"`
	Description string `json:"description"`
}

func (s *dailyWorkCheckService) MarkDone(input DailyWorkCheckInput) error {
	dwc := domain.DailyWorkCheck{
		PermitID:    input.PermitID,
		ActivityID:  input.ActivityID,
		Nik:         input.Nik,
		Description: input.Description,
		Date:        time.Now(),
		Status:      "Done",
	}

	if err := s.repo.Create(&dwc); err != nil {
		return err
	}

	if err := s.activityRepo.UpdateStatus(input.ActivityID, "Done"); err != nil {
		return err
	}

	return nil
}

// --- NEW ---
func (s *dailyWorkCheckService) GetActivitiesByWorker(permitID uint, nik string) ([]domain.Activity, error) {
	return s.activityRepo.FindByPermitAndWorker(permitID, nik)
}
