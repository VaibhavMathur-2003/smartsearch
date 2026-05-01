package searchrepo

type Urls struct {
	ID      int    `json:"id"`
	Domain  string `json:"domain"`
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
