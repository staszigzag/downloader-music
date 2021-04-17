package repository

import (
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/staszigzag/downloader-music/internal/config"
	"github.com/staszigzag/downloader-music/pkg/logger"
)

const (
	usersTable      = "users"
	audioTable      = "audio"
	usersAudioTable = "users_audio"
)

func NewPostgresDB(cfg config.DbConfig) (*sqlx.DB, error) {
	connConfig := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Database: cfg.Dbname,
		User:     cfg.User,
		Password: cfg.Password,
	}
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		// TODO config
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf(" failed to create connections pool: %v", err)
	}

	nativeDB := stdlib.OpenDBFromPool(connPool)

	db := sqlx.NewDb(nativeDB, "pgx")

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf(" error to ping connection pool: %v", err)
	}
	logger.Info("Data base Postgres is started...")

	return db, nil
}
