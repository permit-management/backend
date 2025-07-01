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

type CreateTagRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
	Desc string `json:"desc" binding:"required,min=2,max=100"`
}

type UpdateTagRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"max=100"`
}

// // TagEntity is exported and used by the delivery layer.
// type TagEntity struct {
// 	ID        string  `json:"id"`
// 	Name      string  `json:"name"`
// 	CreatedBy string  `json:"created_by"`
// 	UpdatedBy *string `json:"updated_by"`
// }

// func (e domain.TagEntity) FromDomain(d domain.TagModel) *TagEntity {
// 	e.ID = d.ID
// 	e.Name = d.Name
// 	e.CreatedBy = d.CreatedBy
// 	e.UpdatedBy = d.UpdatedBy
// 	return e
// }

type TagService struct {
	ctx  context.Context
	db   *gorm.DB
	repo repository.TagRepository
}

func NewTagService(ctx context.Context, db *gorm.DB) *TagService {
	svc := &TagService{ctx: ctx, db: db}
	svc.repo = repository.NewTagRepository(ctx, db)
	return svc
}

func (svc *TagService) GetTag(param *constants.IDRequest) (*domain.TagModel, *errcode.Error) {
	tag, err := svc.repo.GetTag(param.ID)
	if err != nil {
		return nil, errcode.BadRequest.WithDetails(err.Error())
	}
	return tag, nil
}

func (svc *TagService) GetTagListWithCnt(pager *app.Pager) ([]*domain.TagModel, int, *errcode.Error) {
	cnt, err := svc.repo.CountTag(pager)
	if err != nil {
		return nil, 0, errcode.BadRequest.WithDetails(err.Error())
	}

	tags, err := svc.repo.GetTagList(pager)
	if err != nil {
		return nil, 0, errcode.BadRequest.WithDetails(err.Error())
	}

	// Convert to excpected type
	// tagList := []domain.TagEntity{}
	// for _, tag := range tags {
	// 	tagList = append(tagList, (&TagEntity{}).FromDomain(*tag))
	// }

	return tags, int(cnt), nil
}

func (svc *TagService) CreateTag(param *CreateTagRequest) (*domain.TagModel, *errcode.Error) {
	model, err := svc.repo.CreateTag(param.Name)
	if err != nil {
		return nil, errcode.BadRequest.WithDetails(err.Error())
	}
	return model, nil
}

func (svc *TagService) UpdateTag(param *UpdateTagRequest) (*domain.TagModel, *errcode.Error) {
	model, err := svc.repo.UpdateTag(param.ID, param.Name)
	if err != nil {
		return nil, errcode.BadRequest.WithDetails(err.Error())
	}
	return model, nil
}

func (svc *TagService) DeleteTag(param *constants.IDRequest) *errcode.Error {
	err := svc.repo.DeleteTag(param.ID)
	if err != nil {
		return errcode.BadRequest.WithDetails(err.Error())
	}
	return nil
}
