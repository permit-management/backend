package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/pkg/app"
)

// TagRepository interface to be implemented by repository
type TagRepository interface {
	GetTag(id uint) (*domain.TagModel, error)
	GetTagList(pager *app.Pager) ([]*domain.TagModel, error)
	CountTag(pager *app.Pager) (int64, error)
	CreateTag(name string) (*domain.TagModel, error)
	UpdateTag(id uint, name string) (*domain.TagModel, error)
	DeleteTag(id uint) error
}

type tagRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewTagRepository(ctx context.Context, db *gorm.DB) TagRepository {
	return &tagRepository{
		ctx: ctx,
		db:  db.WithContext(ctx),
	}
}

func (r *tagRepository) searchCriteria(db *gorm.DB) app.SearchCriteria {
	return func(k string, v string) {
		switch k {
		case "name":
			db = db.Where("UPPER(name) LIKE UPPER(?)", "%"+v+"%")
		}
	}
}

func (d *tagRepository) GetTag(id uint) (*domain.TagModel, error) {
	var tag = &domain.TagModel{
		Model: domain.Model{ID: id},
	}
	err := d.db.First(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (d *tagRepository) GetTagList(pager *app.Pager) ([]*domain.TagModel, error) {
	var tags []*domain.TagModel
	db := d.db.Model(&domain.TagModel{})
	pager.SearchCriteria(d.searchCriteria(db))
	if err := db.Offset(pager.Offset()).Limit(pager.PageSize).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (d *tagRepository) CountTag(pager *app.Pager) (int64, error) {
	var count int64
	db := d.db.Model(&domain.TagModel{})
	pager.SearchCriteria(d.searchCriteria(db))
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *tagRepository) CreateTag(name string) (*domain.TagModel, error) {
	tag := domain.TagModel{
		Name: name,
	}
	if err := d.db.Create(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (d *tagRepository) UpdateTag(id uint, name string) (*domain.TagModel, error) {
	values := map[string]any{}
	if name != "" {
		values["name"] = name
	}
	result := d.db.Model(&domain.TagModel{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("row not found")
	}

	var tag domain.TagModel
	d.db.First(&tag, id)
	return &tag, nil
}

func (d *tagRepository) DeleteTag(id uint) error {
	return d.db.Delete(&domain.TagModel{Model: domain.Model{ID: id}}).Error
}
