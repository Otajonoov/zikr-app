package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	HttpPort         string
	Environment      string
	LogLevel         string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	DatabaseUrl      string
}

func Load() Config {
	c := Config{}
	c.HttpPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", "5005"))
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "info"))

	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "zikr_app"))

	// Update the DatabaseUrl using the provided information
	c.DatabaseUrl = cast.ToString(GetOrReturnDefault("DATABASE_URL", generateDatabaseUrl(c)))
	return c
}

func generateDatabaseUrl(c Config) string {
	return "postgres://" + c.PostgresUser + ":" + c.PostgresPassword +
		"@" + c.PostgresHost + ":" + c.PostgresPort + "/" + c.PostgresDatabase +
		"?sslmode=disable"
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
