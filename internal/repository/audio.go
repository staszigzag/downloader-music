package repository

import (
	"fmt"
	"io"

	"github.com/staszigzag/downloader-music/internal/domain"

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

func (r *AudioRepo) CreateAudioDb(videoId, name, path string) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (video_id, name, path) values ($1, $2, $3) RETURNING id", audioTable)

	row := r.db.QueryRow(query, videoId, name, path)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AudioRepo) GetAudioDbByVideoId(videoId string) *domain.Audio {
	user := new(domain.Audio)
	query := fmt.Sprintf("SELECT * FROM %s WHERE video_id=$1", audioTable)
	err := r.db.Get(user, query, videoId)
	if err != nil {
		return nil
	}
	return user
}

func (r *AudioRepo) CreateAudioUserLink(audioId, userId int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, audio_id) values ($1, $2)", usersAudioTable)

	rr, err := r.db.Exec(query, userId, audioId)
	fmt.Println(rr)
	if err != nil {
		return err
	}

	return nil
}
