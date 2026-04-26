package searxng

type SearxRequest struct {
	Query      string   `json:"q"`
	Engines    []string `json:"engines,omitempty"`
	Language   string   `json:"language,omitempty"`
	Page       int      `json:"pageno,omitempty"`
	SafeSearch int      `json:"safesearch,omitempty"`
}

type SearxResult struct {
	URL     string `json:"url"`
	Engine  string `json:"engine"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SearxResponse struct {
	Results []SearxResult `json:"results"`
}
