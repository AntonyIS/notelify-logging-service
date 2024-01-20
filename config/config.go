package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV               string
	SERVER_PORT       string
	LOGGING_TABLE     string
	SECRET_KEY        string
	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_PASSWORD string
	DEBUG             bool
	TEST              bool
}

func NewConfig() (*Config, error) {
	ENV := os.Getenv("ENV")

	if ENV == "test" {
		err := godotenv.Load("../../../.env")
		if err != nil {
			return nil, err
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
		ENV = os.Getenv("ENV")
	}

	var (
		SECRET_KEY        = os.Getenv("SECRET_KEY")
		POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		POSTGRES_USER     = "postgres"
		POSTGRES_DB       = "postgres"
		POSTGRES_HOST     = "postgres"
		POSTGRES_PORT     = "5432"
		SERVER_PORT       = "8002"
		LOGGING_TABLE     = "Logs"
		DEBUG             = false
		TEST              = false
	)

	switch ENV {
	case "production":
		TEST = false
		DEBUG = false

	case "production_test":
		TEST = true
		DEBUG = true
		LOGGING_TABLE = "TestLogs"

	case "developement_test":
		TEST = true
		DEBUG = true
		LOGGING_TABLE = "TestLogs"

	case "development":
		TEST = true
		DEBUG = true
		POSTGRES_HOST = "localhost"
		LOGGING_TABLE = "DevLogs"

	case "docker":
		TEST = true
		DEBUG = true
		LOGGING_TABLE = "DockerLogs"

	case "docker_test":
		TEST = true
		DEBUG = true
		LOGGING_TABLE = "DockerLogs"
	}

	config := Config{
		ENV:               ENV,
		SERVER_PORT:       SERVER_PORT,
		LOGGING_TABLE:     LOGGING_TABLE,
		SECRET_KEY:        SECRET_KEY,
		DEBUG:             DEBUG,
		TEST:              TEST,
		POSTGRES_DB:       POSTGRES_DB,
		POSTGRES_USER:     POSTGRES_USER,
		POSTGRES_HOST:     POSTGRES_HOST,
		POSTGRES_PORT:     POSTGRES_PORT,
		POSTGRES_PASSWORD: POSTGRES_PASSWORD,
	}

	return &config, nil
}
