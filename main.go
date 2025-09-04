package main

import (
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


var (
	list   []int
	mutex  sync.Mutex
	Logger = logrus.New()
)

func initLogger() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	filename := "logs/app-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, file)

	Logger.SetOutput(mw)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.SetLevel(logrus.InfoLevel)
}


func handleNumber(c *gin.Context) {
	var input struct {
		Number int `json:"number"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		Logger.Error("Invalid input")
		return
	}

	number := input.Number
	if number == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "zero not allowed"})
		Logger.Error("Zero not allowed")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if len(list) == 0 {
		list = append(list, number)
		Logger.Infof("List was empty, appended: %d", number)
		c.JSON(http.StatusOK, gin.H{"updated_list": list})
		return
	}

	if (list[0] > 0 && number > 0) || (list[0] < 0 && number < 0) {
		list = append(list, number)
		Logger.Infof("Same sign, appended %d → %v", number, list)
	} else {
		toReduce := abs(number)
		Logger.Infof("Opposite sign, reducing %d from %v", toReduce, list)

		for toReduce > 0 && len(list) > 0 {
			if abs(list[0]) > toReduce {
				if list[0] > 0 {
					list[0] -= toReduce
				} else {
					list[0] += toReduce
				}
				toReduce = 0
			} else {
				toReduce -= abs(list[0])
				list = list[1:]
			}
		}

		if toReduce > 0 {
			list = append(list, number/abs(number)*toReduce)
		}
		Logger.Infof("After reduction → %v", list)
	}

	c.JSON(http.StatusOK, gin.H{"updated_list": list})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	initLogger()
	Logger.Info("Starting application...")

	r := gin.Default()
	r.GET("/number", func(c *gin.Context) {
		c.JSON(200, gin.H{"current_list": list})
	})
	r.POST("/number", handleNumber)
	r.Run(":8080")
}
