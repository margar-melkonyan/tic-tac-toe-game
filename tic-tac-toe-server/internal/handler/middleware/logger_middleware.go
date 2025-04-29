package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

type statusRecorder struct {
	statusCode int
	http.ResponseWriter
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sr := &statusRecorder{
			statusCode:     200,
			ResponseWriter: w,
		}
		next.ServeHTTP(sr, r)
		slog.Info(fmt.Sprintf("[%v %v] %v", r.Method, sr.statusCode, r.URL))
	})
}
