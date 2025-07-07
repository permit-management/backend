package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/permit-management/backend/internal/handler/v1"
	"github.com/permit-management/backend/pkg/setting"
	"gorm.io/gorm"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration, db *gorm.DB) {
	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	apiv1 := r.Group("/api/v1/permit")
	// apiv1.Use(middleware.Authenticated()) // Uncomment this when middleware is ready

	// === Tag routes ===
	tagHandler := v1.NewTagHandler(db, cfg)
	tags := apiv1.Group("/tags")
	{
		tags.POST("", tagHandler.Create)
		tags.GET("", tagHandler.List)
		tags.GET("/:id", tagHandler.Get)
		tags.PUT("/:id", tagHandler.Update)
		tags.DELETE("/:id", tagHandler.Delete)
	}

	// === User routes ===
	userHandler := v1.NewUserHandler(db)
	users := apiv1.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.List)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	// === Departement routes ===
	departementHandler := v1.NewDepartementHandler(db, cfg)
	departements := apiv1.Group("/departements")
	{
		departements.POST("", departementHandler.Create)
		departements.GET("", departementHandler.List)
		departements.GET("/:id", departementHandler.Get)
		departements.PUT("/:id", departementHandler.Update)
		departements.DELETE("/:id", departementHandler.Delete)
	}

	// === Role routes ===
	roleHandler := v1.NewRoleHandler(db, cfg)
	roles := apiv1.Group("/roles")
	{
		roles.POST("", roleHandler.Create)
		roles.GET("", roleHandler.List)
		roles.GET("/:id", roleHandler.Get)
		roles.PUT("/:id", roleHandler.Update)
		roles.DELETE("/:id", roleHandler.Delete)
	}
}
