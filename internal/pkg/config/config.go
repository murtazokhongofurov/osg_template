package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	HttpPort         string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error while connected .env")
	}
	c := Config{}

	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "port"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "databasename"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "host"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "port"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "username"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
