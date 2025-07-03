package service

import (
	"context"

	"github.com/permit-management/backend/internal/constants"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/pkg/app"
	"github.com/permit-management/backend/pkg/errcode"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"id" binding:"required"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService struct {
	ctx  context.Context
	db   *gorm.DB
	repo repository.UserRepository
}

func NewUserService(ctx context.Context, db *gorm.DB) *UserService {
	return &UserService{
		ctx:  ctx,
		db:   db,
		repo: repository.NewUserRepository(ctx, db),
	}
}

func (s *UserService) GetUser(param *constants.IDRequest) (*domain.UserModel, *errcode.Error) {
	user, err := s.repo.GetUser(param.ID)
	if err != nil {
		return nil, errcode.BadRequest.WithDetails(err.Error())
	}
	return user, nil
}

func (s *UserService) ListUsers(pager *app.Pager) ([]*domain.UserModel, int, *errcode.Error) {
	cnt, err := s.repo.CountUsers(pager)
	if err != nil {
		return nil, 0, errcode.BadRequest.WithDetails(err.Error())
	}
	users, err := s.repo.ListUsers(pager)
	if err != nil {
		return nil, 0, errcode.BadRequest.WithDetails(err.Error())
	}
	return users, int(cnt), nil
}

func (s *UserService) CreateUser(param *CreateUserRequest) (*domain.UserModel, *errcode.Error) {
	user, err := s.repo.CreateUser(param.Username, param.Email, param.Password)
	if err != nil {
		return nil, errcode.BadRequest.WithDetails(err.Error())
	}
	return user, nil
}

func (s *UserService) UpdateUser(param *UpdateUserRequest) (*domain.UserModel, *errcode.Error) {
	user, err := s.repo.UpdateUser(param.ID, param.Username, param.Email)
	if err != nil {
		return nil, errcode.BadRequest.WithDetails(err.Error())
	}
	return user, nil
}

func (s *UserService) DeleteUser(param *constants.IDRequest) *errcode.Error {
	if err := s.repo.DeleteUser(param.ID); err != nil {
		return errcode.BadRequest.WithDetails(err.Error())
	}
	return nil
}