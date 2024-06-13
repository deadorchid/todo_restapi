package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(logger *log.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) PrintInfo(info string) {
	l.logger.Println(fmt.Sprintf("[INFO] %s", info))
}

func (l *Logger) PrintWarning(warning string) {
	l.logger.Println(fmt.Sprintf("[WARNING] %s", warning))
}

func (l *Logger) PrintError(err error) {
	l.logger.Fatal(fmt.Sprintf("[ERROR] %v", err))
}
