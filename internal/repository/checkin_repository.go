package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type CheckInRepository interface {
	Save(checkIn *domain.DailyCheckIn) error
}

type checkInRepository struct {
	db *gorm.DB
}

func NewCheckInRepository(db *gorm.DB) CheckInRepository {
	return &checkInRepository{db: db}
}

func (r *checkInRepository) Save(checkIn *domain.DailyCheckIn) error {
	return r.db.Create(checkIn).Error
}
