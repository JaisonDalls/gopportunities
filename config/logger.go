package config

import (
	"io"
	"log"
	"os"
)

// Logger is ...
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// NewLogger is ...
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Debug is ...
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Panicln(v...)
}

// INFO is ...
func (l *Logger) INFO(v ...interface{}) {
	l.info.Panicln(v...)
}

// Warning is ...
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Panicln(v...)
}

// Error is ...
func (l *Logger) Error(v ...interface{}) {
	l.err.Panicln(v...)
}

// Debugf is ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Panicf(format, v...)
}

// Infof is ...
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Panicf(format, v...)
}

// Warningf is ...
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Panicf(format, v...)
}

// Errorf is ...
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Panicf(format, v...)
}
