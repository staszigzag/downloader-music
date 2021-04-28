package repository

import (
	"fmt"
	"io"

	"github.com/jmoiron/sqlx"
	"github.com/staszigzag/downloader-music/pkg/filestorage"
)

type AudioRepo struct {
	storage filestorage.FileStorage
	db      *sqlx.DB
}

func NewAudioRepo(storage filestorage.FileStorage, db *sqlx.DB) *AudioRepo {
	return &AudioRepo{storage, db}
}

func (s *AudioRepo) CreateAudioFile(name string, data io.ReadCloser) (filepath string, err error) {
	return s.storage.Create(name, data)
}

func (r *AudioRepo) CreateAudioDb(videoId, name, path string) error {
	query := fmt.Sprintf("INSERT INTO %s (video_id, name, path) values ($1, $2, $3)", audioTable)

	rr, err := r.db.Exec(query, videoId, name, path)
	fmt.Println(rr)
	if err != nil {
		return err
	}

	return nil
}
