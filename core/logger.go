package core

import (
	"os"
	"log"
)

var (
	LogInf, LogWar, LogErr *log.Logger
	logFile *os.File
	err error
)

func init() {
	logFile, err = os.OpenFile("core.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
		return
	}
	var flag int = log.Ldate | log.Ltime | log.Lshortfile
	LogInf = log.New(logFile, "INFO: ", flag)
	LogErr = log.New(logFile, "ERROR: ", flag)
	LogWar = log.New(logFile, "WARNING: ", flag)
}

func LogClose() {
	logFile.Close()
}