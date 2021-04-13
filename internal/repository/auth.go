package repository

import (
	"github.com/staszigzag/downloader-music/internal/domain"
)

// AuthPostgres
type Auth struct {
	db interface{}
}

func NewAuth(db interface{}) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user domain.User) (int, error) {
	return 42, nil
}
