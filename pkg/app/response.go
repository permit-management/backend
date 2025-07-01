package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/pkg/errcode"
)

type BaseResponse interface {
	ToErrorResponse(err *errcode.Error)
}

type Response struct {
	Ctx *gin.Context
}

type ResponseSuccess struct {
	Code int       `json:"code"`
	Msg  string    `json:"message"`
	Data any       `json:"data,omitempty"`
	Time time.Time `json:"timestamp,omitempty"`
}

type ListResponse[T any] struct {
	Page      int       `json:"page"`
	PageSize  int       `json:"size"`
	Data      []T       `json:"data"`
	TotalRows int       `json:"total_record"`
	TotalPage int       `json:"total_page"`
	Time      time.Time `json:"timestamp,omitempty"`
}

// Untuk documentasi swagger
type PagerResponse struct {
	Page      int        `json:"page"`
	PageSize  int        `json:"size"`
	Data      []struct{} `json:"data"`
	TotalRows int        `json:"total_record"`
	TotalPage int        `json:"total_page"`
	Time      time.Time  `json:"timestamp,omitempty"`
}

type MapResponse map[string]string

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data any) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, &ResponseSuccess{
		Code: 0,
		Msg:  "Success",
		Data: data,
		Time: time.Now(),
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	err.WithTimestamp()
	r.Ctx.JSON(err.Status, err)
}

type PagingResponse[T any] struct {
	Response
}

func NewPagingResponse[T any](ctx *gin.Context) *PagingResponse[T] {
	return &PagingResponse[T]{
		Response: *NewResponse(ctx),
	}
}

func (r *PagingResponse[T]) ToResponse(data []T, pager *Pager, count int) {
	r.Ctx.JSON(http.StatusOK, &ListResponse[T]{
		Page:      pager.Page,
		PageSize:  pager.PageSize,
		Data:      data,
		TotalRows: count,
		TotalPage: pager.TotalPage(count),
		Time:      time.Now(),
	})
}
