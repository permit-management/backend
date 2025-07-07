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
	"github.com/permit-management/backend/pkg/errcode"
	"github.com/permit-management/backend/pkg/logger"
	"github.com/permit-management/backend/pkg/setting"
)

type userHandler struct {
	db  *gorm.DB
	cfg *setting.Configuration
}

func NewUserHandler(db *gorm.DB, cfg *setting.Configuration) userHandler {
	return userHandler{
		db:  db,
		cfg: cfg,
	}
}

func (h *userHandler) Create(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)

	// Bind JSON body ke struct
	if err := c.ShouldBindJSON(&param); err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "create user")
	svc := service.NewUserService(ctx, h.db)

	user, err := svc.CreateUser(&param)
	if err != nil {
		logger.WithTrace(c).Error(err)
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(user)
}

func (h *userHandler) List(c *gin.Context) {
	response := app.NewPagingResponse[*domain.UserModel](c)
	svc := service.NewUserService(c, h.db)
	pager := app.NewPager(c, h.cfg)

	users, count, err := svc.ListUsers(pager)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(users, pager, count)
}

func (h *userHandler) Get(c *gin.Context) {
	param := constants.IDRequest{
		ID: uint(convert.StrTo(c.Param("id")).MustInt()),
	}
	response := app.NewResponse(c)

	if app.Validation(c, &param, response, false) != nil {
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "get user")
	svc := service.NewUserService(ctx, h.db)

	user, err := svc.GetUser(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(user)
}

func (h *userHandler) Update(c *gin.Context) {
	param := service.UpdateUserRequest{
		ID: uint(convert.StrTo(c.Param("id")).MustInt()),
	}
	response := app.NewResponse(c)

	if err := c.ShouldBindJSON(&param); err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "update user")
	svc := service.NewUserService(ctx, h.db)

	user, err := svc.UpdateUser(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(user)
}

func (h *userHandler) Delete(c *gin.Context) {
	param := constants.IDRequest{
		ID: uint(convert.StrTo(c.Param("id")).MustInt()),
	}
	response := app.NewResponse(c)

	if app.Validation(c, &param, response, true) != nil {
		return
	}

	ctx := context.WithValue(c.Request.Context(), "username", "delete user")
	svc := service.NewUserService(ctx, h.db)

	if err := svc.DeleteUser(&param); err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{})
}
