package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const (
	Error = iota
	Warning
	Info
	Debug
)

func ConfigureLogger(file string, level int) {

	var logLevel logrus.Level
	switch level {
	case Error:
		logLevel = logrus.ErrorLevel
	case Debug:
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
	logrus.SetOutput(os.Stdout)
}

func SetLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Debug(r.URL.Path + " " + time.Now().String())
		next.ServeHTTP(w, r)
	})
}
