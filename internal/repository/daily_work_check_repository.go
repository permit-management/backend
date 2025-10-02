package repository

import (
	"github.com/permit-management/backend/internal/domain"

	"gorm.io/gorm"
)

type DailyWorkCheckRepository interface {
	Create(dwc *domain.DailyWorkCheck) error
}

type dailyWorkCheckRepository struct {
	db *gorm.DB
}

func NewDailyWorkCheckRepository(db *gorm.DB) DailyWorkCheckRepository {
	return &dailyWorkCheckRepository{db}
}

func (r *dailyWorkCheckRepository) Create(dwc *domain.DailyWorkCheck) error {
	return r.db.Create(dwc).Error
}
