package searchdomain

type SearchRequest struct {
	Query string `json:"q"`
}

type SearchResult struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SearchResponse struct {
	Results []SearchResult `json:"results"`
}
