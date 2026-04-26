package ollama

import (
	"context"
	"errors"

	"github.com/ollama/ollama/api"
)

type Client struct {
	client *api.Client
}

func NewClient() (*Client, error) {
	c, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, err
	}
	return &Client{client: c}, nil
}

func (c *Client) Generate(
	ctx context.Context,
	req *OllamaRequest,
) (*OllamaResponse, error) {

	if req.Model == "" {
		return nil, errors.New("model is required")
	}

	stream := false
	if req.Stream != nil {
		stream = *req.Stream
	}

	reqOllama := &api.GenerateRequest{
		Model:  req.Model,
		System: req.System,
		Prompt: req.Prompt,
		Stream: &stream,
	}

	var response string
	var totalDuration int64

	err := c.client.Generate(ctx, reqOllama, func(resp api.GenerateResponse) error {
		response += resp.Response
		if resp.Done {
			totalDuration = resp.TotalDuration.Milliseconds()
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &OllamaResponse{
		Model:         req.Model,
		Response:      response,
		TotalDuration: totalDuration,
	}, nil
}
