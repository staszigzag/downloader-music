package repository

import (
	"io"

	"github.com/staszigzag/downloader-music/pkg/filestorage"

	"github.com/staszigzag/downloader-music/internal/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Audio interface {
	CreateAudio(name string, data io.ReadCloser) (filepath string, err error)
}

type Repository struct {
	Authorization
	Audio
}

func NewRepository(db interface{}, storage filestorage.FileStorage) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Audio:         NewAudioRepo(storage),
	}
}
