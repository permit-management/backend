package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/permit-management/backend/internal/constants"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/service"
	"github.com/permit-management/backend/pkg/app"
	"github.com/permit-management/backend/pkg/convert"
	_ "github.com/permit-management/backend/pkg/errcode"
	"github.com/permit-management/backend/pkg/logger"

	"github.com/permit-management/backend/pkg/setting"
)

type tagHandler struct {
	db  *gorm.DB
	cfg *setting.Configuration
}

func NewTagHandler(db *gorm.DB, cfg *setting.Configuration) tagHandler {
	return tagHandler{
		db:  db,
		cfg: cfg,
	}
}

func (h *tagHandler) List(c *gin.Context) {
	response := app.NewPagingResponse[*domain.TagModel](c)
	svc := service.NewTagService(c, h.db)
	pager := app.NewPager(c, h.cfg)

	data, cnt, err := svc.GetTagListWithCnt(pager)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(data, pager, cnt)
}

func (h *tagHandler) Get(c *gin.Context) {
	param := constants.IDRequest{ID: uint(convert.StrTo(c.Param("id")).MustInt())}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response, false) != nil {
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "current user")
	svc := service.NewTagService(ctx, h.db)
	tag, err := svc.GetTag(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(tag)
}

func (h *tagHandler) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response, true) != nil {
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "create user")
	svc := service.NewTagService(ctx, h.db)
	tag, err := svc.CreateTag(&param)
	if err != nil {
		logger.WithTrace(c).Info(err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(tag)
}

func (h *tagHandler) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: uint(convert.StrTo(c.Param("id")).MustInt()),
	}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response, true) != nil {
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "update user")
	svc := service.NewTagService(ctx, h.db)
	tag, err := svc.UpdateTag(&param)
	if err != nil {
		logger.WithTrace(c).Info(err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(tag)
}

func (h *tagHandler) Delete(c *gin.Context) {
	param := constants.IDRequest{
		ID: uint(convert.StrTo(c.Param("id")).MustInt()),
	}
	response := app.NewResponse(c)
	if app.Validation(c, &param, response, true) != nil {
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "delete user")
	svc := service.NewTagService(ctx, h.db)
	if err := svc.DeleteTag(&param); err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{})
}
