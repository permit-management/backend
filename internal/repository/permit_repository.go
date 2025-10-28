package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PermitRepository interface {
	Create(permit *domain.Permit) error
	FindAll() ([]domain.Permit, error)
	FindByID(id uint) (domain.Permit, error)
	Update(permit *domain.Permit) error
	Delete(id uint) error

	// tambahan
	UpdateStatus(permitID uint, status string) error
}

type permitRepository struct {
	db *gorm.DB
}

func NewPermitRepository(db *gorm.DB) PermitRepository {
	return &permitRepository{db}
}

// hindari duplikasi data relasi
func (r *permitRepository) Create(permit *domain.Permit) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// insert permit tanpa auto-insert relasi
		if err := tx.Omit(clause.Associations).Create(permit).Error; err != nil {
			return err
		}

		// insert activities (manual, tapi hanya sekali)
		for i := range permit.Activities {
			permit.Activities[i].PermitID = permit.ID
			if err := tx.Create(&permit.Activities[i]).Error; err != nil {
				return err
			}
		}

		// insert workers (manual, tapi hanya sekali)
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
	err := r.db.Preload("Activities").Preload("Workers").Preload("WorkType").Find(&permits).Error
	return permits, err
}

func (r *permitRepository) FindByID(id uint) (domain.Permit, error) {
	var permit domain.Permit
	err := r.db.Preload("Activities").Preload("Workers").Preload("WorkType").Preload("CheckIns").Preload("Incidents").First(&permit, id).Error
	return permit, err
}

func (r *permitRepository) Update(permit *domain.Permit) error {
	return r.db.Save(permit).Error
}

func (r *permitRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Permit{}, id).Error
}

// tambahan: update status permit
func (r *permitRepository) UpdateStatus(permitID uint, status string) error {
	return r.db.Model(&domain.Permit{}).
		Where("id = ?", permitID).
		Update("status", status).Error
}
