package config

var defaultConfig = map[string]interface{}{
	"App.RunMode":               "release",
	"App.Listen":                "0.0.0.0",
	"App.HTTPPort":              8080,
	"App.ReadTimeout":           "30s",
	"App.WriteTimeout":          "60s",
	"App.DefaultPageSize":       50,
	"App.MaxPageSize":           100,
	"App.ServerShutdownTimeout": "30s",

	"Log.LogSavePath": "logs",
	"Log.LogFileName": "application.log",
	"Log.MaxSize":     20,
	"Log.MaxBackups":  10,
	"Log.Compress":    false,
	"Log.Level":       "trace",

	"Database.Username":     "postgres",
	"Database.Password":     "password",
	"Database.Host":         "127.0.0.1",
	"Database.Port":         5432,
	"Database.DBName":       "SMMF_DEV",
	"Database.Schema":       "public",
	"Database.Charset":      "utf8",
	"Database.ParseTime":    true,
	"Database.MaxIdleConns": 10,
	"Database.MaxOpenConns": 30,
	"Database.MigrationURL": "file://db/migration",

	"Ratelimit.Enable": false,
}
