package logger

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

type RequestID struct{}

var id RequestID

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestid := uuid.NewString()
		ctx := context.WithValue(r.Context(), "id", requestid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func GetRequestID(ctx context.Context) string {
	if val, ok := ctx.Value("").(string); ok {
		return val
	}
	return ""
}
