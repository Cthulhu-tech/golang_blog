package utils

import (
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
	}
}

func LogInfo(message string) {
	Logger.Info(message)
}

func LogError(err error) {
	Logger.WithFields(logrus.Fields{
		"error": err,
	}).Error("An error occurred")
}

func LogFatal(err error) {
	Logger.WithFields(logrus.Fields{
		"error": err,
	}).Fatal("Critical error")
}

func LogDebug(message string) {
	Logger.Debug(message)
}
