package mlogger

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var logger1 *logger
var once sync.Once

func GetInstance() *logger {
	once.Do(func() {
		logger1 = createLogger("mlogger.log")
	})
	return logger1
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "gomessenger ", log.Lshortfile),
	}
}
