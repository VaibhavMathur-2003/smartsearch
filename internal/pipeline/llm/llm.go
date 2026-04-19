package llm

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

func LlmCall(ctx context.Context, model string, systemPrompt string, userPrompt string) {
	client, err := api.ClientFromEnvironment()
	req := &api.GenerateRequest{
		Model:  "llama3",
		Prompt: "Write a short poem about Go programming.",
		Stream: nil, // default streaming
	}
	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		if resp.Done {
			fmt.Println("\n[Generation complete]")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
