package service

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/staszigzag/downloader-music/internal/domain"
	"github.com/staszigzag/downloader-music/internal/repository"
	"github.com/staszigzag/downloader-music/pkg/youtubedl"
)

type Authorization interface {
	Authorization(msg *tgbotapi.Message) (*domain.User, error)
	CreateUser(user *domain.User) error
	GetUser(userId int) *domain.User
}

type Downloader interface {
	Download(ctx context.Context, url string, user *domain.User) (string, error)
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
