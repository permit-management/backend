package service

import (
	"errors"
	"time"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type CheckInService interface {
	CreateCheckIn(permitID uint, workerName string, Nik string, photoURL string) (*domain.DailyCheckIn, error)
}

type checkInService struct {
	repo repository.CheckInRepository
}

func NewCheckInService(r repository.CheckInRepository) CheckInService {
	return &checkInService{repo: r}
}

func (s *checkInService) CreateCheckIn(permitID uint, workerName string, Nik string, photoURL string) (*domain.DailyCheckIn, error) {
	if permitID == 0 {
		return nil, errors.New("invalid permit")
	}


	checkIn := &domain.DailyCheckIn{
		PermitID:   permitID,
		WorkerName: workerName,
		Nik: 	  	Nik,
		PhotoURL:   photoURL,
		Status:     "CHECKIN",
		Date:       time.Now(),
	}

	if err := s.repo.Save(checkIn); err != nil {
		return nil, err
	}

	return checkIn, nil
}