package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	GATEWAY_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	MONGO_URI                  string
	MONGO_DB_NAME              string
	MONGO_CHAT_COLLECTION_NAME string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.GATEWAY_PORT = cast.ToString(coalesce("GATEWAY_PORT", ":2030"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "portgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "sport_db"))

	config.MONGO_URI = cast.ToString(coalesce("MONGO_URI", "mongodb://localhost:27017"))
	config.MONGO_DB_NAME = cast.ToString(coalesce("MONGO_DB_NAME", "sport"))
	config.MONGO_CHAT_COLLECTION_NAME = cast.ToString(coalesce("MONGO_CHAT_COLLECTION_NAME", "chat"))

	return &config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return defaultValue
}