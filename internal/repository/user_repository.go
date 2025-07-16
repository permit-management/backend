package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error

	FindByEmail(email string) (*domain.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepo) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepo) Update(user *domain.User) error {
	// Update seluruh record dari user yang dikirim
	return r.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepo) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
