package summary

type SummaryRequest struct {
	URL     string
	Content string
}

type SummaryResponse struct {
	URL     string
	Summary string
}
