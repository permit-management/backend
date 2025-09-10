package repository 

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type WorkTypeRepository interface {
	Create(workType *domain.WorkType) error
	FindAll() ([]domain.WorkType, error)
	FindByID(id uint) (*domain.WorkType, error)
	Update(workType *domain.WorkType) error
	Delete(id uint) error
}

type workTypeRepository struct {
	db *gorm.DB
}

func NewWorkTypeRepository(db *gorm.DB) WorkTypeRepository {
	return &workTypeRepository{db}
}

func (r *workTypeRepository) Create(workType *domain.WorkType) error {
	return r.db.Create(workType).Error
}

func (r *workTypeRepository) FindAll() ([]domain.WorkType, error) {
	var workTypes []domain.WorkType
	err := r.db.Preload("ApprovalUser1").Preload("ApprovalUser2").Preload("ApprovalUser3").Find(&workTypes).Error
	return workTypes, err
}

func (r *workTypeRepository) FindByID(id uint) (*domain.WorkType, error) {
	var workType domain.WorkType
	err := r.db.Preload("ApprovalUser1").Preload("ApprovalUser2").Preload("ApprovalUser3").First(&workType, id).Error
	return &workType, err
}

func (r *workTypeRepository) Update(workType *domain.WorkType) error {
	return r.db.Model(&domain.WorkType{}).
	Where("id = ?", workType.ID).
	Omit("created_at").
	Updates(workType).Error
}

func (r *workTypeRepository) Delete(id uint) error {
	return r.db.Delete(&domain.WorkType{}, id).Error
}
