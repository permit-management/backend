package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/permit-management/backend/pkg/setting"
)

func DBInit(dbc *setting.DatabaseSettingS) (engine *gorm.DB, err error) {
	engine, err = newDBEngine(dbc)
	engine.Debug() // Debug Engine
	if err != nil {
		return nil, err
	}

	return
}

func newDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable TimeZone=UTC",
		databaseSetting.Host,
		databaseSetting.Port,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.DBName,
		databaseSetting.Schema,
	)

	// https://gorm.io/docs/performance.html
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // Default transaction is enabled
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	if setting.Conf.App.RunMode == "debug" {
		db = db.Debug()
	}
	// db.SingularTable(true)

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
