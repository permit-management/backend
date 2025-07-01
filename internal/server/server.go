package server

import (
	"context"
	"net/http"

	"github.com/permit-management/backend/internal/middleware"
	"github.com/permit-management/backend/pkg/setting"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//	@title			Template Service
//	@description	Template Service API
//	@contact.name	SMMF Digital Team
//	@contact.url	https://github.com/permit-management/backend
//	@license.name	MIT License
//	@license.url	https://github.com/permit-management/backend/blob/main/LICENSE
//	@host			localhost:8080
//	@BasePath		/
//	@schemes		http https

type Server struct {
	Router *gin.Engine
	Svr    *http.Server
	Config *setting.Configuration
	DB     *gorm.DB
}

func NewServer(cfg *setting.Configuration, db *gorm.DB) *Server {
	r := gin.New()

	middleware.UseDefault(r, cfg)

	SetRouters(r, cfg, db)

	srv := &http.Server{
		Addr:           cfg.App.Listen + ":" + cfg.App.HTTPPort,
		Handler:        r,
		ReadTimeout:    cfg.App.ReadTimeout,
		WriteTimeout:   cfg.App.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		Config: cfg,
		Svr:    srv,
		Router: r,
		DB:     db,
	}
}

func (s *Server) Start() error {
	// Timeout: https://adam-p.ca/blog/2022/01/golang-http-server-timeouts/
	go func() {
		log.Info("Starting HTTP Server at :", s.Config.App.HTTPPort)
		if err := s.Svr.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("HTTP Server exception. ", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Svr.Shutdown(ctx)
}
