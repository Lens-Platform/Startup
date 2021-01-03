package middleware

import (
	"net/http"
)

type VersionrMw struct {
	version string
}

// NewVersionMw returns a new instance of the version errors middleware
func NewVersionMw(version string) *VersionrMw {
	return &VersionrMw{}
}

// VersionMiddleware runs the version middleware
func (mw *VersionrMw) VersionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-API-Version", mw.version)
		r.Header.Set("X-API-Revision", mw.version)

		next.ServeHTTP(w, r)
	})
}
