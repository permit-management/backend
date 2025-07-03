package repository

import (
	"context"
	"errors"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/pkg/app"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id uint) (*domain.UserModel, error)
	ListUsers(pager *app.Pager) ([]*domain.UserModel, error)
	CountUsers(pager *app.Pager) (int64, error)
	CreateUser(username, email, password string) (*domain.UserModel, error)
	UpdateUser(id uint, username, email string) (*domain.UserModel, error)
	DeleteUser(id uint) error
}

type userRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewUserRepository(ctx context.Context, db *gorm.DB) UserRepository {
	return &userRepository{
		db:  db.WithContext(ctx),
		ctx: ctx,
	}
}

func (r *userRepository) GetUser(id uint) (*domain.UserModel, error) {
	var user domain.UserModel
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ListUsers(pager *app.Pager) ([]*domain.UserModel, error) {
	var users []*domain.UserModel
	db := r.db.Model(&domain.UserModel{})
	pager.SearchCriteria(nil)
	if err := db.Offset(pager.Offset()).Limit(pager.PageSize).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CountUsers(pager *app.Pager) (int64, error) {
	var count int64
	db := r.db.Model(&domain.UserModel{})
	pager.SearchCriteria(nil)
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) CreateUser(username, email, password string) (*domain.UserModel, error) {
	user := domain.UserModel{
		Username: username,
		Email:    email,
		Password: password,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(id uint, username, email string) (*domain.UserModel, error) {
	values := map[string]interface{}{
		"username": username,
		"email":    email,
	}
	result := r.db.Model(&domain.UserModel{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	var user domain.UserModel
	r.db.First(&user, id)
	return &user, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&domain.UserModel{}, id).Error
}
