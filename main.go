package main

import (
	constants "01-basic-api-calling/contants"
	"01-basic-api-calling/pkg/env"
	"context"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	if err := env.LoadEnv(); err != nil {
		panic(err)
	}

	apiKey, err := env.LoadAPIKey()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(constants.XAIBaseURL),
	)

	question := "whats my name?"

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(question)},
		Model: constants.GrokModel,
	})
	if err != nil {
		panic(err)
	}

	println(resp.OutputText())
}
