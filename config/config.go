package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

var Cfg Config

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
	}
	Database struct {
		Username string `envconfig:"DB_USERNAME"`
		Password string `envconfig:"DB_PASSWORD"`
		Host     string `envconfig:"DB_HOST"`
		Port     string `envconfig:"DB_PORT"`
		Name     string `envconfig:"DB_NAME"`
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	if err := envconfig.Process("", &Cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
