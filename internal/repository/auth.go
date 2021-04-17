package repository

import (
	"github.com/staszigzag/downloader-music/internal/domain"
)

type AuthRepo struct {
	db interface{}
}

func NewAuthRepo(db interface{}) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(user domain.User) (int, error) {
	return 42, nil
}
