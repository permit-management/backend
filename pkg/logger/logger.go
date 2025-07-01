package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/pkg/setting"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *lumberjack.Logger

func SetupLogger(logset *setting.LogSettingS) {
	// TODO: use lumberjack.Logger as config
	logger := &lumberjack.Logger{
		Filename:   filepath.Join(logset.LogSavePath, logset.LogFileName),
		MaxSize:    logset.MaxSize, // megabytes
		MaxBackups: logset.MaxBackups,
		MaxAge:     14,              // days
		Compress:   logset.Compress, // disabled by default
	}

	lvl, err := logrus.ParseLevel(logset.Level)
	if err != nil {
		logrus.SetLevel(lvl)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// use lumberjack to write to implement rotation.
	mw := io.MultiWriter(os.Stdout, logger)

	logrus.SetOutput(mw) // set output to file and console at the same time
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Set default log output to logrus
	log.SetOutput(logrus.StandardLogger().Writer())
}

func Close() {
	logger.Close()
}

func WithTrace(ctx *gin.Context) *logrus.Entry {
	fields := logrus.Fields{}
	if len(ctx.GetString("X-Trace-ID")) > 0 {
		fields["trace_id"] = ctx.GetString("X-Trace-ID")
	}
	if len(ctx.GetString("X-Span-ID")) > 0 {
		fields["span_id"] = ctx.GetString("X-Span-ID")
	}
	if len(ctx.GetString("Authorization")) > 0 {
		fields["auth_token"] = ctx.GetString("Authorization")
	}
	return logrus.WithFields(fields)
}

func Log() *logrus.Logger {
	return logrus.StandardLogger()
}
