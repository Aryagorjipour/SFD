package loggerInfoLogger

import (
	"log"
	"os"
)

var (
	// InfoLogger is a logger for informational messages
	InfoLogger *log.Logger
	// ErrorLogger is a logger for error messages
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an informational message
func Info(msg string) {
	InfoLogger.Println(msg)
}

// Error logs an error message
func Error(msg string) {
	ErrorLogger.Println(msg)
}
