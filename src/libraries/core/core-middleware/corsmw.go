package middleware

import (
	"net/http"
)

type CorsMw struct {
	origin string
}

// NewCorsMw returns a new instance of the cors middleware
func NewCorsMw(origins []string) *CorsMw {
	var origin = ""
	if len(origins) == 0 {
		origin = "*"
	}

	for i, singleOrigin := range origins {
		if i == 0 {
			origin += singleOrigin
		} else {
			origin += "," + singleOrigin
		}
	}

	return &CorsMw{origin: origin}
}

// CorsMiddleware runs the Cors middleware
func (c *CorsMw) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", c.origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
