package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production

	PostsCRUDServiceHost string
	PostsCRUDServicePort int

	HttpPort string
}

func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	config.PostsCRUDServiceHost = cast.ToString(getOrReturnDefault("POSTS_CRUD_SERVICE_HOST", "localhost"))
	config.PostsCRUDServicePort = cast.ToInt(getOrReturnDefault("POSTS_CRUD_SERVICE_PORT", 8081))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
