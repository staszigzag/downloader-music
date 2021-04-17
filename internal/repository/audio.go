package repository

import (
	"io"

	"github.com/staszigzag/downloader-music/pkg/filestorage"
)

type AudioRepo struct {
	storage filestorage.FileStorage
}

func NewAudioRepo(storage filestorage.FileStorage) *AudioRepo {
	return &AudioRepo{storage}
}

func (s *AudioRepo) CreateAudio(name string, data io.ReadCloser) (filepath string, err error) {
	return s.storage.Create(name, data)
}
