package middleware

import (
	"net/http"

	core_logging "github.com/Lens-Platform/Platform/src/libraries/core/core-logging/json"
)

type LoggingMiddleware struct {
}

// NewLoggingMiddleware returns a new instance of the logging middleware
func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

// LoggingMiddleware runs the logging middleware
func (m *LoggingMiddleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		core_logging.JSONLogger.Info(
			"request started",
			"proto", r.Proto,
			"uri", r.RequestURI,
			"method", r.Method,
			"remote", r.RemoteAddr,
			"user-agent", r.UserAgent(),
		)
		next.ServeHTTP(w, r)
	})
}
