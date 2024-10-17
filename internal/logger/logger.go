package logger

import (
	"log"
	"os"
	"time"
	"path/filepath"
	"fmt"
)

type Logger struct {
	file *os.File
}
// Create new logger and log file in "logs" folder.
func New() (*Logger, error) {
	currentTime := time.Now().Format("020106")
	filename := fmt.Sprintf("%s_log.txt", currentTime)
	filePath := filepath.Join("logs", filename)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{file: file}, nil
}
// Log message to log file.
func (l *Logger) Log(message string) {
	log.SetOutput(l.file)
	log.Printf("%s %s", time.Now().Format("17/10/2024 15:04"), message)
}

func (l *Logger) Close() {
	l.file.Close()
}