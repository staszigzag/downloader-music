package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/staszigzag/downloader-music/internal/domain"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, chat_id) values ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.ChatId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
