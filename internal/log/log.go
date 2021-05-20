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

func Infof(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Infof(message, args...)
}

func Fatalf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Fatalf(message, args...)
}

func Warnf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Warnf(message, args...)
}

func Debugf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Debugf(message, args...)
}

func Errorf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Errorf(message, args...)
}

func Panicf(fields Fields, message string, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Panicf(message, args...)
}
