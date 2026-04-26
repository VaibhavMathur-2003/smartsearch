package llm

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func LlmCall(ctx context.Context, model string, systemPrompt string, userPrompt string) (string, error) {
	client, err := api.ClientFromEnvironment()
	stream := false
	req := &api.GenerateRequest{
		Model:  model,
		Prompt: userPrompt,
		Stream: &stream, // default streaming
	}
	response := ""
	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		response = resp.Response
		if resp.Done {
			fmt.Println("\n[Generation complete]")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return response, nil
}
