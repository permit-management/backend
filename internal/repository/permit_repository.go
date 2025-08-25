package repository

import (
	"gorm.io/gorm"
	"github.com/permit-management/backend/internal/domain"
)

type PermitRepository interface {
	Create(permit *domain.Permit) error
	FindAll() ([]domain.Permit, error)
	FindByID(id uint) (domain.Permit, error)
	Update(permit *domain.Permit) error
	Delete(id uint) error
}

type permitRepository struct {
	db *gorm.DB
}

func NewPermitRepository(db *gorm.DB) PermitRepository {
	return &permitRepository{db}
}

func (r *permitRepository) Create(permit *domain.Permit) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&permit).Error; err != nil {
			return err
		}
		for i := range permit.Activities {
			permit.Activities[i].PermitID = permit.ID
			if err := tx.Create(&permit.Activities[i]).Error; err != nil {
				return err
			}
		}
		for i := range permit.Workers {
			permit.Workers[i].PermitID = permit.ID
			if err := tx.Create(&permit.Workers[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *permitRepository) FindAll() ([]domain.Permit, error) {
	var permits []domain.Permit
	err := r.db.Preload("Activities").Preload("Workers").Find(&permits).Error
	return permits, err
}

func (r *permitRepository) FindByID(id uint) (domain.Permit, error) {
	var permit domain.Permit
	err := r.db.Preload("Activities").Preload("Workers").First(&permit, id).Error
	return permit, err
}

func (r *permitRepository) Update(permit *domain.Permit) error {
	return r.db.Save(&permit).Error
}

func (r *permitRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Permit{}, id).Error
}
