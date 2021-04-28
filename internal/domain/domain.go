package domain

import "time"

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	UserName  string `json:"userName" db:"user_name"`
	ChatId    int64  `json:"chatId" db:"chat_id"`
}

type Audio struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	VideoId   string    `json:"video_id" db:"video_id"`
	Path      string    `json:"path" db:"path"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
