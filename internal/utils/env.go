package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		LogInfo(fmt.Sprintf("Error loading .env file: %v", err))
	}

	LogInfo(fmt.Sprintf("Successfully loaded .env file"))
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		LogInfo(fmt.Sprintf("Environment variable %s is required but not set", key))
	}

	LogInfo(fmt.Sprintf("Environment variable %s received", key))
	return value
}

func GetPostgresql() string {
	host := GetEnv("POSTGRES_HOST")
	port := GetEnv("POSTGRES_PORT")
	user := GetEnv("POSTGRES_USER")
	password := GetEnv("POSTGRES_PASSWORD")
	dbname := GetEnv("POSTGRES_DB")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	return connectionString
}
