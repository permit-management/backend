package service

import (
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type DepartementService interface {
	Create(dept *domain.Departement) error
	GetAll() ([]domain.Departement, error)
	GetByID(id uint) (*domain.Departement, error)
	Update(dept *domain.Departement) error
	UpdateWithoutCreatedAt(dept *domain.Departement) error
	Delete(id uint) error
}

type departementService struct {
	repo repository.DepartementRepository
}

func NewDepartementService(r repository.DepartementRepository) DepartementService {
	return &departementService{repo: r}
}

func (s *departementService) Create(dept *domain.Departement) error {
	return s.repo.Create(dept)
}

func (s *departementService) GetAll() ([]domain.Departement, error) {
	return s.repo.GetAll()
}

func (s *departementService) GetByID(id uint) (*domain.Departement, error) {
	return s.repo.GetByID(id)
}

func (s *departementService) Update(dept *domain.Departement) error {
	return s.repo.Update(dept)
}

func (s *departementService) UpdateWithoutCreatedAt(dept *domain.Departement) error {
	return s.repo.UpdateWithoutCreatedAt(dept)
}

func (s *departementService) Delete(id uint) error {
	return s.repo.Delete(id)
}
