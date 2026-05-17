package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type ResponseWiterStatus struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWiterStatus) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
func LoggerMiddleware(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			wr := &ResponseWiterStatus{ResponseWriter: w, status: http.StatusOK}
			reqID := GetRequestID(r.Context())

			next.ServeHTTP(wr, r)

			switch {
			case wr.status >= 500:
				log.Error("ошибка сервера",
					zap.String("request_id", reqID),
					zap.Int("status", wr.status),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("agent", r.UserAgent()),
					zap.Duration("duration", time.Since(start)),
				)
			case wr.status >= 400:
				log.Warn("ошибка сервера",
					zap.String("request_id", reqID),
					zap.Int("status", wr.status),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("agent", r.UserAgent()),
					zap.Duration("duration", time.Since(start)),
				)
			default:
				log.Info("успешный запрос",
					zap.String("request_id", reqID),
					zap.Int("status", wr.status),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("agent", r.UserAgent()),
					zap.Duration("duration", time.Since(start)),
				)
			}
		})
	}
}
