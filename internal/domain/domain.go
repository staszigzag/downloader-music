package domain

type User struct {
	Name string `json:"name"`
}

type Audio struct {
	Name    string `json:"name"`
	VideoId string `json:"videoId"`
	Path    string `json:"path"`
}
