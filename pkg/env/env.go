package env

import (
	constants "01-basic-api-calling/contants"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (err error) {
	if err = godotenv.Load(".env"); err != nil {
		return
	}

	return
}

func LoadAPIKey() (apiKey string, err error) {
	apiKey = os.Getenv(constants.XAIAPIKeyEnv)
	if len(apiKey) == 0 {
		err = errors.New("API Key is not set.")
		return
	}

	return
}
