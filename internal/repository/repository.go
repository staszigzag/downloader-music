package repository

import "github.com/staszigzag/downloader-music/internal/domain"

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db interface{}) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
	}
}
