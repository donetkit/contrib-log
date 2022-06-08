package glog

import (
	"github.com/sirupsen/logrus"
)

type LoggerEntry struct {
	logger *logrus.Logger
}

func (entry *LoggerEntry) Trace(args ...interface{}) {
	entry.logger.Trace(args...)
}

func (entry *LoggerEntry) Debug(args ...interface{}) {
	entry.logger.Debug(args...)
}

func (entry *LoggerEntry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *LoggerEntry) Info(args ...interface{}) {
	entry.logger.Info(args...)
}

func (entry *LoggerEntry) Warn(args ...interface{}) {
	entry.logger.Warn(args...)
}

func (entry *LoggerEntry) Warning(args ...interface{}) {
	entry.Warn(args...)
}

func (entry *LoggerEntry) Error(args ...interface{}) {
	entry.logger.Error(args...)
}

func (entry *LoggerEntry) Fatal(args ...interface{}) {
	entry.logger.Fatal(args...)
	entry.logger.Exit(1)
}

func (entry *LoggerEntry) Panic(args ...interface{}) {
	entry.logger.Panic(args...)
}

func (entry *LoggerEntry) Tracef(format string, args ...interface{}) {
	entry.logger.Tracef(format, args...)
}

func (entry *LoggerEntry) Debugf(format string, args ...interface{}) {
	entry.logger.Debugf(format, args...)
}

func (entry *LoggerEntry) Infof(format string, args ...interface{}) {
	entry.logger.Infof(format, args...)
}

func (entry *LoggerEntry) Printf(format string, args ...interface{}) {
	entry.Infof(format, args...)
}

func (entry *LoggerEntry) Warnf(format string, args ...interface{}) {
	entry.logger.Warnf(format, args...)
}

func (entry *LoggerEntry) Warningf(format string, args ...interface{}) {
	entry.logger.Warningf(format, args...)
}

func (entry *LoggerEntry) Errorf(format string, args ...interface{}) {
	entry.logger.Errorf(format, args...)
}

func (entry *LoggerEntry) Fatalf(format string, args ...interface{}) {
	entry.logger.Fatalf(format, args...)
	entry.logger.Exit(1)
}

func (entry *LoggerEntry) Panicf(format string, args ...interface{}) {
	entry.logger.Panicf(format, args...)
}

func (entry *LoggerEntry) Traceln(args ...interface{}) {
	entry.logger.Traceln(args...)
}

func (entry *LoggerEntry) Debugln(args ...interface{}) {
	entry.logger.Debugln(args...)
}

func (entry *LoggerEntry) Infoln(args ...interface{}) {
	entry.logger.Infoln(args...)
}

func (entry *LoggerEntry) Println(args ...interface{}) {
	entry.logger.Println(args...)
}

func (entry *LoggerEntry) Warnln(args ...interface{}) {
	entry.logger.Warnln(args...)
}

func (entry *LoggerEntry) Warningln(args ...interface{}) {
	entry.logger.Warningln(args...)
}

func (entry *LoggerEntry) Errorln(args ...interface{}) {
	entry.logger.Errorln(args...)
}

func (entry *LoggerEntry) Fatalln(args ...interface{}) {
	entry.logger.Fatalln(args...)
	entry.logger.Exit(1)
}

func (entry *LoggerEntry) Panicln(args ...interface{}) {
	entry.logger.Panicln(args...)
}
