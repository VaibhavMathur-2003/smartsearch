package domain

type Urls struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Searx struct {
	UrlData []Urls `json:"results"`
}
