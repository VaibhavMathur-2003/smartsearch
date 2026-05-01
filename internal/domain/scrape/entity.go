package scrapedomain

type ScrapeRequest struct {
	URL string
}

type ScrapeData struct {
	URL      string         `json:"url"`
	RawHTML  string         `json:"html"`
	Content  string         `json:"content"`
	Metadata map[string]any `json:"metadata"`
}
