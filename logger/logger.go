package logger

import (
	"context"

	"deluxe-gin/monitor"
	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	return logrus.StandardLogger()
}

func GetLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Warning(args ...interface{}) {
	logrus.Warning(args...)
}

func Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}

func SetTraceID(ctx context.Context) *logrus.Entry {
	traceID := monitor.GetTraceID(ctx)
	return WithField("trace_id", traceID)
}
