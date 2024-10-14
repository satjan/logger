package logger

import (
	"github.com/sirupsen/logrus"
)

func Tracef(format string, args ...interface{}) {
	WithCaller().Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	WithCaller().Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	WithCaller().Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	WithCaller().Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	WithCaller().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	WithCaller().Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	WithCaller().Logf(logrus.FatalLevel, format, args...)
	Log.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	WithCaller().Logf(logrus.PanicLevel, format, args...)
}

func Trace(args ...interface{}) {
	WithCaller().Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	WithCaller().Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	WithCaller().Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	WithCaller().Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	WithCaller().Warn(args...)
}

func Error(args ...interface{}) {
	WithCaller().Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	WithCaller().Log(logrus.FatalLevel, args...)
	Log.Exit(1)
}

func Panic(args ...interface{}) {
	WithCaller().Log(logrus.PanicLevel, args...)
}

func Traceln(args ...interface{}) {
	WithCaller().Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	WithCaller().Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	WithCaller().Logln(logrus.InfoLevel, args...)
}

func Warnln(args ...interface{}) {
	WithCaller().Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	WithCaller().Warnln(args...)
}

func Errorln(args ...interface{}) {
	WithCaller().Logln(logrus.ErrorLevel, args...)
}

func Fatalln(args ...interface{}) {
	WithCaller().Logln(logrus.FatalLevel, args...)
	Log.Exit(1)
}

func Panicln(args ...interface{}) {
	WithCaller().Logln(logrus.PanicLevel, args...)
}
