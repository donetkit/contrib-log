package glog

import (
	"fmt"
	"log"
	"os"
	"time"
)

type DefaultLogger struct {
	logger       *log.Logger
	config       *Config
	dateFormat   string
	class        string
	logFormatter func(interface{}, bool) string
}

func NewDefaultLogger(opts ...Option) ILogger {
	hostName, _ := os.Hostname()
	cfg := &Config{
		logLevel: DEBUG,
		log2File: false,
		hostName: hostName,
		ip:       getHostIp(),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	logger := NewLoggerWith(log.New(os.Stdout, "", 0))
	logger.config = cfg
	logger.SetCustomLogFormat(defaultLogFormatter)
	return logger
}

func NewLoggerWith(log *log.Logger) *DefaultLogger {
	logger := &DefaultLogger{logger: log, dateFormat: defaultDateFormat}
	logger.SetCustomLogFormat(defaultLogFormatter)
	return logger
}

func (log *DefaultLogger) SetClass(className string) {
	log.class = className
}

func (log *DefaultLogger) SetLogLevel(level LogLevel) {
	log.config.logLevel = level
}

func (log *DefaultLogger) SetCustomLogFormat(logFormatterFunc func(logInfo interface{}, color bool) string) {
	log.logFormatter = logFormatterFunc
}

func (log *DefaultLogger) SetDateFormat(format string) {
	log.dateFormat = format
}

func (log *DefaultLogger) log(level LogLevel, format string, a ...interface{}) {
	if level < log.config.logLevel {
		return
	}
	message := format
	message = fmt.Sprintf(format, a...)

	start := time.Now()
	info := LogInfo{
		StartTime: start.Format(log.dateFormat),
		Level:     LevelString[level],
		Class:     log.class,
		Host:      log.config.hostName,
		IP:        log.config.ip,
		Message:   message,
	}
	if log.config.logColor {
		info.Level = LevelColorString[level]
	}
	log.logger.Println(log.logFormatter(info, log.config.logColor))
}

func (log *DefaultLogger) Debug(format string, a ...interface{}) {
	if DEBUG < log.config.logLevel {
		return
	}
	log.log(DEBUG, format, a...)
}

func (log *DefaultLogger) Info(format string, a ...interface{}) {
	if INFO < log.config.logLevel {
		return
	}
	log.log(INFO, format, a...)
}

func (log *DefaultLogger) Warning(format string, a ...interface{}) {
	if WARNING < log.config.logLevel {
		return
	}
	log.log(WARNING, format, a...)
}

func (log *DefaultLogger) Error(format string, a ...interface{}) {
	if ERROR < log.config.logLevel {
		return
	}
	log.log(ERROR, format, a...)
}
