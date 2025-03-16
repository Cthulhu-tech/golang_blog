package utils

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Logger.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.SetOutput(os.Stdout)
		Logger.Warn("Не удалось открыть файл логов, используется stdout")
	}
}

func LogInfo(message string) {
	if Logger == nil {
		fmt.Println("Сообщение:", message)
		return
	}
	Logger.Info(message)
}

func LogError(err error) {
	if Logger == nil {
		fmt.Println("Logger не инициализирован! Ошибка:", err)
		return
	}
	Logger.WithFields(logrus.Fields{
		"error": err,
	}).Error("An error occurred")
}

func LogFatal(err error) {
	if Logger == nil {
		fmt.Println("Logger не инициализирован! Критическая ошибка:", err)
		os.Exit(1)
	}
	Logger.WithFields(logrus.Fields{
		"error": err,
	}).Fatal("Critical error")
}

func LogDebug(message string) {
	if Logger == nil {
		fmt.Println("Logger не инициализирован! Debug:", message)
		return
	}
	Logger.Debug(message)
}
