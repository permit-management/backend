package main

import (
	"github.com/permit-management/backend/db"
	"github.com/permit-management/backend/internal/config"
	"github.com/permit-management/backend/internal/server"
	"github.com/permit-management/backend/pkg/setting"
	"github.com/permit-management/backend/pkg/signals"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := config.AppInit(""); err != nil {
		log.Error("init server failed.", err)
		return
	}

	// init db
	engine, err := db.DBInit(&setting.Conf.Database)
	if err != nil {
		log.Error("init db failed.", err)
		return
	}

	// start http server
	svr := server.NewServer(setting.Conf, engine)
	if err := svr.Start(); err != nil {
		log.Error("init server failed.", err)
		return
	}

	// graceful shutdown
	stopCh := signals.SetupSignalHandler()
	sd, _ := signals.NewShutdown(setting.Conf.App.ServerShutdownTimeout)
	sd.Graceful(stopCh, svr, engine)
}
