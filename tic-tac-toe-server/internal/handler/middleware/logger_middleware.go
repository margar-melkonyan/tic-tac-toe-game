package middleware

import (
	"bufio"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type statusRecorder struct {
	statusCode int
	http.ResponseWriter
}

func (sr *statusRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := sr.ResponseWriter.(http.Hijacker); ok {
		return h.Hijack()
	}
	return nil, nil, errors.New("the ResponseWriter doesn't support hijacking")
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
		start := time.Now()
		next.ServeHTTP(sr, r)
		slog.Info(
			"Request info",
			slog.String("method", r.Method),
			slog.Int("status", sr.statusCode),
			slog.String("uri", r.URL.String()),
			slog.String("duration", time.Since(start).String()),
		)
	})
}
