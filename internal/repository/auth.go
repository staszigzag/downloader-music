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

func (r *AuthRepo) CreateUser(user *domain.User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, first_name, last_name, user_name, chat_id) values ($1, $2, $3, $4, $5)", usersTable)

	rr, err := r.db.Exec(query, user.Id, user.FirstName, user.LastName, user.UserName, user.ChatId)
	fmt.Println(rr)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) GetUser(userId int) (*domain.User, error) {
	user := new(domain.User)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(user, query, userId)
	return user, err
}
