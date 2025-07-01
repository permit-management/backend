package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/pkg/setting"
	log "github.com/sirupsen/logrus"
)

// UseDefault set the default middleware
func UseDefault(r *gin.Engine, cfg *setting.Configuration) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.Static("/static", "./web/dist")
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

}
