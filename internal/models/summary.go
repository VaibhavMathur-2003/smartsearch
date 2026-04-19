package models

type Summary struct {
	ID      int    `json:"id"`
	UrlID   string `json:"url_id"`
	Summary string `json:"summary"`
}
