package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/permit-management/backend/internal/handler/v1"
	"github.com/permit-management/backend/internal/middleware"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/internal/service"
	"github.com/permit-management/backend/pkg/setting"
	"gorm.io/gorm"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration, db *gorm.DB) {
	// health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Static("/uploads", cfg.App.UploadFolder)

	// auth
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg.App.JWTSecret)
	authHandler := v1.NewAuthHandler(authService)

	//auth mobile
	authMobileRepo := repository.NewAuthMobileRepository(db)
	authMobileService := service.NewAuthMobileService(authMobileRepo, cfg.App.JWTSecret)
	authMobileHandler := v1.NewAuthMobileHandler(authMobileService)

	auth := r.Group("/api/v1/permit/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)

		//ini route khusus mobile login
		auth.POST("/mobile/login", authMobileHandler.Login)
	}
	// protected routes
	protected := r.Group("/api/v1/permit")
	protected.Use(middleware.JWT())
	unprotected := r.Group("/api/permit")

	// handlers
	userHandler := v1.NewUserHandler(db)
	tagHandler := v1.NewTagHandler(db, cfg)
	departementHandler := v1.NewDepartementHandler(db, cfg)
	roleHandler := v1.NewRoleHandler(db, cfg)
	permitHandler := v1.NewPermitHandler(db, cfg)
	workTypeHandler := v1.NewWorkTypeHandler(db, cfg)

	// repository & service utk permit approvals
	permitApprovalRepo := repository.NewPermitApprovalRepository(db)
	permitRepo := repository.NewPermitRepository(db)
	permitApprovalService := service.NewPermitApprovalService(permitApprovalRepo, permitRepo)
	permitApprovalHandler := v1.NewPermitApprovalHandler(permitApprovalService)

	// repository & service utk check-in
	checkInRepo := repository.NewCheckInRepository(db)
	checkInService := service.NewCheckInService(checkInRepo)
	checkInHandler := v1.NewCheckInHandler(checkInService)

	// repository & service utk daily work check
	dailyWorkCheckRepo := repository.NewDailyWorkCheckRepository(db)
	activityRepo := repository.NewActivityRepository(db)
	dailyWorkCheckService := service.NewDailyWorkCheckService(dailyWorkCheckRepo, activityRepo)
	dailyWorkCheckHandler := v1.NewDailyWorkCheckHandler(dailyWorkCheckService)

	incidentReportRepo := repository.NewIncidentReportRepository(db)
	incidentReportService := service.NewIncidentReportService(incidentReportRepo)
	incidentReportHandler := v1.NewIncidentReportHandler(incidentReportService)

	// untuk user
	users := protected.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.List)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	// untuk tags
	tags := protected.Group("/tags")
	{
		tags.POST("", tagHandler.Create)
		tags.GET("", tagHandler.List)
		tags.GET("/:id", tagHandler.Get)
		tags.PUT("/:id", tagHandler.Update)
		tags.DELETE("/:id", tagHandler.Delete)
	}

	// untuk departements
	departements := protected.Group("/departements")
	{
		departements.POST("", departementHandler.Create)
		departements.GET("", departementHandler.List)
		departements.GET("/:id", departementHandler.Get)
		departements.PUT("/:id", departementHandler.Update)
		departements.DELETE("/:id", departementHandler.Delete)
	}

	// untuk roles
	roles := protected.Group("/roles")
	{
		roles.POST("", roleHandler.Create)
		roles.GET("", roleHandler.List)
		roles.GET("/:id", roleHandler.Get)
		roles.PUT("/:id", roleHandler.Update)
		roles.DELETE("/:id", roleHandler.Delete)
	}

	// untuk permits
	permits := protected.Group("/permits")
	{
		permits.POST("", permitHandler.CreatePermit)
		permits.GET("", permitHandler.GetAllPermits)
		permits.GET("/:id", permitHandler.GetPermitByID)
		permits.PUT("/:id", permitHandler.UpdatePermit)
		permits.DELETE("/:id", permitHandler.DeletePermit)

		unprotected.POST("/permits", permitHandler.CreatePermit)
	}

	// untuk work types
	workTypes := protected.Group("/work-types")
	{
		workTypes.POST("", workTypeHandler.Create)
		workTypes.GET("", workTypeHandler.List)
		workTypes.GET("/:id", workTypeHandler.Get)
		workTypes.PUT("/:id", workTypeHandler.Update)
		workTypes.DELETE("/:id", workTypeHandler.Delete)

		unprotected.GET("/work-types", workTypeHandler.List)
		unprotected.GET("/work-types/:id", workTypeHandler.Get)
	}

	// untuk permit approvals
	permitApprovals := protected.Group("/permit-approvals")
	{
		permitApprovals.POST("", permitApprovalHandler.ApprovePermit)
		// kalau mau lihat history approval per permit
		// permitApprovals.GET("/:permit_id", permitApprovalHandler.ListByPermitID)
	}

	authMobile := r.Group("/api/v1/mobile/auth")
	{
		authMobile.POST("/login", authMobileHandler.Login)
	}

	// untuk check-in
	checkin := r.Group("/api/v1/mobile")
	{
		checkin.POST("/checkin", checkInHandler.CheckIn)
	}

	// untuk mobile (daily work check)
	mobile := r.Group("/api/v1/mobile")
	{
		// worker selesai mengerjakan aktivitas
		mobile.POST("/daily-work-check/done", dailyWorkCheckHandler.MarkDone)

		// worker ambil semua aktivitas berdasarkan permit & nik
		mobile.GET("/daily-work-check/:permit_id/:nik", dailyWorkCheckHandler.GetActivitiesByWorker)
	}

	// untuk mobile (incident report)
	mobileIncident := r.Group("/api/v1/mobile")
	{
		mobileIncident.POST("/incident-report", incidentReportHandler.Create)
	}
}
