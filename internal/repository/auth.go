package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/staszigzag/downloader-music/internal/domain"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(user domain.User) (int, error) {
	return 42, nil
}
