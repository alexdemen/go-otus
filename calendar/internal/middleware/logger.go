package middleware

import "net/http"

type Logger struct {
}

func (l Logger) SetLogger(next http.Handler) http.Handler {
	return next
}
