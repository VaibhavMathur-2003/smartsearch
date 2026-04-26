package searchrepo

type Urls struct {
	ID      int    `json:"id"`
	Domain  string `json:"domain"`
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
