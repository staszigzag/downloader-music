package domain

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	ChatId int64  `json:"chat_id"`
}

type Audio struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	VideoId string `json:"video_id"`
	Path    string `json:"path"`
	// created_at
}
