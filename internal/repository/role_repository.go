package repository

import (
    "time"

    "github.com/permit-management/backend/internal/domain"
    "gorm.io/gorm"
)

type RoleRepository interface {
    Create(role *domain.Role) error
    GetByID(id uint) (*domain.Role, error)
    GetAll() ([]domain.Role, error)
    Update(role *domain.Role) error
    Delete(id uint) error

    CreateRolePermission(rp *domain.RolePermission) error
    DeleteRolePermissionsByRoleID(roleID uint) error
}

type roleRepo struct {
    db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
    return &roleRepo{db: db}
}

func (r *roleRepo) Create(role *domain.Role) error {
    // GORM otomatis mengisi CreatedAt & UpdatedAt jika pakai autoCreateTime dan autoUpdateTime tag
    return r.db.Create(role).Error
}

func (r *roleRepo) GetByID(id uint) (*domain.Role, error) {
    var role domain.Role
    err := r.db.First(&role, id).Error
    return &role, err
}

func (r *roleRepo) GetAll() ([]domain.Role, error) {
    var roles []domain.Role
    err := r.db.Find(&roles).Error
    return roles, err
}

func (r *roleRepo) Update(role *domain.Role) error {
    // Update fields role_id, role_name dan updated_at secara manual pakai Updates
    return r.db.Model(&domain.Role{}).Where("id = ?", role.ID).Updates(map[string]interface{}{
        "role_id":    role.RoleID,
        "role_name":  role.RoleName,
        "updated_at": time.Now(),
    }).Error
}

func (r *roleRepo) Delete(id uint) error {
    return r.db.Delete(&domain.Role{}, id).Error
}

// RolePermission methods

func (r *roleRepo) CreateRolePermission(rp *domain.RolePermission) error {
    return r.db.Create(rp).Error
}

func (r *roleRepo) DeleteRolePermissionsByRoleID(roleID uint) error {
    return r.db.Where("role_id = ?", roleID).Delete(&domain.RolePermission{}).Error
}
