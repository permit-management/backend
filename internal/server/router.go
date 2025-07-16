package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/permit-management/backend/internal/handler/v1"
	"github.com/permit-management/backend/pkg/setting"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/internal/service"
	"github.com/permit-management/backend/internal/middleware"
	"gorm.io/gorm"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration, db *gorm.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg.App.JWTSecret)
	authHandler := v1.NewAuthHandler(authService)

	auth := r.Group("/api/v1/permit/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api/v1/permit")
	protected.Use(middleware.JWT())

	userHandler := v1.NewUserHandler(db)
	tagHandler := v1.NewTagHandler(db, cfg)
	departementHandler := v1.NewDepartementHandler(db, cfg)
	roleHandler := v1.NewRoleHandler(db, cfg)

	users := protected.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.List)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	tags := protected.Group("/tags")
	{
		tags.POST("", tagHandler.Create)
		tags.GET("", tagHandler.List)
		tags.GET("/:id", tagHandler.Get)
		tags.PUT("/:id", tagHandler.Update)
		tags.DELETE("/:id", tagHandler.Delete)
	}

	departements := protected.Group("/departements")
	{
		departements.POST("", departementHandler.Create)
		departements.GET("", departementHandler.List)
		departements.GET("/:id", departementHandler.Get)
		departements.PUT("/:id", departementHandler.Update)
		departements.DELETE("/:id", departementHandler.Delete)
	}

	roles := protected.Group("/roles")
	{
		roles.POST("", roleHandler.Create)
		roles.GET("", roleHandler.List)
		roles.GET("/:id", roleHandler.Get)
		roles.PUT("/:id", roleHandler.Update)
		roles.DELETE("/:id", roleHandler.Delete)
	}
}