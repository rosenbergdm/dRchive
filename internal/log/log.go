package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetOutput(os.Stdout)
}

func Config(level Level, target io.Writer) {
	logrus.SetLevel(logrus.Level(level))
	logrus.SetOutput(target)
}

type Fields map[string]interface{}

type Level uint64

const (
	PanicLevel Level = Level(logrus.PanicLevel)
	FatalLevel Level = Level(logrus.FatalLevel)
	ErrorLevel Level = Level(logrus.ErrorLevel)
	WarnLevel  Level = Level(logrus.WarnLevel)
	InfoLevel  Level = Level(logrus.InfoLevel)
	DebugLevel Level = Level(logrus.DebugLevel)
	TraceLevel Level = Level(logrus.TraceLevel)
)

func Info(message string, fields Fields) {
	logrus.WithFields(logrus.Fields(fields)).Info(message)
}

func Fatal(message string, fields Fields) {
	logrus.WithFields(logrus.Fields(fields)).Fatal(message)
}

func Warn(message string, fields Fields) {
	logrus.WithFields(logrus.Fields(fields)).Warn(message)
}

func Debug(message string, fields Fields) {
	logrus.WithFields(logrus.Fields(fields)).Debug(message)
}

func Error(message string, fields Fields) {
	logrus.WithFields(logrus.Fields(fields)).Error(message)
}

func Panic(message string, fields Fields) {
	logrus.WithFields(logrus.Fields(fields)).Panic(message)
}

func LogInfof(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Infof(message, args...)
}

func LogFatalf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Fatalf(message, args...)
}

func LogWarnf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Warnf(message, args...)
}

func LogDebugf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Debugf(message, args...)
}

func LogErrorf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Errorf(message, args...)
}

func LogPanicf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Panicf(message, args...)
}
