package repository

import (
	"io"

	"github.com/jmoiron/sqlx"

	"github.com/staszigzag/downloader-music/pkg/filestorage"

	"github.com/staszigzag/downloader-music/internal/domain"
)

const (
	usersTable      = "users"
	audioTable      = "audio"
	usersAudioTable = "users_audio"
)

type Authorization interface {
	CreateUser(user *domain.User) error
	GetUser(userId int) (*domain.User, error)
}

type Audio interface {
	CreateAudioFile(name string, data io.ReadCloser) (filepath string, err error)
	CreateAudioDb(videoId, name, path string) error
}

type Repository struct {
	Authorization
	Audio
}

func NewRepository(db *sqlx.DB, storage filestorage.FileStorage) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Audio:         NewAudioRepo(storage, db),
	}
}
