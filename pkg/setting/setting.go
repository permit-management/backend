package setting

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	Conf = &Configuration{}
)

type AppSettingS struct {
	RunMode      string        `mapstructure:"RunMode" env:"APP_MODE"`
	Listen       string        `mapstructure:"Listen" env:"APP_LISTEN"`
	HTTPPort     string        `mapstructure:"HTTPPort" env:"APP_PORT"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout" env:"APP_READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout" env:"APP_WRITE_TIMEOUT"`

	DefaultPageSize       int           `mapstructure:"DefaultPageSize" env:"APP_PAGE_SIZE"`
	MaxPageSize           int           `mapstructure:"MaxPageSize" env:"APP_MAX_PAGE"`
	ServerShutdownTimeout time.Duration `mapstructure:"ServerShutdownTimeout" env:"APP_SHUTDOWN_TIMEOUT"`
	JWTSecret             string        `mapstructure:"JWTSecret" env:"APP_JWT_SECRET"`

	UploadFolder string `mapstructure:"UploadFolder" env:"APP_UPLOAD_FOLDER"`
}

type LogSettingS struct {
	LogSavePath string `mapstructure:"LogSavePath" env:"LOG_PATH"`
	LogFileName string `mapstructure:"LogFileName" env:"LOG_NAME"`
	MaxSize     int    `mapstructure:"MaxSize" env:"LOG_MAX_SIZE"`
	MaxBackups  int    `mapstructure:"MaxBackups" env:"LOG_MAX_BACKUP"`
	Compress    bool   `mapstructure:"Compress" env:"LOG_COMPRESS"`
	Level       string `mapstructure:"Level" env:"LOG_LEVEL"`
}

type DatabaseSettingS struct {
	Username     string `mapstructure:"Username" env:"DB_USER"`
	Password     string `mapstructure:"Password" env:"DB_PASSWORD"`
	Host         string `mapstructure:"Host" env:"DB_HOST"`
	Port         int    `mapstructure:"Port" env:"DB_PORT"`
	DBName       string `mapstructure:"DBName" env:"DB_NAME"`
	Schema       string `mapstructure:"Schema" env:"DB_SCHEMA"`
	Charset      string `mapstructure:"Charset" env:"DB_CHARSET"`
	ParseTime    bool   `mapstructure:"ParseTime" env:"DB_PARSE_TIME"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns" env:"DB_CONN_IDLE"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns" env:"DB_CONN_OPEN"`
	MigrationURL string `mapstructure:"MigrationURL"`
}

type Configuration struct {
	App      AppSettingS      `mapstructure:"App"`
	Log      LogSettingS      `mapstructure:"Log"`
	Database DatabaseSettingS `mapstructure:"Database"`
}

func Load(cfgFile string, defaultConfig map[string]interface{}) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("configs/")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	for key, val := range defaultConfig {
		viper.SetDefault(key, val)
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.Unmarshal(Conf)
	if err != nil {
		return err
	}

	return nil
}
