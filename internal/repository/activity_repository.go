package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	FindByID(id uint) (*domain.Activity, error)
	FindByPermitID(permitID uint) ([]domain.Activity, error)
	Update(activity *domain.Activity) error
	UpdateStatus(id uint, status string) error
	FindByPermitAndWorker(permitID uint, nik string) ([]domain.Activity, error) // << baru
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) FindByID(id uint) (*domain.Activity, error) {
	var activity domain.Activity
	err := r.db.First(&activity, id).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *activityRepository) FindByPermitID(permitID uint) ([]domain.Activity, error) {
	var activities []domain.Activity
	err := r.db.Where("permit_id = ?", permitID).Find(&activities).Error
	return activities, err
}

func (r *activityRepository) Update(activity *domain.Activity) error {
	return r.db.Save(activity).Error
}

func (r *activityRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&domain.Activity{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// --- NEW ---
// Untuk sekarang kita filter by permit_id saja (nik tidak dipakai di tbl_activity).
func (r *activityRepository) FindByPermitAndWorker(permitID uint, nik string) ([]domain.Activity, error) {
	var activities []domain.Activity
	err := r.db.Where("permit_id = ?", permitID).Find(&activities).Error
	return activities, err
}