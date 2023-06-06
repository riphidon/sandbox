package logs

import (
	"io"
	"log"
	"os"
)

const (
	start string = "App Start |"
	debug string = "Debug Log | "
	info  string = "Info Log  | "
	warn  string = "Warn Log  | "
)

var logFile *os.File

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Startf(format string, args ...interface{})
	FatalF(format string, args ...interface{})
}

type AppLogger struct {
	logger *log.Logger
}

func init() {
	file, err := os.OpenFile("general-log.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	logFile = file
}

func NewAppLogger() *AppLogger {
	return &AppLogger{logger: log.New(io.MultiWriter(os.Stdout, logFile), "", 0)}
}

func (l *AppLogger) Debugf(format string, args ...interface{}) {
	l.logger.SetPrefix(debug)
	l.logger.Printf(format, args...)
}

func (l *AppLogger) Infof(format string, args ...interface{}) {
	l.logger.SetPrefix(info)
	l.logger.Printf(format, args...)
}

func (l *AppLogger) Startf(format string, args ...interface{}) {
	l.logger.SetPrefix(start)
	l.logger.Printf(format, args...)
}

func (l *AppLogger) Fatalf(format string, args ...interface{}) {
	l.logger.SetPrefix(start)
	l.logger.Fatalf(format, args...)
}

func (l *AppLogger) SetMode(prefix string) {
	l.logger.SetPrefix(prefix)
}
