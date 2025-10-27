package repository

import (
	"github.com/permit-management/backend/internal/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *domain.Role) error
	GetAll() ([]domain.Role, error)
	GetByID(id uint) (*domain.Role, error)
	Update(role *domain.Role) error
	Delete(id uint) error

	CreateRolePermission(rp *domain.RolePermission) error
	DeleteRolePermissionsByRoleID(roleID uint) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetAll() ([]domain.Role, error) {
	var roles []domain.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) GetByID(id uint) (*domain.Role, error) {
	var role domain.Role

	// Ambil data role berdasarkan ID
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, err
	}

	// Ambil data permissions dari tbl_role_permissions berdasarkan role_id (string)
	var permissions []domain.RolePermission
	if err := r.db.Where("role_id = ?", role.RoleID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	role.Permissions = permissions

	return &role, nil
}

func (r *roleRepository) Update(role *domain.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Role{}, id).Error
}

func (r *roleRepository) CreateRolePermission(rp *domain.RolePermission) error {
	return r.db.Create(rp).Error
}

func (r *roleRepository) DeleteRolePermissionsByRoleID(roleID uint) error {
	return r.db.Where("role_id = ?", roleID).Delete(&domain.RolePermission{}).Error
}