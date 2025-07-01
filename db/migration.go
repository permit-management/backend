package db

import (
	"embed"

	"gorm.io/gorm"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var source embed.FS

func Status(conn *gorm.DB) (version uint, rerr error) {
	goose.SetBaseFS(source)

	if err := goose.SetDialect("postgres"); err != nil {
		rerr = err
		return
	}

	db, err := conn.DB()
	if err != nil {
		rerr = err
		return
	}

	if err := goose.Status(db, "migrations"); err != nil {
		rerr = err
		return
	}

	return
}

func Migrate(conn *gorm.DB) (version uint, rerr error) {
	goose.SetBaseFS(source)

	if err := goose.SetDialect("postgres"); err != nil {
		rerr = err
		return
	}

	db, err := conn.DB()
	if err != nil {
		rerr = err
		return
	}

	if err := goose.Up(db, "migrations"); err != nil {
		rerr = err
		return
	}

	return
}

func Rollback(conn *gorm.DB) (version uint, rerr error) {
	goose.SetBaseFS(source)

	if err := goose.SetDialect("postgres"); err != nil {
		rerr = err
		return
	}

	db, err := conn.DB()
	if err != nil {
		rerr = err
		return
	}

	if err := goose.Down(db, "migrations"); err != nil {
		rerr = err
		return
	}

	return
}
