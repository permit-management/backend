package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/pkg/convert"
	"github.com/permit-management/backend/pkg/setting"
)

type Pager struct {
	Page     int    `json:"page"`
	PageSize int    `json:"size"`
	Search   string `json:"search"`
}

func NewPager(c *gin.Context, cfg *setting.Configuration) *Pager {
	return &Pager{
		Page:     getPage(c),
		PageSize: getPageSize(c, cfg.App.DefaultPageSize, cfg.App.MaxPageSize),
		Search:   c.Query("search"),
	}
}

func (p *Pager) Offset() int {
	result := 0
	if p.Page > 0 {
		result = (p.Page - 1) * p.PageSize
	}
	return result
}

func (p *Pager) TotalPage(count int) int {
	return (count + p.PageSize - 1) / p.PageSize
}

type SearchCriteria func(string, string)

func (p *Pager) SearchCriteria(cb SearchCriteria) {
	if strings.TrimSpace(p.Search) != "" {
		searchStr := strings.Split(p.Search, ",")
		for i := 0; i < len(searchStr); i++ {
			vals := strings.Split(searchStr[i], ":")
			if len(vals) == 2 {
				key, val := strings.TrimSpace(vals[0]), strings.TrimSpace(vals[1])
				if key != "" && val != "" {
					cb(key, val)
				}
			}
		}
	}
}

func getPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func getPageSize(c *gin.Context, defaultSize int, maxSize int) int {
	pageSize := convert.StrTo(c.Query("size")).MustInt()
	if pageSize <= 0 {
		return defaultSize
	}
	if pageSize > maxSize {
		return maxSize
	}

	return pageSize
}
