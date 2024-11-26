package helper

import (
	"log"
	"os"
)

var (
	InfoLogger  = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger  = log.New(os.Stdout, "WARN", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(msg string) {
	InfoLogger.Println(msg)
}
func Error(msg string) {
	ErrorLogger.Println(msg)
}
func Warn(msg string) {
	WarnLogger.Println(msg)
}
