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

func ConfigLogger(level logrus.Level, target io.Writer) {
	logrus.SetLevel(level)
	logrus.SetOutput(target)
}

type Fields map[string]interface{}

func LogInfo(fields Fields, message string) {
	logrus.WithFields(logrus.Fields(fields)).Info(message)
}

func LogFatal(fields Fields, message string) {
	logrus.WithFields(logrus.Fields(fields)).Fatal(message)
}

func LogWarn(fields Fields, message string) {
	logrus.WithFields(logrus.Fields(fields)).Warn(message)
}

func LogDebug(fields Fields, message string) {
	logrus.WithFields(logrus.Fields(fields)).Debug(message)
}

func LogError(fields Fields, message string) {
	logrus.WithFields(logrus.Fields(fields)).Error(message)
}

func LogPanic(fields Fields, message string) {
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
