package service

import (
	"context"

	"github.com/staszigzag/downloader-music/internal/domain"
	"github.com/staszigzag/downloader-music/internal/repository"
	"github.com/staszigzag/downloader-music/pkg/youtubedl"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Downloader interface {
	Download(ctx context.Context, url string) (string, error)
}

type Services struct { // Authorization
	Authorization
	Downloader
}

type Deps struct {
	Repos      *repository.Repository
	Downloader youtubedl.Downloader
	// TokenManager           auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		Authorization: NewAuthService(deps.Repos),
		Downloader:    NewDownloaderService(deps.Downloader, deps.Repos),
	}
}
