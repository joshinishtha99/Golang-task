package logrus

import (
	"os"
	"time"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	filename := "logs/app-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	Logger.SetOutput(file)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.SetLevel(logrus.InfoLevel)

	Logger.SetOutput(os.Stdout)
}
