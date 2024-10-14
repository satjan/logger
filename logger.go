package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"regexp"
	"runtime"
)

var Log = logrus.New()

// CustomFormatter is a logger formatter that masks sensitive fields
type CustomFormatter struct {
	FieldNames []string
	logrus.JSONFormatter
}

// Format formats the log entry and masks sensitive fields
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	for _, fieldName := range f.FieldNames {
		re := regexp.MustCompile(fieldName)
		entry.Message = re.ReplaceAllString(entry.Message, `********:"********"`)
	}
	return f.JSONFormatter.Format(entry)
}

func Init(filename string) {
	// Create a custom formatter with sensitive fields to mask
	formatter := &CustomFormatter{
		FieldNames: []string{
			`password:"[^"]+"`,
			`passwordMatch:"[^"]+"`,
			`token:"[^"]+"`,
			`secret:"[^"]+"`,
			`tranPassword:"[^"]+"`,
			`"password":"[^"]+"`,
			`"passwordMatch":"[^"]+"`,
			`"token":"[^"]+"`,
			`"secret":"[^"]+"`,
			`"tranPassword":"[^"]+"`,
			`<Pwd>[^<]+</Pwd>`,
		},
		JSONFormatter: logrus.JSONFormatter{
			DataKey: "data",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "log.level",
				logrus.FieldKeyMsg:   "message",
			},
		},
	}

	Log.SetFormatter(formatter)

	Log.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,    // Max size in megabytes before log rotation occurs
		MaxBackups: 5,    // Max number of old log files to keep
		MaxAge:     30,   // Max number of days to retain old log files
		Compress:   true, // Whether to compress the rotated log files
	})
}

// Helper function to get the true caller's file, line, and function name
func getCaller() (string, int, string) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return "", 0, ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	return file, line, funcName
}

// WithCaller adds file, line, and function name information to the log entry
func WithCaller() *logrus.Entry {
	file, line, funcName := getCaller()
	return Log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"func": funcName,
	})
}
