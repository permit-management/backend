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
	return r.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"user_id":        user.UserID,
		"name":           user.Name,
		"number_phone":   user.NumberPhone,
		"email":          user.Email,
		"departements_id": user.DepartementID,
		"role_id":        user.RoleID,
		"updated_at":     user.UpdatedAt,
	}).Error
}

func (r *userRepo) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
