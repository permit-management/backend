package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type AuthMobileRepository interface {
	FindByPermitAndNIK(permitNumber, nik string) (*domain.Worker, error)
}

type authMobileRepository struct {
	db *gorm.DB
}

func NewAuthMobileRepository(db *gorm.DB) AuthMobileRepository {
	return &authMobileRepository{db}
}

func (r *authMobileRepository) FindByPermitAndNIK(permitNumber, nik string) (*domain.Worker, error) {
	var worker domain.Worker
	err := r.db.
		Joins("JOIN tbl_permit p ON p.id = tbl_worker.permit_id").
		Where("p.permit_number = ? AND tbl_worker.nik = ?", permitNumber, nik).
		First(&worker).Error
	if err != nil {
		return nil, err
	}
	return &worker, nil
}