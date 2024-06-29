package dnutlogger

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	SUCCESS
	WARN
	ERROR
)

var appliedLevel LogLevel = INFO

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DBG"
	case INFO:
		return "INF"
	case WARN:
		return "WAR"
	case SUCCESS:
		return "SUC"
	case ERROR:
		return "ERR"
	default:
		return "UKN"
	}
}

var useColor = true

func UseColor(enabled bool) {
	useColor = enabled
}

func log(level LogLevel, format string, a ...any) {
	if level < appliedLevel {
		return
	}
	var colorSet string
	var colorReset string
	if useColor {
		colorReset = "\033[0m"
		switch level {
		case DEBUG:
			colorSet = "\033[0;90m"
		case INFO:
			colorSet = "\033[0;94m"
		case WARN:
			colorSet = "\033[0;93m"
		case SUCCESS:
			colorSet = "\033[0;92m"
		case ERROR:
			colorSet = "\033[0;91m"
		}
	}
	fmt.Printf("%s%-3s [%s]%s %s\n", colorSet, level, time.Now().Format("2006-01-02 15:04:05"), colorReset, fmt.Sprintf(format, a...))
}

func FromString(level string) LogLevel {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG", "DBG":
		return DEBUG
	case "INFO", "INF":
		return INFO
	case "WARN", "WAR":
		return WARN
	case "SUCCESS", "SUC":
		return SUCCESS
	case "ERROR", "ERR":
		return ERROR
	default:
		return INFO
	}
}

func SetMinLevel(level LogLevel) {
	appliedLevel = level
}

func Debugf(format string, a ...any) {
	log(DEBUG, format, a...)
}

func Debug(a ...any) {
	log(DEBUG, "%+v", a)
}

func Infof(format string, a ...any) {
	log(INFO, format, a...)
}

func Info(a ...any) {
	log(INFO, "%+v", a)
}

func Warnf(format string, a ...any) {
	log(WARN, format, a...)
}

func Warn(a ...any) {
	log(WARN, "%+v", a)
}

func Successf(format string, a ...any) {
	log(SUCCESS, format, a...)
}

func Success(a ...any) {
	log(SUCCESS, "%+v", a)
}

func Err(exit bool, err error) {
	if err == nil {
		return
	}
	Errorf(exit, "%+v", err)
}

func Errorf(exit bool, format string, a ...any) {
	debug.PrintStack()
	log(ERROR, format, a...)
	if exit {
		os.Exit(1)
	}
}

func Error(exit bool, a ...any) {
	debug.PrintStack()
	log(ERROR, "%+v", a)
	if exit {
		os.Exit(1)
	}
}
