package domain

type Website struct {
	ID    int    `json:"id"`
	UrlID int    `json:"url_id"`
	Text  string `json:"text"`
}
