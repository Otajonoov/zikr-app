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
	JwtSecretKet     string
}

type JwtConfig struct {
	JWTSecret string
}

func NewConfig() *JwtConfig {
	var config JwtConfig

	config.JWTSecret = getEnv("JWT_SECRET_KEY", "")

	return &config
}

func Load() Config {
	c := Config{}
	c.HttpPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", "50055"))
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "info"))

	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "zikr_app"))
	c.JwtSecretKet = cast.ToString(GetOrReturnDefault("JWT_SECRET_KEY", "WHic3i9cGl"))

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

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}

//package config
//
//import (
//	"github.com/ilyakaznacheev/cleanenv"
//	"log"
//	"os"
//	"time"
//)
//
//type Config struct {
//	Env         string `yaml:"env" env-default:"local"`
//	Environment string `yaml:"environment" env-default:"develop"`
//	LogLevel    string `yaml:"logLevel" env-default:"info"`
//	DatabaseUrl string `yaml:"database_url" env-required:"true"`
//	Postgres
//	HTTPServer
//}
//
//type HTTPServer struct {
//	Address     string        `yaml:"address" env-default:"localhost:50055"`
//	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
//	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
//	User        string        `yaml:"user" env-required:"true"`
//	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
//}
//
//type Postgres struct {
//	PostgresHost     string `yaml:"postgres_host" env-required:"true"`
//	PostgresPort     string `yaml:"postgres_port" env-required:"true"`
//	PostgresUser     string `yaml:"postgres_user" env-required:"true"`
//	PostgresPassword string `yaml:"postgres_password" env-required:"true"`
//	PostgresDatabase string `yaml:"postgres_database" env-required:"true"`
//}
//
//func MustLoad() *Config {
//	configPath := os.Getenv("CONFIG_PATH")
//	if configPath == "" {
//		log.Fatal("CONFIG_PATH is not set")
//	}
//
//	// check if file exists
//	if _, err := os.Stat(configPath); os.IsNotExist(err) {
//		log.Fatalf("config file does not exist: %s", configPath)
//	}
//
//	var cfg Config
//
//	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
//		log.Fatalf("cannot read config: %s", err)
//	}
//
//	return &cfg
//}
//
//func generateDatabaseUrl(c Config) string {
//	return "postgres://" + c.PostgresUser + ":" + c.PostgresPassword +
//		"@" + c.PostgresHost + ":" + c.PostgresPort + "/" + c.PostgresDatabase +
//		"?sslmode=disable"
//}
//
//func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
//	_, exists := os.LookupEnv(key)
//	if exists {
//		return os.Getenv(key)
//	}
//	return defaultValue
//}
