package domain

import "time"

type Role struct {
	ID          uint             `json:"id" gorm:"primaryKey"`
	RoleID      string           `json:"role_id" gorm:"unique;not null"`
	RoleName    string           `json:"role_name" gorm:"not null"`
	CreatedAt   time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	Permissions []RolePermission `json:"permissions" gorm:"-"`
}

func (Role) TableName() string {
	return "tbl_role"
}

type RolePermission struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
	AllowRead    bool `json:"allow_read"`
	AllowCreate  bool `json:"allow_create"`
	AllowUpdate  bool `json:"allow_update"`
	AllowDelete  bool `json:"allow_delete"`
}

func (RolePermission) TableName() string {
	return "tbl_role_permissions"
}
