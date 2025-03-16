package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		LogFatal(fmt.Errorf("Ошибка загрузки .env файла: %v", err))
	}

	LogInfo("Файл .env успешно загружен")
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		LogFatal(fmt.Errorf("Не установлена переменная окружения %s", key))
	}

	return value
}

func GetPostgresql() string {
	host := GetEnv("POSTGRES_HOST")
	port := GetEnv("POSTGRES_PORT")
	user := GetEnv("POSTGRES_USER")
	password := GetEnv("POSTGRES_PASSWORD")
	dbname := GetEnv("POSTGRES_DB")

	LogInfo(fmt.Sprintf("Подключение к БД с параметрами: HOST=%s, PORT=%s, USER=%s, DB=%s", host, port, user, dbname))

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)

	return connectionString
}
