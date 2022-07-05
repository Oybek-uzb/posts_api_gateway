package config

import (
	"github.com/spf13/cast"
	"os"
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

	config.PostsCRUDServiceHost = cast.ToString(getOrReturnDefault("BOND_SERVICE_GRPC_HOST", "bond-service"))
	config.PostsCRUDServicePort = cast.ToInt(getOrReturnDefault("BOND_SERVICE_GRPC_PORT", 80))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
