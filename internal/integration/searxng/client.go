package searxng

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) Search(ctx context.Context, req *SearxRequest) (*SearxResponse, error) {
	params := url.Values{}
	params.Set("q", req.Query)
	params.Set("format", "json")

	endpoint := c.baseURL + "/search?" + params.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("searx error: status %d", resp.StatusCode)
	}

	var out SearxResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
