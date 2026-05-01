package scrape

type ScrapeData struct {
	ID       int
	URL      string
	RawHTML  string
	Content  string
	Metadata map[string]any
}
