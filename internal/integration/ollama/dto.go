package ollama

type OllamaRequest struct {
	Model  string `json:"model"`
	System string `json:"system,omitempty"`
	Prompt string `json:"prompt"`
	Format string `json:"format,omitempty"`
	Stream *bool  `json:"stream"`
}

type OllamaResponse struct {
	Model         string `json:"model"`
	Response      string `json:"response"`
	TotalDuration int64  `json:"total_duration"`
}
