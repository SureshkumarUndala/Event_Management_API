package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LogDirectoryFolder = "./logs"
)

type Log struct {
	logDir string
}

var Logger *Log

func Loggerinit() *Log {
	if _, err := os.Stat(LogDirectoryFolder); os.IsNotExist(err) {
		err := os.Mkdir(LogDirectoryFolder, 0766)
		if err != nil {
			return nil
		}
	}
	return &Log{
		logDir: LogDirectoryFolder,
	}
}

func fetchLogFile() *os.File {
	year, month, day := time.Now().Date()
	filename := "log file" + fmt.Sprintf("%v-%v-%v.log", day, int(month), year)
	filepath, _ := os.OpenFile(LogDirectoryFolder+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	return filepath
}

func (l *Log) Error() *log.Logger {
	return log.New(fetchLogFile(), "ERROR : ", log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
}
func (l *Log) Info() *log.Logger {
	return log.New(fetchLogFile(), "INFO : ", log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
}
func (l *Log) Fatal() *log.Logger {
	return log.New(fetchLogFile(), "FATAL : ", log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
}
func (l *Log) Success() *log.Logger {
	return log.New(fetchLogFile(), "SUCCESS : ", log.Ldate|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
}
