package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
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
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
	// 	databaseSetting.Host,
	// 	databaseSetting.Port,
	// 	databaseSetting.Username,
	// 	databaseSetting.Password,
	// 	databaseSetting.DBName,
	// )
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC", databaseSetting.Username,
		databaseSetting.Password, databaseSetting.Host, databaseSetting.Port, databaseSetting.DBName)

	log.Println(dsn)

	// https://gorm.io/docs/performance.html
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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
