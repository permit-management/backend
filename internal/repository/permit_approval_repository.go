package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type PermitApprovalRepository interface {
	Create(approval *domain.PermitApproval) error
	FindByPermitID(permitID uint) ([]domain.PermitApproval, error)
}

type permitApprovalRepository struct {
	db *gorm.DB
}

func NewPermitApprovalRepository(db *gorm.DB) PermitApprovalRepository {
	return &permitApprovalRepository{db}
}

func (r *permitApprovalRepository) Create(approval *domain.PermitApproval) error {
	return r.db.Create(approval).Error
}

func (r *permitApprovalRepository) FindByPermitID(permitID uint) ([]domain.PermitApproval, error) {
	var approvals []domain.PermitApproval
	err := r.db.Where("permit_id = ?", permitID).Find(&approvals).Error
	return approvals, err
}
