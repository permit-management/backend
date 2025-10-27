package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type DepartementRepository interface {
	Create(dept *domain.Departement) error
	GetByID(id uint) (*domain.Departement, error)
	GetAll() ([]domain.Departement, error)
	Update(dept *domain.Departement) error
	UpdateWithoutCreatedAt(dept *domain.Departement) error
	Delete(id uint) error
}

type departementRepo struct {
	db *gorm.DB
}

func NewDepartementRepository(db *gorm.DB) DepartementRepository {
	return &departementRepo{db: db}
}

func (r *departementRepo) Create(dept *domain.Departement) error {
	return r.db.Create(dept).Error
}

func (r *departementRepo) GetByID(id uint) (*domain.Departement, error) {
	var dept domain.Departement
	err := r.db.First(&dept, id).Error
	return &dept, err
}

func (r *departementRepo) GetAll() ([]domain.Departement, error) {
	var depts []domain.Departement
	err := r.db.Find(&depts).Error
	return depts, err
}

func (r *departementRepo) Update(dept *domain.Departement) error {
	return r.db.Save(dept).Error
}

func (r *departementRepo) UpdateWithoutCreatedAt(dept *domain.Departement) error {
	return r.db.Model(&domain.Departement{}).
		Where("id = ?", dept.ID).
		Select("departements_id", "departements_name", "updated_at").
		Updates(dept).Error
}

func (r *departementRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Departement{}, id).Error
}
