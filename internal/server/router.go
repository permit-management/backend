package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/permit-management/backend/internal/handler/v1"
	"github.com/permit-management/backend/pkg/setting"
	"gorm.io/gorm"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration, db *gorm.DB) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiv1 := r.Group("/api/v1/permit-management")
	// apiv1.Use(middleware.Authenticated())

	var tagHandler = v1.NewTagHandler(db, cfg)
	tags := apiv1.Group("/tags")
	tags.POST("", tagHandler.Create)
	tags.DELETE("/:id", tagHandler.Delete)
	tags.PUT("/:id", tagHandler.Update)
	tags.GET("", tagHandler.List)
	tags.GET("/:id", tagHandler.Get)

}
