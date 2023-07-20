package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	ServerHost string
	HTTPPort   string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     int

	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10
	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "localhost:"))
	cfg.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", "8080"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "crud"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "123123"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
