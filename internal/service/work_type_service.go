package service

import (
	"errors"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type WorkTypeService interface {
	Create(workType *domain.WorkType) error
	GetAll() ([]domain.WorkType, error)
	GetByID(id uint) (*domain.WorkType, error)
	Update(workType *domain.WorkType) error
	Delete(id uint) error
}

type workTypeService struct {
	repo repository.WorkTypeRepository
}


func NewWorkTypeService(repo repository.WorkTypeRepository) WorkTypeService {
	return &workTypeService{repo: repo}
}

func (s *workTypeService) Create(workType *domain.WorkType) error {
	if workType.WorkType == "" {
		return errors.New("work type is required")
	}
	if workType.Approval1 == workType.Approval2 || workType.Approval1 == workType.Approval3 || workType.Approval2 == workType.Approval3 {
		return errors.New("approval levels must be unique")
	}
	return s.repo.Create(workType)
}

func (s *workTypeService) GetAll() ([]domain.WorkType, error) {
	return s.repo.FindAll()
}

func (s *workTypeService) GetByID(id uint) (*domain.WorkType, error) {
	if id == 0 {
		return &domain.WorkType{}, errors.New("invalid ID")
	}
	return s.repo.FindByID(id)
}

func (s *workTypeService) Update(workType *domain.WorkType) error {
	if workType.ID == 0 {
		return errors.New("id is required for update")
	}
	if workType.Approval1 == workType.Approval2 || workType.Approval1 == workType.Approval3 || workType.Approval2 == workType.Approval3 {
		return errors.New("approval must be unique")
	}
	return s.repo.Update(workType)
}

func (s *workTypeService) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid ID")
	}
	return s.repo.Delete(id)
}