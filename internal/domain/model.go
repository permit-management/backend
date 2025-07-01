package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	CreateByColumn        = "created_by"
	UpdateByColumn        = "updated_by"
	DeleteByColumn        = "deleted_by"
	FlagDeleted    uint32 = 1
	FlagNotDeleted uint32 = 0
)

// Basic type for all models. https://gorm.io/docs/models.html
type Model struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	// ID uint `json:"id" gorm:"primaryKey"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy *string    `json:"updated_by" gorm:"<-:update"`
	DeletedBy *string    `json:"deleted_by,omitempty" gorm:"<-:update"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"<-:update"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("username").(string); ok {
		tx.Statement.SetColumn(CreateByColumn, userID)
	}
	// if m.ID == "" {
	// 	m.ID = uuid.NewString()
	// }
	return nil
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("username").(string); ok {
		tx.Statement.SetColumn(UpdateByColumn, userID)
	}
	return nil
}

func (m *Model) BeforeDelete(tx *gorm.DB) (err error) {
	// FIXME DeleteByColumn not persisted
	if userID, ok := tx.Statement.Context.Value("username").(string); ok {
		tx.Statement.SetColumn(DeleteByColumn, userID)
	}
	return nil
}

func WithUserContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, "username", username)
}
