package db

import (
	"database/sql"
	"fmt"
	"github.com/ingener-vladimir/go_practices/db_generator/internal/config"
)

const (
	dsnFormat = "%s:%s@tcp(%s:%d)/%s?timeout=5s"
)

func NewConnection(cfg *config.Config) (conn *sql.DB, err error) {
	dsn := fmt.Sprintf(
		dsnFormat,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.NameDB,
	)

	db, err := sql.Open(cfg.Type, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
