package logger

import (
	"github.com/sirupsen/logrus"
)

func Tracef(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.FatalLevel, format, args...)
	Log.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logf(logrus.PanicLevel, format, args...)
}

func Trace(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Warn(args...)
}

func Error(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.FatalLevel, args...)
	Log.Exit(1)
}

func Panic(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Log(logrus.PanicLevel, args...)
}

func Traceln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.InfoLevel, args...)
}

func Warnln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Warnln(args...)
}

func Errorln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.ErrorLevel, args...)
}

func Fatalln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.FatalLevel, args...)
	Log.Exit(1)
}

func Panicln(args ...interface{}) {
	file, line := getCaller()
	Log.WithField("file", file).
		WithField("line", line).
		Logln(logrus.PanicLevel, args...)
}
