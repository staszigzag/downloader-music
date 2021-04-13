package service

import (
	"github.com/staszigzag/downloader-music/internal/domain"
	"github.com/staszigzag/downloader-music/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Services struct { // Authorization
	Authorization
}

type Deps struct {
	Repos *repository.Repository
	// TokenManager           auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		Authorization: NewAuthService(deps.Repos),
	}
}
