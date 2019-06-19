package logger

import (
	"log"
	"os"
)

var logger *log.Logger

// NewLog - creates new logger
func NewLog() *log.Logger {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger = log.New(logFile, "", log.LstdFlags)
	return logger
}
