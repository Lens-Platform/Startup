package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	core_auth_sdk "github.com/Lens-Platform/Platform/src/libraries/core/core-auth-sdk"
	core_logging "github.com/Lens-Platform/Platform/src/libraries/core/core-logging/json"
)

type contextKey struct {
	name string
}

type AuthnMW struct {
	client *core_auth_sdk.Client
	logger core_logging.ILog
}

var ctxKey *contextKey

// NewAuthnMw returns a new instance of the authentication middleware
func NewAuthnMw(c *core_auth_sdk.Client, log core_logging.ILog, serviceName string) *AuthnMW {
	if serviceName == "" {
		panic("service name should be provided")
	}

	ctxKey = &contextKey{serviceName}
	return &AuthnMW{client: c, logger: log}
}

// AuthenticationMiddleware wraps the authentication middleware around an http call
func (mw *AuthnMW) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		authorization := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authorization, "Bearer ")
		decodedToken, err := mw.client.SubjectFrom(token)
		if err != nil {
			mw.logger.Error(err, "error")
			next.ServeHTTP(w, r)
			return
		}

		ctx = context.WithValue(ctx, ctxKey, decodedToken)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// IsAuthenticated returns wether or not the user is authenticated.
// REQUIRES Middleware to have run.
func IsAuthenticated(ctx context.Context) bool {
	return ctx.Value(ctxKey) != nil
}

// GetTokenFromCtx extracts the token from the context
func GetTokenFromCtx(ctx context.Context) (string, error) {
	if IsAuthenticated(ctx) {
		token, ok := ctx.Value(ctxKey).(string)
		if !ok {
			return "", errors.New("token cannot be converted to string")
		}

		return token, nil
	}

	return "", errors.New("token not found in context")
}

// InjectContextWithMockToken injects a token into the context. Useful for testing
func InjectContextWithMockToken(ctx context.Context, token string, serviceName string) context.Context {
	ctxKey = &contextKey{serviceName}
	return context.WithValue(ctx, ctxKey, token)
}
