package main

import (
	"io"
	"log"
	"os"
)

var (
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
	LogNet     *log.Logger
)

func InitLogs() error {

	//TODO: make data/log directory if not exists

	logFile, err := os.OpenFile("./data/log/logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)

	if err != nil {
		return err
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	LogInfo = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime)
	LogWarning = log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime)
	LogError = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime)
	LogNet = log.New(multiWriter, "NET: ", log.Ldate|log.Ltime)

	return nil
}
