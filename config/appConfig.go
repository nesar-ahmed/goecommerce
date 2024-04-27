package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	fmt.Println(os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	fmt.Println(httpPort)
	if len(httpPort) > 1 {
		return AppConfig{}, errors.New("env variable not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}