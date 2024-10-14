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

// Helper function to get the caller's file and line number
func getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(2) // Adjust the caller depth based on the function call stack
	if !ok {
		return "", 0
	}
	return file, line
}
