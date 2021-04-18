package repository

import (
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/staszigzag/downloader-music/internal/config"
)

const (
	usersTable      = "users"
	audioTable      = "audio"
	usersAudioTable = "users_audio"
)

func NewPostgresDB(cfg config.DbConfig, log logrus.FieldLogger) (*sqlx.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
		cfg.Sslmode,
	)

	// Prep config
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf(" failed to parse config: %v", err)
	}

	// Prep logger
	connConfig.Logger = logrusadapter.NewLogger(log)
	// config.LogLevel, err = pgx.LogLevelFromString()

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf(" failed to create connection db: %v", err)
	}

	// Configure pool limits
	// TODO
	// conn.SetConnMaxIdleTime(cfg.connMaxIdle)
	// conn.SetConnMaxLifetime(cfg.connMaxLife)
	// conn.SetMaxIdleConns(cfg.connMaxOpenIdle)
	// conn.SetMaxOpenConns(cfg.connMaxOpen)

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf(" error to ping connection pool: %v", err)
	}
	log.Info("Data base Postgres is started...")
	return dbx, nil
}
