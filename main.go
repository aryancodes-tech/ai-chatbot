package main

import (
	"context"
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

const (
	xAIAPIKeyEnv = "XAI_API_KEY"
	// xAIBaseURL is the OpenAI-compatible inference endpoint.
	xAIBaseURL = "https://api.x.ai/v1"
	grokModel  = "grok-4.3"
)

func loadAPIKey() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	apiKey := os.Getenv(xAIAPIKeyEnv)
	if len(apiKey) == 0 {
		return "", errors.New("XAI_API_KEY is not set")
	}

	return apiKey, nil
}

func main() {
	apiKey, err := loadAPIKey()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(xAIBaseURL),
	)

	question := "whats my name?"

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(question)},
		Model: grokModel,
	})
	if err != nil {
		panic(err)
	}

	println(resp.OutputText())
}
