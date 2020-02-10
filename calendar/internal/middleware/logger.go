package middleware

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

const (
	Error = iota
	Warning
	Info
	Debug
)

var logFile *os.File = nil

func ConfigureLogger(file string, level int) error {
	logrus.SetLevel(getLoggerLevel(level))
	output, err := getLoggerOutput(file)
	if err != nil {
		return err
	}
	logrus.SetOutput(output)
	return nil
}

func getLoggerLevel(level int) logrus.Level {
	var logLevel logrus.Level
	switch level {
	case Error:
		logLevel = logrus.ErrorLevel
	case Debug:
		logLevel = logrus.DebugLevel
	case Warning:
		logLevel = logrus.WarnLevel
	case Info:
		logLevel = logrus.InfoLevel
	}
	return logLevel
}

func getLoggerOutput(file string) (io.Writer, error) {
	if file == "" {
		return os.Stdout, nil
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	logFile = f
	return f, nil
}

func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

func SetLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"path": r.URL.Path,
		}).Info()

		next.ServeHTTP(w, r)
	})
}
