package dnutlogger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelSuccess
	LevelWarning
	LevelError
	LevelFatal
)

var levelPrefix = [...]string{
	LevelDebug:   "DBG",
	LevelInfo:    "INF",
	LevelSuccess: "SUC",
	LevelWarning: "WAR",
	LevelError:   "ERR",
	LevelFatal:   "FTL",
}

func (l LogLevel) Color() string {
	switch l {
	case LevelDebug:
		return "\033[0;90m"
	case LevelInfo:
		return "\033[0;94m"
	case LevelSuccess:
		return "\033[0;92m"
	case LevelWarning:
		return "\033[0;93m"
	case LevelError:
		return "\033[0;91m"
	case LevelFatal:
		return "\033[0;31m"
	}

	return ""
}

const ColorReset = "\033[0m"

type FormattingOptions struct {
	UseColor        bool
	ShowTimestamp   bool
	TimestampFormat string
}

type Logger struct {
	level      LogLevel
	writer     io.Writer
	outputFunc func(level LogLevel, msg string)
}

func NewLogger(level LogLevel, writer io.Writer, opts FormattingOptions) *Logger {
	logger := &Logger{
		level:  level,
		writer: writer,
	}

	formatStr := ""
	if opts.UseColor {
		formatStr += level.Color()
	}
	formatStr += "%-3s "
	if opts.ShowTimestamp {
		formatStr += "[%s]"
	}
	if opts.UseColor {
		formatStr += ColorReset
	}
	formatStr += " %s\n"

	logger.outputFunc = func(level LogLevel, msg string) {
		if opts.ShowTimestamp {
			timestamp := time.Now().Format(opts.TimestampFormat)
			fmt.Fprintf(writer, formatStr, levelPrefix[level], timestamp, msg)
		} else {
			fmt.Fprintf(writer, formatStr, levelPrefix[level], msg)
		}
	}

	return logger
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
}

func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	message := fmt.Sprintf(format, v...)

	l.outputFunc(level, message)
}

func (l *Logger) Debug(v ...interface{}) {
	l.log(LevelDebug, "%+v", v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(LevelDebug, format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.log(LevelInfo, "%+v", v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(LevelInfo, format, v...)
}

func (l *Logger) Success(v ...interface{}) {
	l.log(LevelSuccess, "%+v", v...)
}

func (l *Logger) Successf(format string, v ...interface{}) {
	l.log(LevelSuccess, format, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.log(LevelWarning, "%+v", v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.log(LevelWarning, format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.log(LevelError, "%+v", v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(LevelError, format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.log(LevelFatal, "%+v", v...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(LevelFatal, format, v...)
	os.Exit(1)
}

var defaultLogger = NewLogger(LevelInfo, os.Stdout, FormattingOptions{
	UseColor:        true,
	ShowTimestamp:   true,
	TimestampFormat: time.RFC822,
})

func SetLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

func SetWriter(writer io.Writer) {
	defaultLogger.SetWriter(writer)
}

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Success(v ...interface{}) {
	defaultLogger.Success(v...)
}

func Successf(format string, v ...interface{}) {
	defaultLogger.Successf(format, v...)
}

func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}
