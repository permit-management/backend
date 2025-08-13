package service

import (
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
)

type RoleService interface {
	Create(role *domain.Role) error
	GetAll() ([]domain.Role, error)
	GetByID(id uint) (*domain.Role, error)
	Update(role *domain.Role) error
	Delete(id uint) error
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) Create(role *domain.Role) error {
	err := s.repo.Create(role)
	if err != nil {
		return err
	}

	for i := range role.Permissions {
		role.Permissions[i].RoleID = role.ID
		err = s.repo.CreateRolePermission(&role.Permissions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *roleService) GetAll() ([]domain.Role, error) {
	return s.repo.GetAll()
}

func (s *roleService) GetByID(id uint) (*domain.Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) Update(role *domain.Role) error {
	// Ambil role lama dari DB
	existingRole, err := s.repo.GetByID(role.ID)
	if err != nil {
		return err
	}

	// Update field dasar
	existingRole.RoleID = role.RoleID
	existingRole.RoleName = role.RoleName

	// Hapus permission lama
	err = s.repo.DeleteRolePermissionsByRoleID(role.ID)
	if err != nil {
		return err
	}

	// Simpan permission baru
	for i := range role.Permissions {
		role.Permissions[i].RoleID = role.ID
		err = s.repo.CreateRolePermission(&role.Permissions[i])
		if err != nil {
			return err
		}
	}

	// Update role utama
	return s.repo.Update(existingRole)
}

func (s *roleService) Delete(id uint) error {
	return s.repo.Delete(id)
}