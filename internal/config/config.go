package config

import (
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/pkg/logger"
	"github.com/permit-management/backend/pkg/setting"
)

func AppInit(cfgFile string) error {
	if err := setting.Load(cfgFile, defaultConfig); err != nil {
		return errors.Wrap(err, "loading config file failed")
	}
	gin.SetMode(setting.Conf.App.RunMode)
	logger.SetupLogger(&setting.Conf.Log)
	return nil
}
