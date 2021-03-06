package glog

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"sync/atomic"
)

var defaultDateFormat = "2006-01-02 15:04:05.000"

// Level type
type Level uint32

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type Logger struct {
	logger *logrus.Logger
	config *Config
}

type Config struct {
	noColors       bool
	noFieldsColors bool
	level          Level
	reportCaller   bool
	log2File       bool
	hostName       string
	ip             string
	displayFields  bool
	className      string
	fields         map[string]interface{}
	fileName       string
	maxSize        int
	maxAge         int
	maxBackups     int
	localTime      bool
	compress       bool
}

func New(opts ...Option) ILogger {
	hostName, _ := os.Hostname()
	cfg := &Config{
		level:        DebugLevel,
		reportCaller: false,
		log2File:     false,
		hostName:     hostName,
		ip:           getHostIp(),
		fileName:     "./logs/log.log", //日志文件存放目录
		maxSize:      1,                //文件大小限制,单位MB
		maxBackups:   15,               //最大保留日志文件数量
		maxAge:       7,                //日志文件保留天数
	}
	for _, opt := range opts {
		opt(cfg)
	}

	logger := logrus.New()
	logger.SetReportCaller(cfg.reportCaller)
	logger.SetLevel(logrus.Level(cfg.level))
	//logger.AddHook(NewContextHook(logrus.Level(cfg.level)))
	if cfg.log2File {
		cfg.noColors = true
		cfg.noFieldsColors = true
		logger.SetOutput(&WriteLogger{
			FileName:   cfg.fileName,   //日志文件存放目录
			MaxSize:    cfg.maxSize,    //文件大小限制,单位MB
			MaxBackups: cfg.maxBackups, //最大保留日志文件数量
			MaxAge:     cfg.maxAge,     //日志文件保留天数
		})
	} else {
		logger.SetOutput(os.Stdout)
	}
	logger.SetFormatter(&Formatter{
		FieldsOrder:     []string{"host", "ip", "class"},
		NoColors:        cfg.noColors,
		NoFieldsColors:  cfg.noFieldsColors,
		TimestampFormat: defaultDateFormat,
		HideKeys:        true,
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return path.Base(frame.File)
		},
	})
	return &Logger{logger: logger, config: cfg}
}

func (log *Logger) level() logrus.Level {
	return logrus.Level(atomic.LoadUint32((*uint32)(&log.config.level)))
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param
func (log *Logger) IsLevelEnabled(level logrus.Level) bool {
	return log.level() >= level
}

func (log *Logger) With(fields map[string]interface{}) *logrus.Entry {
	fieldsMap := make(map[string]interface{})
	//fieldsMap["prefix"] = "prefix"
	if log.config.displayFields {
		if log.config.hostName != "" {
			fieldsMap["host"] = log.config.hostName
		}
		if log.config.ip != "" {
			fieldsMap["ip"] = log.config.ip
		}
		if log.config.className != "" {
			fieldsMap["class"] = log.config.className
		}
	}
	if fields != nil {
		for k, v := range fields {
			fieldsMap[k] = v
		}
	}
	//fieldsMap["message"] = message
	return log.logger.WithFields(fieldsMap)
}

func (log *Logger) logf(level logrus.Level, format string, args ...interface{}) {
	if log.IsLevelEnabled(level) {
		log.With(log.config.fields).Logf(level, format, args)
	}
}

func (log *Logger) log(level logrus.Level, args ...interface{}) {
	if log.IsLevelEnabled(level) {
		log.With(log.config.fields).Log(level, args)
	}
}

func (log *Logger) logIn(level logrus.Level, args ...interface{}) {
	if log.IsLevelEnabled(level) {
		log.With(log.config.fields).Logln(level, args)
	}
}

func (log *Logger) AddHook(hook logrus.Hook) {
	log.logger.AddHook(hook)
}

func (log *Logger) WithClassName(className string) {
	log.config.className = className
}

func (log *Logger) SetFormatter(formatter logrus.Formatter) {
	log.logger.SetFormatter(formatter)
}

func (log *Logger) SetReportCaller(reportCaller bool) {
	log.logger.SetReportCaller(reportCaller)
}

func (log *Logger) SetNoLock() {
	log.logger.SetNoLock()
}

func (log *Logger) SetLevel(level Level) {
	log.config.level = level
	log.logger.SetLevel(logrus.Level(level))
}

func (log *Logger) WithField(key string, value interface{}) *logrus.Entry {
	fieldsMap := make(map[string]interface{})
	fieldsMap[key] = value
	return log.With(fieldsMap)
}

func (log *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return log.With(fields)
}

func (log *Logger) WithError(err error) *logrus.Entry {
	return log.With(log.config.fields).WithError(err)
}

func (log Logger) Debugf(format string, args ...interface{}) {
	log.logf(logrus.DebugLevel, format, args)
}
func (log Logger) Infof(format string, args ...interface{}) {
	log.logf(logrus.InfoLevel, format, args)
}
func (log *Logger) Printf(format string, args ...interface{}) {
	log.logger.Printf(format, args)
}
func (log Logger) Warnf(format string, args ...interface{}) {
	log.logf(logrus.WarnLevel, format, args)
}
func (log Logger) Warningf(format string, args ...interface{}) {
	log.Warnf(format, args)
}
func (log Logger) Errorf(format string, args ...interface{}) {
	log.logf(logrus.ErrorLevel, format, args)
}
func (log Logger) Fatalf(format string, args ...interface{}) {
	log.logf(logrus.FatalLevel, format, args)
}
func (log Logger) Panicf(format string, args ...interface{}) {
	log.logf(logrus.PanicLevel, format, args)
}

func (log Logger) Debug(args ...interface{}) {
	log.log(logrus.DebugLevel, args)
}
func (log Logger) Info(args ...interface{}) {
	log.log(logrus.InfoLevel, args)
}
func (log *Logger) Print(args ...interface{}) {
	log.logger.Print(args)
}
func (log Logger) Warn(args ...interface{}) {
	log.log(logrus.WarnLevel, args)
}
func (log Logger) Warning(args ...interface{}) {
	log.Warn(args)
}
func (log Logger) Error(args ...interface{}) {
	log.log(logrus.ErrorLevel, args)
}
func (log Logger) Fatal(args ...interface{}) {
	log.log(logrus.FatalLevel, args)
}
func (log Logger) Panic(args ...interface{}) {
	log.log(logrus.PanicLevel, args)
}

func (log Logger) Debugln(args ...interface{}) {
	log.logIn(logrus.DebugLevel, args)
}
func (log Logger) Infoln(args ...interface{}) {
	log.logIn(logrus.InfoLevel, args)
}
func (log *Logger) Println(args ...interface{}) {
	log.logger.Println(args)
}
func (log Logger) Warnln(args ...interface{}) {
	log.logIn(logrus.WarnLevel, args)
}
func (log Logger) Warningln(args ...interface{}) {
	log.Warnln(args)
}
func (log Logger) Errorln(args ...interface{}) {
	log.logIn(logrus.ErrorLevel, args)
}
func (log Logger) Fatalln(args ...interface{}) {
	log.logIn(logrus.FatalLevel, args)
}
func (log Logger) Panicln(args ...interface{}) {
	log.logIn(logrus.PanicLevel, args)
}
