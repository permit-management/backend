package service

import (
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
	"time"
)

type UserService interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *domain.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// TODO: hash password before saving
	return s.repo.Create(user)
}

func (s *userService) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetByID(id uint) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Update(user *domain.User) error {
	user.UpdatedAt = time.Now()
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
