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

func LogInfo(fields logrus.Fields, message string) {
	logrus.WithFields(fields).Info(message)
}

func LogFatal(fields logrus.Fields, message string) {
	logrus.WithFields(fields).Fatal(message)
}

func LogWarn(fields logrus.Fields, message string) {
	logrus.WithFields(fields).Warn(message)
}

func LogDebug(fields logrus.Fields, message string) {
	logrus.WithFields(fields).Debug(message)
}

func LogError(fields logrus.Fields, message string) {
	logrus.WithFields(fields).Error(message)
}

func LogPanic(fields logrus.Fields, message string) {
	logrus.WithFields(fields).Panic(message)
}

func LogInfof(fields logrus.Fields, message string, args ...interface{}) {
	logrus.WithFields(fields).Infof(message, args...)
}

func LogFatalf(fields logrus.Fields, message string, args ...interface{}) {
	logrus.WithFields(fields).Fatalf(message, args...)
}

func LogWarnf(fields logrus.Fields, message string, args ...interface{}) {
	logrus.WithFields(fields).Warnf(message, args...)
}

func LogDebugf(fields logrus.Fields, message string, args ...interface{}) {
	logrus.WithFields(fields).Debugf(message, args...)
}

func LogErrorf(fields logrus.Fields, message string, args ...interface{}) {
	logrus.WithFields(fields).Errorf(message, args...)
}

func LogPanicf(fields logrus.Fields, message string, args ...interface{}) {
	logrus.WithFields(fields).Panicf(message, args...)
}
